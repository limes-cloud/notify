package types

type (
	ListNotifyCategoryRequest struct {
		Page     uint32  `json:"page"`
		PageSize uint32  `json:"pageSize"`
		Name     *string `json:"name"`
	}

	ListNotifyRequest struct {
		Page       uint32  `json:"page"`
		PageSize   uint32  `json:"pageSize"`
		CategoryId *uint32 `json:"categoryId"`
		Status     *string `json:"status"`
		Keyword    *string `json:"keyword"`
		Name       *string `json:"name"`
		Priority   *uint32 `json:"priority"`
	}

	SendNotifyUser struct {
		Email          string `json:"email"`
		OfficialOpenid string `json:"officialOpenid"`
	}

	SendNotifyRequest struct {
		FromServer string            `json:"fromServer"`
		Notify     string            `json:"notify"`    // 标识
		User       *SendNotifyUser   `json:"user"`      // 用户
		Variable   map[string]string `json:"variable"`  // 变量
		Timestamp  int64             `json:"timestamp"` // 时间
		IP         string            `json:"ip"`
	}

	SendNotifyInnerRequest struct {
		FromServer string            `json:"fromServer"`
		NotifyId   uint32            `json:"notifyId"`
		ChannelId  uint32            `json:"channelId"`
		User       string            `json:"user"`
		Variable   map[string]string `json:"variable"`
		Title      string            `json:"title"`
		Content    string            `json:"content"`
		Type       string            `json:"type"`
		AK         string            `json:"ak"`
		SK         string            `json:"sk"`
		Extra      string            `json:"extra"`
		IP         string            `json:"ip"`
	}
)
