package main

import (
	"context"
	"log"
	"os"

	"github.com/gogo/protobuf/proto"
	"github.com/kchristidis/nested/lib/common"
	"github.com/kchristidis/nested/lib/extended"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var logger *zap.SugaredLogger

func main() {
	var err error

	// Setup logging
	if plainLogger, err := zap.NewDevelopment(); err != nil {
		log.Fatalf("Failed to set up logging: %s", err)
	} else {
		logger = plainLogger.Sugar()
	}

	// Parse command-line arguments
	args := os.Args[1:]

	// Launch client
	client, err := newClient(args[0])
	if err != nil {
		logger.Fatal(err)
	}

	// Create extended request
	extpld := extended.Payload{Common: &common.Payload{Value: 42}}
	if args[1] == "extended" { // See: https://golang.org/src/strings/compare.go
		extpld.Extra = "hello world"
	}
	var extpldmrs []byte
	if extpldmrs, err = proto.Marshal(&extpld); err != nil {
		logger.Fatal(errors.Wrap(err, "failed to pack extended payload"))
	}
	extreq := common.Request{Payload: extpldmrs}

	// Invoke GET
	logger.Infof("Invoking %s GET with 'common.value' set to '%d' and 'extra' set to '%s'", args[1], extpld.GetCommon().GetValue(), extpld.GetExtra())
	extresp, err := client.Get(context.Background(), &extreq)
	if err != nil {
		logger.Errorf("GET failed: %s", err)
	}
	logger.Infof("GET response: %s", extresp.String())
}

func newClient(serverAddress string) (extended.GetClient, error) {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial server")
	}
	return extended.NewGetClient(conn), nil
}
