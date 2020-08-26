package main

import (
	"log"
	"net"
	"os"

	"github.com/devlights/go-grpc-uds-example/internal/pb"
	"github.com/devlights/go-grpc-uds-example/internal/service"
	"google.golang.org/grpc"
)

const (
	protocol = "unix"
	sockAddr = "/tmp/echo.sock"
)

func main() {
	// - https://qiita.com/marnie_ms4/items/4582a1a0db363fe246f3
	// - http://yamahiro0518.hatenablog.com/entry/2016/02/01/215908

	cleanup := func() {
		if _, err := os.Stat(sockAddr); err == nil {
			if err := os.RemoveAll(sockAddr); err != nil {
				log.Fatal(err)
			}
		}
	}

	cleanup()

	listener, err := net.Listen(protocol, sockAddr)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	echo := service.NewEchoService()

	pb.RegisterEchoServer(server, echo)

	server.Serve(listener)
}
