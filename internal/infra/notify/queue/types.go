package queue

type (
	Queue interface {
		Push(data *Item) error
		Pop() (*Item, error)
		Name() string
	}

	Item struct {
		Priority int64
		ItemData *ItemData
	}

	ItemData struct {
		Trace string `json:"trace"`
		Span  string `json:"span"`
		Data  any    `json:"data"`
	}
)
