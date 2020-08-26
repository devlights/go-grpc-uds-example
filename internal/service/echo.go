package service

import (
	"context"
	"strings"

	"github.com/devlights/go-grpc-uds-example/internal/pb"
)

type EchoServiceImpl struct {
}

var _ pb.EchoServer = (*EchoServiceImpl)(nil)

func NewEchoService() pb.EchoServer {
	return new(EchoServiceImpl)
}

func (e *EchoServiceImpl) Echo(ctx context.Context, message *pb.EchoMessage) (*pb.EchoResponse, error) {
	s := strings.ToUpper(message.Data)
	r := &pb.EchoResponse{
		Data: s,
	}

	return r, nil
}
