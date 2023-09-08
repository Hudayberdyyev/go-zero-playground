package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/app"
	"github.com/Hudayberdyyev/go-zero-playground/third-party/cache"
	"log"

	"github.com/Hudayberdyyev/go-zero-playground/api/internal/config"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/handler"
	"github.com/Hudayberdyyev/go-zero-playground/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/beletvideo-server.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	appCtx := context.Background()

	cacheClient, err := cache.NewRedisClient(appCtx, cache.Params{
		Host: "0.0.0.0",
		Port: "6385",
		DB:   0,
	})
	if err != nil {
		log.Fatal(err)
	}

	app := app.Domain{
		Cache:         app.NewCache(cacheClient),
		Authorization: nil,
	}
	ctx := svc.NewServiceContext(c, app)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
