package hdl

import (
	"context"
	proto "go-project/api/user/v1"

	"github.com/gin-gonic/gin"
	"github.com/quexer/utee"
	"google.golang.org/grpc/grpclog"
)

type Users struct {
	UserClient proto.UserClient
}

func (p *Users) Mount(g gin.IRouter) {
	g.GET("/world", p.World)
}

func (p *Users) World(c *gin.Context) {
	req := &proto.IdRequest{Id: 1}
	res, err := p.UserClient.GetUserById(context.Background(), req)

	if err != nil {
		grpclog.Fatalln(err)
	}

	c.JSON(200, utee.J{"data": res})
}
