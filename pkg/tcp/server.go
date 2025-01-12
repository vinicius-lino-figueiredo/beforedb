package tcp

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
)

func NewServer() Server {
	return &server{}
}

type Server interface {
	Serve(ctx context.Context, port int) error
	AddCertificate(cert tls.Certificate)
}

type server struct {
	certificates []tls.Certificate
}

func (d *server) Serve(ctx context.Context, port int) (err error) {

	cfg := net.ListenConfig{}

	listener, err := cfg.Listen(ctx, "tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return
	}

	tlsCfg := &tls.Config{
		Certificates: d.certificates,
	}

	listener = tls.NewListener(listener, tlsCfg)

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go d.HandleConnection(conn)
	}
}

func (d *server) HandleConnection(conn net.Conn) {
}

func (d *server) AddCertificate(cert tls.Certificate) {
	d.certificates = append(d.certificates, cert)
}
