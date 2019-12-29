package main

import (
	"flag"
	"net"

	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

var (
	docRestAddr = flag.String("docRest", "127.0.0.1:9911", "The service address to listen for REST calls")
	docGrpcAddr = flag.String("docGrpc", "127.0.0.1:9912", "The service address to listen for gRPC calls")
	// webAppPort  = flag.String("web", "", "The web application UI port")
)

type Server interface{ Serve(net.Listener) error }

func main() {
	flag.Parse()

	log := logrus.New()

	helperGetListener := func(addr string) net.Listener {
		listener, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatalf("Failed to listen address %v due to error: %v", addr, err)
		}
		return listener
	}

	grpcListener := helperGetListener(*docGrpcAddr)
	restListener := helperGetListener(*docGrpcAddr)

	docGrpcServer, docRestServer, err := ComposeInmemRpcDocumentService(log, grpcListener.Addr().String())
	if err != nil {
		log.Fatalf("Failed to compose service due to error: %v", err)
	}

	g := new(errgroup.Group)

	g.Go(func() error {
		log.Infof("Starting to serve RPC requests on %v", grpcListener.Addr().String())
		return docGrpcServer.Serve(grpcListener)
	})

	g.Go(func() error {
		log.Infof("Starting to serve REST requests on %v", restListener.Addr().String())
		return docRestServer.Serve(restListener)
	})

	log.Fatal(g.Wait())

}
