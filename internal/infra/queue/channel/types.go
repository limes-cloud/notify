package channel

import "context"

type (
	Queue interface {
		Push(ctx context.Context, data *Item) error
		Pop() (*Item, error)
		Name() string
	}

	Item struct {
		Trace string `json:"trace"`
		Span  string `json:"span"`
		Data  string `json:"data"`
	}
)
