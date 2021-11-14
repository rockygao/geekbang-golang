//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

// func main() {
// 	router := gin.Default()

// 	router.GET("/", func(c *gin.Context) {
// 		c.String(http.StatusOK, "Hello world")
// 	})
// 	router.GET("/user/:name", func(c *gin.Context) {
// 		name := c.Param("name")
// 		c.String(http.StatusOK, "Hello %s", name)
// 	})
// 	router.Run(":8000")
// }
func main() {
	server, e := Init()
	if e != nil {
		panic(e)
	}
	group, ctx := errgroup.WithContext(context.Background())
	group.Go(func() error {
		go func() {
			<-ctx.Done()
			server.Shutdown()
		}()
		return server.Start(ctx)
	})

	group.Go(func() error {
		return handleSysSignal(ctx)
	})

	err := group.Wait()
	fmt.Printf("exited %s", err)
}

func handleSysSignal(ctx context.Context) error {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	select {
	case <-ctx.Done():
		log.Printf("context done receive, exiting system signal listen")
		return ctx.Err()
	case sig := <-sigs:
		log.Printf("receive sig %s, exiting...", sig)
		return fmt.Errorf("receive sig %v, exiting", sig)
	}
}
