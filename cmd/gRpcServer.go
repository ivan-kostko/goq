package main

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func grpcServer(tls bool, certFile, keyFile string) (*grpc.Server, error) {
	var tlsOpts []grpc.ServerOption
	if tls {
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to setup TLS")
		}
		tlsOpts = []grpc.ServerOption{grpc.Creds(creds)}
	}

	return grpc.NewServer(tlsOpts...), nil

}
