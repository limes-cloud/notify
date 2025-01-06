package types

type ListLogRequest struct {
	Page      uint32  `json:"page"`
	PageSize  uint32  `json:"pageSize"`
	User      *string `json:"user"`
	Index     *uint32 `json:"index"`
	NotifyId  *uint32 `json:"notifyId"`
	ChannelId *uint32 `json:"channelId"`
}
