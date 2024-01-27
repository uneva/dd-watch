package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/uneva/dd-watch/app/watch/internal/biz"
)

type watchRepo struct {
	data *Data
	log  *log.Helper
}

func NewWatchRepo(data *Data, logger log.Logger) biz.WatchRepo {
	return &watchRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/watch")),
	}
}

func (wr *watchRepo) ToLink(ctx context.Context, uid int64) error {
	wr.log.Infof("uid %v", uid)
	return nil
}
