package tcp

import (
	"context"
)

func NewServer() Server {
	return &server{}
}

type Server interface {
	Serve(ctx context.Context, port int) error
}

type server struct {
}

func (d *server) Serve(_ context.Context, _ int) (err error) {
	return
}
