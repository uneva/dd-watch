package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/uneva/dd-watch/api/watch/v1"
	"github.com/uneva/dd-watch/app/watch/internal/biz"
)

func NewWatchService(watch *biz.WatchUsecase, logger log.Logger) *WatchService {
	return &WatchService{
		watch: watch,
		log:   log.NewHelper(logger),
	}
}

func (s *WatchService) ToLink(ctx context.Context, req *v1.ToLinkRequest) (*v1.ToLinkResponse, error) {
	s.log.Infof("input data %v", req)
	return &v1.ToLinkResponse{}, nil
}
