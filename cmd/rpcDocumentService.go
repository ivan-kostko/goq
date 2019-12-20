package cmd

import (
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/ivan-kostko/goq/service"
	pb "github.com/ivan-kostko/goq/service/proto/src/go/documents"
	"github.com/ivan-kostko/goq/storage"
	"github.com/ivan-kostko/goq/storage/inmem"
)

func ComposeInmemRpcDocumentService(log log.Logger) {

	inmemStorage := new(sync.Map)

	storerLog := log.WithField("name", "Storer")
	inmemStorerLog := log.WithField("name", "InmemStorer")

	getterLog := log.WithField("name", "Storer")
	inmemGetterLog := log.WithField("name", "InmemStorer")

	pb.RegisterDocumentsServer(
		grpcServer,
		service.NewDocumentService(
			storage.NewDocumentStorer(storerLog, inmem.NewStorer(inmemStorerLog, inmemStorage)),
			storage.NewDocumentGetter(getterLog, inmem.NewGetter(inmemGetterLog, inmemStorage)),
		),
	)

}
