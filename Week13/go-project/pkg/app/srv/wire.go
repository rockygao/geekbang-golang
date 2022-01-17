//go:build wireinject
// +build wireinject

package srv

import (
	proto "go-project/api/user"
	"go-project/configs"
	"go-project/pkg/app/srv/internal/server/grpc"
	"go-project/pkg/app/srv/internal/server/http"
	"go-project/pkg/boot"
	"go-project/pkg/database"

	"github.com/google/wire"
)

func RunSrv() (*App, func(), error) {
	panic(wire.Build(
		configs.NewConfig,
		database.InitDB,
		wire.Struct(new(App), "*"),
		wire.Struct(new(Bootloader), "*"),

		boot.BaseSet,

		http.Set,
		grpc.Set,
		proto.Set,
	))
}
