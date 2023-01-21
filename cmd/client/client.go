// gRPC with Unix Domain Socket example (client side)
//
// # REFERENCES
// 	- https://qiita.com/marnie_ms4/items/4582a1a0db363fe246f3
// 	- http://yamahiro0518.hatenablog.com/entry/2016/02/01/215908
// 	- https://zenn.dev/hsaki/books/golang-grpc-starting/viewer/client
// 	- https://stackoverflow.com/a/46279623
// 	- https://stackoverflow.com/a/18479916

package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
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
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
	)
	defer mainCxl()

	//
	// Connect
	//
	var (
		credentials = insecure.NewCredentials() // No SSL/TLS
		dialer      = func(ctx context.Context, addr string) (net.Conn, error) {
			var d net.Dialer
			return d.DialContext(ctx, protocol, addr)
		}
		options = []grpc.DialOption{
			grpc.WithTransportCredentials(credentials),
			grpc.WithBlock(),
			grpc.WithContextDialer(dialer),
		}
	)

	conn, err := grpc.Dial(sockAddr, options...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//
	// Send & Recv
	//
	var (
		client = pb.NewEchoClient(conn)
		file   = func() *os.File {
			f, err := os.Open("text-files/agatha_complete.txt")
			if err != nil {
				panic(err)
			}
			return f
		}()
		values = func(r io.ReadCloser) <-chan string {
			ch := make(chan string)
			go func() {
				defer r.Close()
				defer close(ch)
				scanner := bufio.NewScanner(r)
				for scanner.Scan() {
					ch <- scanner.Text()
				}
			}()
			return ch
		}(file)
	)

	for v := range values {
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
