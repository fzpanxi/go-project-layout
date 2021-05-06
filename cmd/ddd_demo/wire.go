// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"ddd_demo/internal/conf"
	"ddd_demo/internal/domain/service"
	"ddd_demo/internal/infra"
	"ddd_demo/internal/server"
	application "ddd_demo/internal/usecase/service"
	"github.com/google/wire"
)

func initApp(c *conf.Conf) (*App, func(), error) {
	panic(wire.Build(server.ProviderSet, infra.ProviderSet, service.ProviderSet, application.ProviderSet, newApp))
}
