package tcp

import (
	"context"
	"fmt"
	"net"
)

func NewServer() Server {
	return &server{}
}

type Server interface {
	Serve(ctx context.Context, port int) error
}

type server struct {
}

func (d *server) Serve(ctx context.Context, port int) (err error) {

	cfg := net.ListenConfig{}

	listener, err := cfg.Listen(ctx, "tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return
	}

	defer listener.Close()

	for {
		_, err := listener.Accept()
		if err != nil {
			return err
		}
	}
}
