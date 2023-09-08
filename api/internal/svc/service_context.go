package svc

import (
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/app"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/config"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config                 config.Config
	VersionVerification    rest.Middleware
	OptionalAuthMiddleware rest.Middleware
	RequiredAuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config, app app.Domain) *ServiceContext {
	return &ServiceContext{
		Config:                 c,
		VersionVerification:    middleware.NewVersionVerificationMiddleware(app.Cache).Handle,
		OptionalAuthMiddleware: middleware.NewOptionalAuthMiddleware().Handle,
		RequiredAuthMiddleware: middleware.NewRequiredAuthMiddleware(app.Authorization).Handle,
	}
}
