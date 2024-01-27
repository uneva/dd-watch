package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

type WatchRepo interface{}

type WatchUsecase struct {
	repo WatchRepo
}

func NewWatchUsecase(repo WatchRepo, logger log.Logger) *WatchUsecase {
	return &WatchUsecase{
		repo: repo,
	}
}
