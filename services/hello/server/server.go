package server

import (
	"context"
	"fmt"

	"github.com/nikunjy/go/protos/common"
	pb "github.com/nikunjy/go/protos/hello"
)

type Server struct {
}

func (h *Server) Greet(ctx context.Context, req *pb.GreetingRequest) (*pb.GreetingResponse, error) {
	return &pb.GreetingResponse{
		Greeting: fmt.Sprintf("Hello %s", req.GetName()),
	}, nil
}

func (h *Server) Bogus(ctx context.Context, req *common.Bogus) (*pb.GreetingResponse, error) {
	return &pb.GreetingResponse{
		Greeting: fmt.Sprintf("Hello %s", "Bogus"),
	}, nil
}
