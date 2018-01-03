package common

import (
	"context"
	"fmt"
	"net"

	"github.com/golang/protobuf/proto"

	"google.golang.org/grpc"
)

const srvAddr = "localhost:5051"

// NewClient returns a new client.
func NewClient() GetClient {
	conn, err := grpc.Dial(srvAddr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return NewGetClient(conn)
}

// NewServer spawns a new server.
func NewServer() error {
	srv := grpc.NewServer()
	RegisterGetServer(srv, &Server{})
	lis, err := net.Listen("tcp", srvAddr)
	if err != nil {
		panic(err)
	}
	return srv.Serve(lis)
}

// Server implements the gRPC server interface.
type Server struct{}

// Get parses the request and returns the appropriate value.
func (s *Server) Get(ctx context.Context, req *Request) (*Response, error) {
	p := Payload{}
	if err := proto.Unmarshal(req.GetPayload(), &p); err != nil {
		panic(err)
	}

	r := &Response{Type: &Response_Bar{Bar: fmt.Sprintf("you asked for %d", p.GetStartFrom())}}
	return r, nil
}
