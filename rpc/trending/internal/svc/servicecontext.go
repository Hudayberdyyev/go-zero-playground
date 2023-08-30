package svc

import "github.com/Hudayberdyyev/go-zero-playground/rpc/trending/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
