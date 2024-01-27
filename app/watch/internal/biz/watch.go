package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type WatchRepo interface{}

type WatchUsecase struct {
	repo WatchRepo
	log  *log.Helper
}

func NewWatchUsecase(repo WatchRepo, logger log.Logger) *WatchUsecase {
	return &WatchUsecase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/watch")),
	}
}

func (uc *WatchUsecase) ToLink(ctx context.Context, uid int64) error {
	uc.log.Infof("uid %v", uid)
	return nil
}
