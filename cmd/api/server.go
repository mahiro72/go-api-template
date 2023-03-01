package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mahiro72/go-api-template/pkg/adapter/http/router"
	"github.com/mahiro72/go-api-template/pkg/infra/registry"
)

func main() {
	repo, err := registry.NewRepository()
	if err != nil {
		panic(fmt.Sprintf("error: registry.NewRepository: %v", err))
	}

	r := router.New(repo)

	srv := &http.Server{Addr: ":8080", Handler: r}
	srvCtx, srvStopCtx := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// 猶予期間が30秒のグレースフルシャットダウンを開始する
		shutdownCtx, shutdownCancel := context.WithTimeout(srvCtx, 30*time.Second)
		defer shutdownCancel()

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Println("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// グレースフルシャットダウンのトリガー
		err := srv.Shutdown(shutdownCtx)
		if err != nil {
			log.Println(err.Error())
		}
		srvStopCtx()
	}()

	err = srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Println(err.Error())
	}

	// サーバーのcontextが終了するまで待機する
	<-srvCtx.Done()
}
