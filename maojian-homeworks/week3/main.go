package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"learn-go/maojian-homeworks/week3/internal/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	c, cf := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(c)

	s1 := server.NewServer(&server.Config{Id: "1", Port: 8080, Context: ctx})
	s2 := server.NewServer(&server.Config{Id: "2", Port: 8083, Context: ctx})

	servers := []server.Server{s1, s2}

	for _, s := range servers {
		eg.Go(s.Start)
	}

	eg.Go(func() error {
		osSig := make(chan os.Signal)
		signal.Notify(osSig, syscall.SIGKILL, syscall.SIGINT)
		select {
		case sig := <-osSig:
			fmt.Printf("\nreceive signal:%s\n", sig)
			cf()
		case <-ctx.Done():
		}
		return nil
	})

	err := eg.Wait()
	if err != nil {
		panic(err)
	}
}
