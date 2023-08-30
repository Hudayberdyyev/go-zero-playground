package svc

import (
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/config"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/middleware"
	"github.com/Hudayberdyyev/go-zero-playground/rpc/trending/trendingclient"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config                  config.Config
	AuthorizationMiddleware rest.Middleware
	OtherMiddleware         rest.Middleware
	TrendingClient          trendingclient.Trending
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                  c,
		AuthorizationMiddleware: middleware.NewAuthorizationMiddleware().Handle,
		OtherMiddleware:         middleware.NewOtherMiddleware().Handle,
		TrendingClient:          trendingclient.NewTrending(zrpc.MustNewClient(c.Trending)),
	}
}
