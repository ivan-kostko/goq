package service

import (
	"bytes"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"

	"context"

	"github.com/ivan-kostko/goq/app"
	pb "github.com/ivan-kostko/goq/service/proto/src/go/documents"
)

type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Error(args ...interface{})
	SetContext(ctx context.Context)
}

type documentsServer struct {
	store  func(d *app.Document) error
	get    func(identifier string) (*app.Document, error)
	logger Logger
}

func (ds *documentsServer) Store(ctx context.Context, req *pb.DocumentStoreReuest) (*pb.Empty, error) {

	log := ds.logger
	log.SetContext(ctx)

	log.Infof("Serving store document request")
	log.Debugf("Requested '%#v'", req)

	if req != nil {
		log.Debugf("Request is not nil")
		if d := req.GetDocument(); d != nil {
			log.Debugf("Requested valid document. Mapping and storing")

			err := errors.Wrapf(ds.store(app.NewDocument(d.GetIdentifier(), bytes.NewBufferString(d.GetContent()))), "Failed to store document")
			if err != nil {
				log.Error(err)
			}
			return nil, err
		}

	}

	return nil, errors.New("Invalid request")
}

func (ds *documentsServer) Get(ctx context.Context, req *pb.DocumentGetReuest) (resp *pb.DocumentGetResponse, e error) {

	log := ds.logger
	log.SetContext(ctx)

	log.Infof("Serving get document request")
	log.Debugf("Requested '%#v'", req)

	if req != nil {
		log.Debugf("Request is not nil")
		if doc, err := ds.get(req.GetIdentifier()); err == nil {
			log.Debugf("Internal document found. Mapping")

			resp.Document = &pb.Document{
				Identifier: doc.GetId(),
				Content:    string(doc.Content()),
			}

			log.Debugf("Returning mapped document %#v", resp)
			return resp, nil
		}
	}
	return resp, errors.New("Invalid request '" + req.String() + "'")
}

func NewDocumentsServer(log Logger, store func(d *app.Document) error, get func(identifier string) (*app.Document, error)) *documentsServer {

	return &documentsServer{
		get:    get,
		store:  store,
		logger: log,
	}
}
