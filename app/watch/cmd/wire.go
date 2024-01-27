//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/uneva/dd-watch/app/watch/internal/biz"
	"github.com/uneva/dd-watch/app/watch/internal/conf"
	"github.com/uneva/dd-watch/app/watch/internal/data"
	"github.com/uneva/dd-watch/app/watch/internal/server"
	"github.com/uneva/dd-watch/app/watch/internal/service"
)

func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		newApp))
}
