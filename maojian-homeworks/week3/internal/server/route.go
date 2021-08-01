package server

import (
	"context"
	"fmt"
	"go.uber.org/atomic"
	"net/http"
)

func initRoute(ctx Context, sm *http.ServeMux) {
	sm.Handle("/doSth", md(ctx, doSth))
}

func md(ctx Context, f func(ctx *ReqCtx)) http.Handler {
	return &h{
		ctx: ctx,
		f:   f,
	}
}

type h struct {
	ctx Context
	f   func(ctx *ReqCtx)
}

var idG = atomic.NewInt32(0)

func (h *h) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("\n\n\n")
	id := idG.Add(1)
	logger.Infof("new request:%d", id)
	defer logger.Infof("request over:%d", id)
	cancel, cf := context.WithCancel(request.Context())
	defer cf()

	ctx := &ReqCtx{
		Context: cancel,
		req:     request,
		writer:  writer,
	}
	cmp := make(chan int)
	close(cmp)
	go func() {
		defer func() { cmp <- 1 }()
		h.f(ctx)
	}()
	for {
		select {
		case <-h.ctx.Done():
			return
		case <-cmp:
			return
		}
	}
}

type ReqCtx struct {
	context.Context
	req    *http.Request
	writer http.ResponseWriter
}
