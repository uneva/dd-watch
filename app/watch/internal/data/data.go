package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/uneva/dd-watch/app/watch/internal/conf"
)

var ProviderSet = wire.NewSet(NewData, NewWatchRepo)

type Data struct{}

func NewData(conf *conf.Data, logger log.Logger) (*Data, func(), error) {
	// log := log.NewHelper(logger)
	return &Data{}, nil, nil
}
