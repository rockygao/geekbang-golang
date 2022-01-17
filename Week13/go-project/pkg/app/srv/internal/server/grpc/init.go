package grpc

import (
	"context"
	"fmt"
	proto "go-project/api/user/v1"
	"net"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// Initializer grpc 初始化器
type Initializer struct {
	Config      *viper.Viper
	UserService *UserService
}

func (p *Initializer) Name() string {
	return "grpc_initializer"
}

func (p *Initializer) IsNeedInit(ctx context.Context) (bool, error) {
	return true, nil
}

// Initialize 进行micro grpc 注册
func (p *Initializer) Initialize(ctx context.Context) error {
	go p.startRpc()
	return nil
}

//v.GetString("data.database.source")
func (p *Initializer) startRpc() error {
	listen, err := net.Listen("tcp", p.Config.GetString("server.grpc.addr"))

	fmt.Println(p.Config.GetString("server.grpc.addr"))
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}

	// 实例化grpc Server
	s := grpc.NewServer()

	// 注册UserService
	proto.RegisterUserServer(s, p.UserService)

	s.Serve(listen)
	return nil
}
