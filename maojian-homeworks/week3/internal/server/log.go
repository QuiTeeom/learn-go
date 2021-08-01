package server

import "go.uber.org/zap"

var logger *zap.SugaredLogger

func init() {
	if l, err := zap.NewProduction(); err != nil {
		panic(err)
	} else {
		logger = l.Sugar()
	}
}
