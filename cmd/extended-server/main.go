package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"go.uber.org/zap"

	"github.com/gogo/protobuf/proto"
	"github.com/kchristidis/nested/lib/common"
	"github.com/kchristidis/nested/lib/extended"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const (
	commonServerAddress   = "localhost:5051"
	extendedServerAddress = "localhost:5052"
)

var client common.GetClient
var logger *zap.SugaredLogger

func main() {
	// Setup logging
	if plainLogger, err := zap.NewDevelopment(); err != nil {
		log.Fatal(errors.Wrap(err, "failed to set up logging"))
	} else {
		logger = plainLogger.Sugar()
	}

	// Launch the common gRPC server
	go func() {
		logger.Infof("Setting up common server at %s", commonServerAddress)
		if err := common.NewServer(commonServerAddress); err != nil {
			logger.Fatal(errors.Wrap(err, "failed to set up common server"))
		}
	}()

	// Launch the common client
	timeout := time.Second * 3
	ticker := time.NewTicker(timeout / 6)
loop:
	for {
		var err error
		select {
		case <-ticker.C:
			if client, err = common.NewClient(commonServerAddress); err == nil {
				ticker.Stop()
				logger.Infof("Common client connected to %s", commonServerAddress)
				break loop
			}
		case <-time.After(timeout):
			ticker.Stop()
			logger.Fatal(errors.Wrap(err, "failed to set up common client"))
		}
	}

	// Launch the extended server
	if err := newServer(extendedServerAddress); err != nil {
		logger.Fatal(errors.Wrap(err, "extended server failed"))
	}
}

func newServer(serverAddress string) error {
	srv := grpc.NewServer()
	extended.RegisterGetServer(srv, &Server{})
	lis, err := net.Listen("tcp", serverAddress)
	if err != nil {
		return errors.Wrap(err, "failed to bind extended server to local address")
	}
	logger.Infof("Setting up extended server at %s", serverAddress)
	return srv.Serve(lis)
}

// Server implements the gRPC server interface.
type Server struct{}

// Get returns the value saved in the data store.
func (s *Server) Get(ctx context.Context, req *common.Request) (*extended.Response, error) {
	var err error

	// Unpack the request
	ep := extended.Payload{}
	if err = proto.Unmarshal(req.GetPayload(), &ep); err != nil {
		return nil, errors.Wrap(err, "failed to unpack request")
	}
	logger.Debugf("Received GET request with 'common.value' set to %d and 'extra' set to '%s'", ep.GetCommon().GetValue(), ep.GetExtra())

	// Use common client
	var cpm []byte // See: https://github.com/golang/go/wiki/CodeReviewComments#declaring-empty-slices
	if cpm, err = proto.Marshal(ep.GetCommon()); err != nil {
		return nil, errors.Wrap(err, "failed to pack request")
	}
	cr := new(common.Response)
	if cr, err = client.Get(context.Background(), &common.Request{Payload: cpm}); err != nil {
		return nil, errors.Wrap(err, "failed on common get")
	}

	// Apply the transformation present in the extended payload
	if len(ep.GetExtra()) == 0 { // Do nothing
		return &extended.Response{Type: &extended.Response_Common{Common: cr}}, nil
	}
	// Dummy way to show how the filter transformation would work
	return &extended.Response{Type: &extended.Response_Extra{Extra: fmt.Sprintf("%d and %s", cr.GetValue(), ep.GetExtra())}}, nil
}
