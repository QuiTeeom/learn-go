package server

import (
	"context"
	"net"
	"net/http"
	"strconv"
	"time"
)

type Server interface {
	Start() error
	Stop()
}

type defaultServer struct {
	ctx    Context
	conf   *Config
	cf     func()
	server http.Server
}

func (d *defaultServer) Start() error {
	logger.Infof("start server:%+v", d.conf)

	sm := http.NewServeMux()

	initRoute(d.ctx, sm)

	d.server = http.Server{
		Addr:    net.JoinHostPort("0.0.0.0", strconv.Itoa(d.conf.Port)),
		Handler: sm,
	}

	errSig := make(chan error)
	close(errSig)
	go func() {
		err := d.server.ListenAndServe()
		if err != nil {
			errSig <- err
		}
	}()
	for {
		select {
		case err := <-errSig:
			return err
		case <-d.ctx.Done():
			d.Stop()
			return nil
		}
	}
}

func (d *defaultServer) Stop() {
	d.cf()
	logger.Info("server:stopping")
	d.server.Shutdown(context.Background())
	<-time.After(3 * time.Second)
}

func NewServer(conf *Config) Server {
	c, cf := context.WithCancel(conf.Context)

	return &defaultServer{
		conf: conf,
		cf:   cf,
		ctx: Context{
			Context:   c,
			StartTime: time.Now(),
		},
	}
}

type Context struct {
	context.Context
	StartTime time.Time
}
