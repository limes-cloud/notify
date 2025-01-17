package repository

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/notify/internal/types"
)

type Queue interface {
	Watch(hook func(ctx kratosx.Context, req *types.SendNotifyRequest) error)

	PushRedis(ctx kratosx.Context, req *types.SendNotifyRequest) error
}
