//go:build wireinject
// +build wireinject

package main

import (
	"go_code/geekbang-golang/Week4/configs"
	"go_code/geekbang-golang/Week4/internal/biz"
	"go_code/geekbang-golang/Week4/internal/data"
	"go_code/geekbang-golang/Week4/internal/server"

	"github.com/google/wire"
)

func Init() (*server.UserServer, error) {
	wire.Build(server.NewUserServer,
		server.NewGrpcServer,
		server.NewHttpServer,
		biz.NewUserService,
		data.NewInMemoUserDao,
		configs.NewConfig,
		configs.InitConfig)
	return &server.UserServer{}, nil
}
