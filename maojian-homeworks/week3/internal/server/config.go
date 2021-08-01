package server

import "context"

type Config struct {
	Id      string
	Port    int
	Context context.Context
}
