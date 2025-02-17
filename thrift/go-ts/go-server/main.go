package main

import (
	"context"
	"fmt"
	"go-server/gen-go/demo"

	"github.com/apache/thrift/lib/go/thrift"
)

type PingHandler struct{}

func (p *PingHandler) Ping(ctx context.Context) (string, error) {
	return "pong", nil
}

func main() {
	handler := &PingHandler{}
	processor := demo.NewPingServiceProcessor(handler)

	transport, err := thrift.NewTServerSocket(":9090")
	if err != nil {
		panic(err)
	}

	server := thrift.NewTSimpleServer4(
		processor,
		transport,
		thrift.NewTBufferedTransportFactory(8192),
		thrift.NewTBinaryProtocolFactoryDefault(),
	)

	fmt.Println("Starting Go Thrift server...")
	server.Serve()
}
