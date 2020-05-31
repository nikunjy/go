package server

import (
	"context"
	"fmt"

	pb "github.com/nikunjy/go/protos/hello"
)

type Server struct {
}

func (h *Server) Greet(ctx context.Context, req *pb.GreetingRequest) (*pb.GreetingResponse, error) {
	return &pb.GreetingResponse{
		Greeting: fmt.Sprintf("Hello %s", req.GetName()),
	}, nil
}
