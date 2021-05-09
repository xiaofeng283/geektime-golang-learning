package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

type myHandler struct {
}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

// HTTP server
func httpServer(ctx context.Context, stops <-chan struct{}) error {
	server := &http.Server{Addr: "127.0.0.1:8000", Handler: &myHandler{}}
	go func() {
		select {
		case <-stops:
			log.Print("Get信号, server shutdown")
			server.Shutdown(ctx)
			return
		case <-ctx.Done():
			log.Print("Get超时, server shutdown")
			server.Shutdown(ctx)
			return
		}
	}()
	return server.ListenAndServe()
}

// ListenSignal
func ListenSignal(ctx context.Context, stops chan struct{}) error {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGQUIT)
	select {
	case <-ch:
		log.Print("信号，closed")
		close(stops)
		return errors.New("receive signal")
	case <-ctx.Done():
		log.Print("超时，closed")
		return ctx.Err()
	}
}

func main() {

	var stops = make(chan struct{})

	c, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*10000))
	defer cancel()
	g, ctx := errgroup.WithContext(c)

	// HTTP服务
	g.Go(func() error {
		err := httpServer(ctx, stops)
		if err != nil {
			return err
		}
		return nil
	})

	// 捕捉信号
	g.Go(func() error {
		err := ListenSignal(ctx, stops)
		if err != nil {
			return err
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("failed：%v", err.Error())
	}
}