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

// HTTP服务
func httpServer(ctx context.Context, Addr string, stops <-chan struct{}) error {
	server := &http.Server{Addr: Addr, Handler: &myHandler{}}
	go func() {
		select {
		case <-stops:
			log.Print("信号, server shutdown")
			server.Shutdown(ctx)
			return
		case <-ctx.Done():
			log.Print("超时, server shutdown")
			server.Shutdown(ctx)
			return
		}
	}()
	return server.ListenAndServe()
}

// 信号
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

	c, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*30000))
	defer cancel()
	g, ctx := errgroup.WithContext(c)

	// HTTP服务
	g.Go(func() error {
		err := httpServer(ctx, ":8000", stops)
		if err != nil {
			return err
		}
		return nil
	})

	g.Go(func() error {
		err := httpServer(ctx, ":8888", stops)
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