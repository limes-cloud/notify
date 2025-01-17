package service

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	kerrors "github.com/go-kratos/kratos/v2/errors"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/library/pool"
	ktypes "github.com/limes-cloud/kratosx/types"

	"github.com/limes-cloud/notify/api/notify/errors"
	"github.com/limes-cloud/notify/internal/conf"
	"github.com/limes-cloud/notify/internal/domain/entity"
	"github.com/limes-cloud/notify/internal/domain/repository"
	"github.com/limes-cloud/notify/internal/infra/sender"
	"github.com/limes-cloud/notify/internal/types"
)

type Notify struct {
	conf     *conf.Config
	repo     repository.Notify
	template repository.Template
	sender   repository.Sender
	log      repository.Log
	queue    repository.Queue
}

const (
	SendModePriority = "priority"
)

func NewNotify(
	conf *conf.Config,
	repo repository.Notify,
	template repository.Template,
	sender repository.Sender,
	log repository.Log,
	queue repository.Queue,
) *Notify {
	notify := &Notify{
		conf:     conf,
		repo:     repo,
		template: template,
		sender:   sender,
		log:      log,
		queue:    queue,
	}
	queue.Watch(notify.SendNotify)
	return notify
}

// ListNotifyCategory 获取通知分类列表
func (u *Notify) ListNotifyCategory(ctx kratosx.Context, req *types.ListNotifyCategoryRequest) ([]*entity.NotifyCategory, uint32, error) {
	list, total, err := u.repo.ListNotifyCategory(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list notify category error", "err", err.Error())
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateNotifyCategory 创建通知分类
func (u *Notify) CreateNotifyCategory(ctx kratosx.Context, req *entity.NotifyCategory) (uint32, error) {
	id, err := u.repo.CreateNotifyCategory(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "create notify category error", "err", err.Error())
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateNotifyCategory 更新通知分类
func (u *Notify) UpdateNotifyCategory(ctx kratosx.Context, req *entity.NotifyCategory) error {
	if err := u.repo.UpdateNotifyCategory(ctx, req); err != nil {
		ctx.Logger().Warnw("msg", "update notify category error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteNotifyCategory 删除通知分类
func (u *Notify) DeleteNotifyCategory(ctx kratosx.Context, id uint32) error {
	if err := u.repo.DeleteNotifyCategory(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete notify category error", "err", err.Error())
		return errors.DeleteError(err.Error())
	}
	return nil
}

// GetNotify 获取指定通知信息
func (u *Notify) GetNotify(ctx kratosx.Context, id uint32) (*entity.Notify, error) {
	notify, err := u.repo.GetNotify(ctx, id)
	if err != nil {
		ctx.Logger().Warnw("msg", "list notify error", "err", err.Error())
		return nil, errors.GetError(err.Error())
	}
	return notify, nil
}

// ListNotify 获取通知列表
func (u *Notify) ListNotify(ctx kratosx.Context, req *types.ListNotifyRequest) ([]*entity.Notify, uint32, error) {
	list, total, err := u.repo.ListNotify(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list notify error", "err", err.Error())
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateNotify 创建通知
func (u *Notify) CreateNotify(ctx kratosx.Context, notify *entity.Notify) (uint32, error) {
	id, err := u.repo.CreateNotify(ctx, notify)
	if err != nil {
		ctx.Logger().Warnw("msg", "create notify error", "err", err.Error())
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateNotify 更新通知
func (u *Notify) UpdateNotify(ctx kratosx.Context, notify *entity.Notify) error {
	if err := u.repo.UpdateNotify(ctx, notify); err != nil {
		ctx.Logger().Warnw("msg", "update notify error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteNotify 删除通知
func (u *Notify) DeleteNotify(ctx kratosx.Context, id uint32) error {
	// 获取当前的notify
	notify, err := u.repo.GetNotify(ctx, id)
	if err != nil {
		return errors.DeleteError(err.Error())
	}

	notify.DeletedAt = ktypes.NullInt64{
		Valid: true,
		Int64: time.Now().Unix(),
	}
	if err := u.repo.UpdateNotify(ctx, notify); err != nil {
		ctx.Logger().Warnw("msg", "update notify error", "err", err.Error())
		return errors.UpdateError(err.Error())
	}
	return nil
}

// AsyncSendNotify 发送通知
func (u *Notify) AsyncSendNotify(ctx kratosx.Context, req *types.SendNotifyRequest) error {
	return u.queue.PushRedis(ctx, req)
}

// SendNotify 发送通知
func (u *Notify) SendNotify(ctx kratosx.Context, req *types.SendNotifyRequest) error {
	getUser := func(user *types.SendNotifyUser, ct string) (string, bool) {
		if u.sender.IsEmail(ct) {
			return user.Email, user.Email != ""
		}
		if u.sender.IsOfficialAccount(ct) {
			return user.OfficialOpenid, user.OfficialOpenid != ""
		}
		return "", false
	}

	// 获取指定通知
	notify, err := u.repo.GetNotifyByKeyword(ctx, req.Notify)
	if err != nil {
		return errors.SendNotifyError(err.Error())
	}

	// 判断是否已经超时
	if notify.Expire != nil && *notify.Expire != 0 {
		if time.Now().Unix()-req.Timestamp > int64(*notify.Expire) {
			return errors.NotifyExpireError()
		}
	}

	// 获取发送模板
	tpls, err := u.template.ListTemplate(ctx, notify.Id)
	if err != nil {
		return errors.SendNotifyError(err.Error())
	}

	// 去除不可用的渠道模板
	var ntpls []*types.SendNotifyInnerRequest
	for _, item := range tpls {
		// 判断渠道是否被禁用
		if item.Channel == nil {
			continue
		}

		// 判断用户是否可用
		user, is := getUser(req.User, item.Channel.Type)
		if !is {
			continue
		}

		// 判断是否已经发送
		if notify.Cache != nil {
		}

		ntpls = append(ntpls, &types.SendNotifyInnerRequest{
			NotifyId:   notify.Id,
			ChannelId:  item.ChannelId,
			FromServer: req.FromServer,
			User:       user,
			Variable:   req.Variable,
			Title:      notify.Title,
			Content:    item.Content,
			Type:       item.Channel.Type,
			AK:         item.Channel.GetAkString(),
			SK:         item.Channel.GetSkString(),
			Extra:      item.Channel.GetExtraString(),
			IP:         req.IP,
		})
	}
	if len(ntpls) == 0 {
		return errors.SendNotifyError("无可用的发送渠道")
	}

	// 获取发送模式
	if notify.SendMode == SendModePriority || len(ntpls) == 1 {
		return u.send(ctx, ntpls[0])
	}

	runner := ctx.WaitRunner(pool.WithWaitRunnerOption(int32(len(ntpls))))
	for _, item := range ntpls {
		tp := *item
		runner.AddTask(ctx, func() error {
			return u.send(ctx, &tp)
		})
	}
	runner.Wait()

	return nil
}

// send 发送通知
func (u *Notify) send(ctx kratosx.Context, req *types.SendNotifyInnerRequest) (err error) {
	// 获取发送模板
	body, err := u.getSendBody(req.Variable, req.Content)
	if err != nil {
		return err
	}

	defer func(sendTime time.Time, body string) {
		var (
			errCode  = 0
			now      = time.Now()
			reason   = ""
			variable = ""
		)
		if err != nil {
			reason = err.Error()
			errCode = int(kerrors.FromError(err).Code)
		}
		vb, _ := json.Marshal(req.Variable)
		variable = string(vb)

		// 写入发送日志
		log := entity.Log{
			NotifyId:   req.NotifyId,
			ChannelId:  req.ChannelId,
			User:       req.User,
			Content:    body,
			Variable:   variable,
			FromServer: req.FromServer,
			FromIp:     req.IP,
			Status:     err == nil,
			Reason:     reason,
			ErrCode:    errCode,
			Cost:       uint32(now.Unix() - sendTime.Unix()),
		}
		if _, er := u.log.CreateLog(ctx, &log); er != nil {
			ctx.Logger().Errorw("msg", "log write db error", "err", er.Error(), "log", log)
		}
	}(time.Now(), body)

	// 组装模板
	initFunc, err := u.sender.GetSender(req.Type)
	if err != nil {
		return errors.SendNotifyError(err.Error())
	}

	// 初始化发送器
	executor, err := initFunc(
		sender.WithAK(req.AK),
		sender.WithSK(req.SK),
		sender.WithExtra(req.Extra),
	)
	if err != nil {
		return errors.SendNotifyError(err.Error())
	}

	// 发送消息
	if err := executor.Send(req.User, req.Title, body); err != nil {
		return errors.SendNotifyError(err.Error())
	}

	// 标记为发送成果
	return nil
}

func (u *Notify) fillKey(val string) string {
	return fmt.Sprintf(`${%s}`, val)
}

func (u *Notify) getSendBody(variable map[string]string, content string) (string, error) {
	for key, val := range variable {
		content = strings.ReplaceAll(content, u.fillKey(key), val)
	}

	reg, _ := regexp.Compile(`\$\{(\w|\.)+}`)
	tempKeys := reg.FindAllString(content, -1)
	if len(tempKeys) != 0 {
		return "", errors.TemplateVariableError(strings.Join(tempKeys, ","))
	}

	return content, nil
}
