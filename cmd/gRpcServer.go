package cmd

import (
	"google.golang.org/grpc"
	
)

func grpcServer(tls bool, certFile, keyFile string) *grpc.Server {
	var tlsOpts []grpc.ServerOption
	if tls {
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("Failed to setup TLS due to error: %v", err)
		}
		tlsOpts = []grpc.ServerOption{grpc.Creds(creds)}
	}

	return grpc.NewServer(tlsOpts...)

}