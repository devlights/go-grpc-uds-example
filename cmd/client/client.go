package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/devlights/go-grpc-uds-example/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	protocol = "unix"
	sockAddr = "/tmp/echo.sock"
)

func main() {
	// - https://qiita.com/marnie_ms4/items/4582a1a0db363fe246f3
	// - http://yamahiro0518.hatenablog.com/entry/2016/02/01/215908
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
	)
	defer mainCxl()

	dialer := func(ctx context.Context, addr string) (net.Conn, error) {
		return net.Dial(protocol, addr)
	}

	conn, err := grpc.Dial(sockAddr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithContextDialer(dialer))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	var (
		client = pb.NewEchoClient(conn)
		values = []string{
			"hello world",
			"golang",
			"goroutine",
			"this program is used gRPC with golang",
		}
	)

	for _, v := range values {
		func() {
			ctx, cancel := context.WithTimeout(mainCtx, 1*time.Second)
			defer cancel()

			message := pb.EchoMessage{Data: v}

			res, err := client.Echo(ctx, &message)
			if err != nil {
				log.Println(err)
				return
			}

			fmt.Println(res)
		}()
	}
}
