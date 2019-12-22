package main

import (
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"context"

	pb "github.com/ivan-kostko/goq/service/proto/src/go/documents"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

var (
	serviceRpcPort  = flag.String("docGrpc", "", "The service port to listen for RPC calls")
	serviceRestPort = flag.String("docRest", "", "The service port to listen for REST calls")
	webAppPort      = flag.String("web", "", "The web application UI port")
)

type Server interface{ Serve(net.Listener) error }

func main() {
	flag.Parse()

	log := logrus.New()

	mux := runtime.

	docGrpcServer := ComposeInmemRpcDocumentService(log)
	docRestService := pb.RegisterDocumentsHandlerFromEndpoint(context.Background())

}
