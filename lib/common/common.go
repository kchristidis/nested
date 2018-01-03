package common

import (
	"context"
	"net"

	"github.com/golang/protobuf/proto"

	"google.golang.org/grpc"
)

// NewClient returns a new client.
func NewClient(serverAddress string) (GetClient, error) {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return NewGetClient(conn), nil
}

// NewServer spawns a new server.
func NewServer(serverAddress string) error {
	srv := grpc.NewServer()
	RegisterGetServer(srv, &Server{})
	lis, err := net.Listen("tcp", serverAddress)
	if err != nil {
		return err
	}
	return srv.Serve(lis)
}

// Server implements the gRPC server interface.
type Server struct{}

// Get parses the request and returns the appropriate value.
func (s *Server) Get(ctx context.Context, req *Request) (*Response, error) {
	in := Payload{}
	if err := proto.Unmarshal(req.GetPayload(), &in); err != nil {
		return nil, err
	}
	resp := &Response{Value: in.GetValue()}
	return resp, nil
}
