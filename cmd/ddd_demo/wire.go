// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"ddd_demo/internal/conf"
	"ddd_demo/internal/domain/service"
	"ddd_demo/internal/infra"
	application "ddd_demo/internal/usecase/service"
	"github.com/google/wire"
)

func initApp(c *conf.Conf) (*application.UserUseCase, func(), error) {
	panic(wire.Build(infra.ProviderSet, service.ProviderSet, application.ProviderSet))
}
