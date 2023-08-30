package main

import (
	"flag"
	"fmt"

	"github.com/Hudayberdyyev/go-zero-playground/rpc/trending/internal/config"
	"github.com/Hudayberdyyev/go-zero-playground/rpc/trending/internal/server"
	"github.com/Hudayberdyyev/go-zero-playground/rpc/trending/internal/svc"
	"github.com/Hudayberdyyev/go-zero-playground/rpc/trending/trending"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/trending.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		trending.RegisterTrendingServer(grpcServer, server.NewTrendingServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
