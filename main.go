package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/simiken/golang-cleanarchitecture/interface/handler"
)

func main() {

	routeSetting()
	srv := &http.Server{Addr: ":28080"}
	// サーバはブロックするので別の goroutine で実行する
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// シグナルを待つ
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM)
	<-sigCh

	// シグナルを受け取ったらShutdown
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatal(err)
	}

}

// TODO
func routeSetting() {

	http.HandleFunc("/", handler.Index())

}

// TODO
func getHandlerFuncByHttpMethod(path string, m string) func(w http.ResponseWriter, r *http.Request) {

	switch m {
	case http.MethodGet:
		return func(w http.ResponseWriter, r *http.Request) {}
	case http.MethodPut:
		return func(w http.ResponseWriter, r *http.Request) {}
	case http.MethodPatch:
		return func(w http.ResponseWriter, r *http.Request) {}
	case http.MethodPost:
		return func(w http.ResponseWriter, r *http.Request) {}
	case http.MethodDelete:
		return func(w http.ResponseWriter, r *http.Request) {}
	default:
		return func(w http.ResponseWriter, r *http.Request) {}
	}

}
