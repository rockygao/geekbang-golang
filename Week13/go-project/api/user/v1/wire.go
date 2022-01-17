package v1

import (
	proto "go-project/api/user/v1"

	"github.com/google/wire"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var Set = wire.NewSet(
	HelloRpcClient,
)

func HelloRpcClient(config *viper.Viper) proto.UserClient {
	// 127.0.0.1:8889
	conn, err := grpc.Dial(config.GetString("grpc.port"), grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}
	// 初始化客户端
	c := proto.NewUserClient(conn)

	return c
}
