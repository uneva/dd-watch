package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "github.com/uneva/dd-watch/api/watch/v1"
	"github.com/uneva/dd-watch/app/watch/internal/biz"
)

var ProviderSet = wire.NewSet(NewWatchService)

type WatchService struct {
	v1.UnimplementedWatchServiceServer

	log *log.Helper

	wc *biz.WatchUsecase
}
