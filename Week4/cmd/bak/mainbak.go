package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

func main() {

	//实现 HTTP server 的启动和关闭

	g, ctx := errgroup.WithContext(context.Background())

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world")
	})
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	//g1 启动http server服务
	g.Go(func() error {
		go func() {
			<-ctx.Done()
			timeoutCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()
			srv.Shutdown(timeoutCtx)
		}()
		return srv.ListenAndServe()
	})

	//g3 linux signal 信号的注册和处理  按ctrl + c 退出程序
	g.Go(func() error {
		c := make(chan os.Signal)
		signal.Notify(c)
		fmt.Println("start ...")
		select {
		case <-ctx.Done():
			fmt.Println("g3 errgroup done ...")
		case s := <-c:
			fmt.Println("os exit end ...", s)
		}
		return errors.New("chan exit...")
	})

	err := g.Wait()
	fmt.Println(err)
	fmt.Println(ctx.Err())
}
