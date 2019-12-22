package main

import (
	"context"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/ivan-kostko/goq/service"
	pb "github.com/ivan-kostko/goq/service/proto/src/go/documents"
	"github.com/ivan-kostko/goq/storage"
	"github.com/ivan-kostko/goq/storage/inmem"
)

func ComposeInmemRpcDocumentService(log log.Logger) {

	inmemStorage := new(sync.Map)

	documentsServerLog := &contextLogger{log.WithField("name", "DocumentsServer")}

	storerLog := log.WithField("name", "Storer")
	inmemStorerLog := log.WithField("name", "InmemStorer")

	getterLog := log.WithField("name", "Storer")
	inmemGetterLog := log.WithField("name", "InmemStorer")

	grpcServer, err := grpcServer(false, "", "")
	if err != nil {
		log.Fatalf("Failed to set up gRPC server: %v", err)
	}

	pb.RegisterDocumentsServer(
		grpcServer,
		service.NewDocumentsServer(
			documentsServerLog,
			storage.NewDocumentStorer(storerLog, inmem.NewStorer(inmemStorerLog, inmemStorage)),
			storage.NewDocumentGetter(getterLog, inmem.NewGetter(inmemGetterLog, inmemStorage)),
		),
	)

	pb.RegisterDocumentsHandlerFromEndpoint(context.Background(), mux)

}

type contextLogger struct {
	*log.Entry
}

func (cl *contextLogger) SetContext(ctx context.Context) {
	cl.Entry = cl.Entry.WithContext(ctx)
}
