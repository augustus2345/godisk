package svc

import (
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/rest"
	"go_disk/core/internal/config"
	"go_disk/core/internal/middleware"
	"go_disk/core/modles"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *xorm.Engine
	RDB    *redis.Client
	Auth   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: modles.Init(c),
		RDB:    modles.InitRedis(c),
		Auth:   middleware.NewAuthMiddleware().Handle,
	}
}
