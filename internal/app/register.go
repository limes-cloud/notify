package app

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/limes-cloud/notify/internal/conf"
)

type registerFunc func(c *conf.Config, hs *http.Server, gs *grpc.Server)

var registries []registerFunc

func register(fn registerFunc) {
	registries = append(registries, fn)
}

func New(c *conf.Config, hs *http.Server, gs *grpc.Server) {
	for _, registry := range registries {
		registry(c, hs, gs)
	}
}
