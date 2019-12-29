package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"

	"sync"

	"github.com/pkg/errors"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"google.golang.org/grpc"

	log "github.com/sirupsen/logrus"

	"github.com/ivan-kostko/goq/service"
	pb "github.com/ivan-kostko/goq/service/proto/src/go/documents"
	"github.com/ivan-kostko/goq/storage"
	"github.com/ivan-kostko/goq/storage/inmem"
)

func ComposeInmemRpcDocumentService(log *log.Logger, rpcAddr string) (Server, Server, error) {

	grpcServer := grpc.NewServer()

	inmemStorage := new(sync.Map)

	documentsServerLog := &contextLogger{log.WithField("name", "DocumentsServer")}

	storerLog := log.WithField("name", "Storer")
	inmemStorerLog := log.WithField("name", "InmemStorer")

	getterLog := log.WithField("name", "Storer")
	inmemGetterLog := log.WithField("name", "InmemStorer")

	pb.RegisterDocumentsServer(
		grpcServer,
		service.NewDocumentsServer(
			documentsServerLog,
			storage.NewDocumentStorer(storerLog, inmem.NewStorer(inmemStorerLog, inmemStorage)),
			storage.NewDocumentGetter(getterLog, inmem.NewGetter(inmemGetterLog, inmemStorage)),
		),
	)

	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &rawJSONPb{runtime.JSONPb{}}),
		runtime.WithProtoErrorHandler(runtime.DefaultHTTPProtoErrorHandler),
	)

	if err := pb.RegisterDocumentsHandlerFromEndpoint(context.Background(), mux, rpcAddr, []grpc.DialOption{grpc.WithInsecure()}); err != nil {
		err = errors.Wrap(err, "Failed to register REST service")
		return nil, nil, err
	}

	return grpcServer, &http.Server{Handler: mux}, nil
}

type contextLogger struct {
	*log.Entry
}

func (cl *contextLogger) SetContext(ctx context.Context) {
	cl.Entry = cl.Entry.WithContext(ctx)
}

var typeOfBytes = reflect.TypeOf([]byte(nil))

type rawJSONPb struct {
	runtime.JSONPb
}

func (*rawJSONPb) NewDecoder(r io.Reader) runtime.Decoder {
	return runtime.DecoderFunc(func(v interface{}) error {
		rawData, err := ioutil.ReadAll(r)
		if err != nil {
			return err
		}
		rv := reflect.ValueOf(v)

		if rv.Kind() != reflect.Ptr {
			return fmt.Errorf("%T is not a pointer", v)
		}

		rv = rv.Elem()
		if rv.Type() != typeOfBytes {
			return fmt.Errorf("Type must be []byte but got %T", v)
		}

		rv.Set(reflect.ValueOf(rawData))
		return nil
	})
}
