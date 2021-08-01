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
	ctx  Context
	conf *Config
	cf   func()
}

var myServer http.Server

func (d *defaultServer) Start() error {
	logger.Infof("start server")

	sm := http.NewServeMux()

	initRoute(d.ctx, sm)

	myServer = http.Server{
		Addr:    net.JoinHostPort("0.0.0.0", strconv.Itoa(d.conf.Port)),
		Handler: sm,
	}
	return myServer.ListenAndServe()
}

func (d *defaultServer) Stop() {
	d.cf()
	logger.Info("server:stopping")
	myServer.Shutdown(context.Background())
	<-time.After(3 * time.Second)
}

func NewServer(conf *Config) Server {
	c, cf := context.WithCancel(context.Background())

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
