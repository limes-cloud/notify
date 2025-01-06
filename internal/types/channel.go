package types

type (
	ListChannelRequest struct {
		Page     uint32  `json:"page"`
		PageSize uint32  `json:"pageSize"`
		Type     *string `json:"type"`
		Name     *string `json:"name"`
		Status   *bool   `json:"status"`
	}
)
