package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"golang.org/x/sync/errgroup"
)

//基于 errgroup 实现一个 http server 的启动和关闭 ，
//以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

//思路梳理

func main() {
	// go serveApp()
	// go serveDebug()
	// select {}

	//老师案例 errgroup 使用
	// g, ctx := errgroup.WithContext(context.Background())

	// var a, b, c []int

	// //调用广告
	// g.Go(func() error {
	// 	time.Sleep(time.Duration(2) * time.Second)
	// 	a = []int{1, 2, 3}
	// 	return errors.New("aaaa")
	// })

	// //调用AI
	// g.Go(func() error {
	// 	time.Sleep(time.Duration(8) * time.Second)
	// 	b = []int{4, 5, 6}
	// 	return errors.New("bbbb")
	// })

	// //调用运营平台
	// g.Go(func() error {
	// 	time.Sleep(time.Duration(10) * time.Second)
	// 	c = []int{1, 3, 5}
	// 	return errors.New("cccc")
	// })

	// err := g.Wait()
	// fmt.Println(a, b, c)
	// fmt.Println(err)
	// fmt.Println(ctx.Err())

	//实现 HTTP server 的启动和关闭

	g, ctx := errgroup.WithContext(context.Background())

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world! This is index!")
	})

	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	//测试http server 阻塞
	serverOut := make(chan int)
	mux.HandleFunc("/out", func(w http.ResponseWriter, r *http.Request) {
		serverOut <- 1
	})

	//g1 启动http server服务
	g.Go(func() error {
		return server.ListenAndServe()
	})

	//g2
	g.Go(func() error {
		select {
		case <-serverOut:
			fmt.Println("server out exit ...")
		case <-ctx.Done():
			fmt.Println("g2 errgroup exit ...")
		}
		fmt.Println("when do here")
		timeoutCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		fmt.Println("shutting down ....")
		return server.Shutdown(timeoutCtx)
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

func serveApp() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hello Qcan!")
	})
	if err := http.ListenAndServe("0.0.0.0:8080", mux); err != nil {
		log.Fatal(err)
	}
}

func serveDebug() {
	if err := http.ListenAndServe("127.0.0.1:8081", http.DefaultServeMux); err != nil {
		log.Fatal(err)
	}
}
