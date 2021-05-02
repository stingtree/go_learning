package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	stop_req := make(chan string)
	muxhttp := http.NewServeMux()
	muxhttp.HandleFunc("/http_start", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("errgroup_httpsev"))
	})
	muxhttp.HandleFunc("/http_stop", func(writer http.ResponseWriter, request *http.Request) {
		stop_req <- "Receviced"
	})

	server := http.Server{
		Handler: muxhttp,
		Addr:    ":8080",
	}

	//启动http服务
	g.Go(func() error {
		return server.ListenAndServe()
	})
	//停止http服务
	g.Go(func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-stop_req:
			log.Fatal("Received http_stop request")
		}
		fmt.Println("Service stopping")
		stopser_ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		return server.Shutdown(stopser_ctx)
	})

	//监听linux信号
	g.Go(func() error {
		//接受linux信号的channel
		linux_s := make(chan os.Signal,0)
		//定义要接受的linux信号类型
		signal.Notify(linux_s, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-linux_s:
			return errors.New("Receviced the Exit Signal")
		}
		return nil
	})
	fmt.Printf("The Error is %v", g.Wait())
}
