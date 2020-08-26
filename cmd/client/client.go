package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/devlights/go-grpc-uds-example/internal/pb"
	"google.golang.org/grpc"
)

const (
	protocol = "unix"
	sockAddr = "/tmp/echo.sock"
)

func main() {
	// - https://qiita.com/marnie_ms4/items/4582a1a0db363fe246f3
	// - http://yamahiro0518.hatenablog.com/entry/2016/02/01/215908

	dialer := func(addr string, t time.Duration) (net.Conn, error) {
		return net.Dial(protocol, addr)
	}

	conn, err := grpc.Dial(sockAddr, grpc.WithInsecure(), grpc.WithDialer(dialer))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	var (
		client  = pb.NewEchoClient(conn)
		rootCtx = context.Background()
		timeout = 1 * time.Second
		values  = []string{
			"hello world",
			"golang",
			"goroutine",
			"this program runs on crostini",
		}
	)

	for _, v := range values {
		message := pb.EchoMessage{Data: v}

		ctx, cancel := context.WithTimeout(rootCtx, timeout)
		defer cancel()

		res, err := client.Echo(ctx, &message)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(res)
	}
}
