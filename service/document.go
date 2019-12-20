package service

import (
	"github.com/ivan-kostko/goq/app"
)

type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Error(args ...interface{})
}

type documentService struct {
	store func(d *app.Document) error
	get   func(identifier string) (*app.Document, error)
}

func (ds *documentService) Store(d *app.Document) error {
	return ds.store(d)
}

func (ds *documentService) Get(identifier string) (*app.Document, error) {
	return ds.get(identifier)
}

func NewDocumentService(store func(d *app.Document) error, get func(identifier string) (*app.Document, error)) *documentService {
	return &documentService{
		get:   get,
		store: store,
	}
}
