package types

type (
	SendNotifyUser struct {
		// 邮箱
		Email string `json:"email"`

		// 手机号
		Phone string `json:"phone"`

		// todo dingding 企业微信
	}

	SendNotifyRequest struct {
		Priority  uint32         `json:"priority"`  // 优先级
		Keyword   string         `json:"keyword"`   // 模板
		User      SendNotifyUser `json:"user"`      // 用户
		Variable  map[string]any `json:"variable"`  // 变量
		Timestamp int64          `json:"timestamp"` // 时间
	}

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
	}
)
