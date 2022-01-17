package boot

import (
	"go-project/pkg/ginex"
	"go-project/pkg/repo"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var BaseSet = wire.NewSet(
	repo.Set,
	DefaultGin,
)

func DefaultGin() *gin.Engine {
	const sessionSalt = "xxx"
	const sessionName = "web-session"
	router := ginex.New(
		ginex.WithLogger(),
		ginex.WithRecovery(),
		ginex.WithPprof(),
		ginex.WithHSTS(),
		ginex.WithCors(),
		ginex.WithCookieSession(sessionName, sessionSalt),
		ginex.WithStatic(),
	)
	return router
}
