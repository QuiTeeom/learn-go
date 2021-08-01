package main

import (
	"fmt"
	"learn-go/maojian-homeworks/week3/internal/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	s := server.NewServer(&server.Config{Port: 8080})

	errSig := make(chan error)

	go func() {
		err := s.Start()
		if err != nil {
			errSig <- err
		}
	}()

	osSig := make(chan os.Signal)
	signal.Notify(osSig, syscall.SIGKILL, syscall.SIGINT)

	for {
		select {
		case err := <-errSig:
			panic(err)
		case sig := <-osSig:
			fmt.Printf("\nreceive signal:%s\n", sig)
			s.Stop()
			os.Exit(0)
		}
	}
}
