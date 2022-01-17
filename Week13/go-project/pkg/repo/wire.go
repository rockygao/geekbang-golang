package repo

import (
	"go-project/pkg/repo/impl"

	"github.com/google/wire"
)

var Set = wire.NewSet(
	wire.Struct(new(DbInitializer), "*"),

	wire.Struct(new(impl.UserRepoImpl), "*"),
	wire.Bind(new(UserRepo), new(*impl.UserRepoImpl)),
)
