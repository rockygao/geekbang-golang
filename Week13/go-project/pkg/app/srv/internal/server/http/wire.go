package http

import (
	"go-project/pkg/app/srv/internal/server/http/hdl"

	"github.com/google/wire"
)

var Set = wire.NewSet(
	wire.Struct(new(Initializer), "*"),
	wire.Struct(new(hdl.Hdl), "*"),
	wire.Struct(new(hdl.Users), "*"),
)
