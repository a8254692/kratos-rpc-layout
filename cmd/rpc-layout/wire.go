// NOTE: https://github.com/google/wire/issues/106
//go:generate go run github.com/google/wire/cmd/wire

//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"gitlab.top.slotssprite.com/my/rpc-layout/internal/biz"
	"gitlab.top.slotssprite.com/my/rpc-layout/internal/data"
	"gitlab.top.slotssprite.com/my/rpc-layout/internal/server"
	"gitlab.top.slotssprite.com/my/rpc-layout/internal/service"
	"gitlab.top.slotssprite.com/my/rpc-layout/party/bootstrap"
)

// wireApp init kratos application.
func initApp(*bootstrap.ServiceInfo, registry.Registrar, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
