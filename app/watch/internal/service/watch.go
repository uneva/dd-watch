package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/uneva/dd-watch/api/watch/v1"
	"github.com/uneva/dd-watch/app/watch/internal/biz"
)

func NewWatchService(watch *biz.WatchUsecase, logger log.Logger) *WatchService {
	return &WatchService{
		wc:  watch,
		log: log.NewHelper(log.With(logger, "module", "service/watch")),
	}
}

func (ws *WatchService) ToLink(ctx context.Context, req *v1.ToLinkRequest) (*v1.ToLinkResponse, error) {
	ws.log.Infof("input data %v", req)

	ws.wc.ToLink(ctx, req.Uid)

	return &v1.ToLinkResponse{}, nil
}
