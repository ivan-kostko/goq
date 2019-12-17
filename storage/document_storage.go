package storage

import (
	"io"
	"github.com/pkg/errors"
	"github.com/ivan-kostko/goq/app"
)

type DocumentStorer func(d *app.Document) error

type DocumentGetter func(identifier string) (*app.Document, error)

type Storer interface {
	Store(identifier string, []byte) error
}

type Getter interface {
	Get(identifier string) (io.Reader, error)
}

type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Error(args ...interface{})
}

func NewDocumentGetter(log Logger, storage Getter) DocumentGetter {
	return func(identifier string) (*app.Document, error) {
		log.Infof("Extracting document '%s'", identifier)

		reader, err := storage.Get(identifier)
		err = errors.Wrapf("Failed to extract document '%s'", identifier)
		if err != nil {
			return nil, err
		}

		log.Debugf("Creating new document from storage response")
		return app.NewDocument(identifier, reader), nil
	}
}

func NewDocumentStorer(log Logger, storage Storer) DocumentStorer {
	return func(d *app.Document) error {

		log.Infof("Persisting document '%s'", d.GetId())

		err := storage.Store(d.GetId(), d.Content())
		err = errors.Wrapf("Failed to persist document '%s'", d.GetId())
		if err != nil {
			return err
		}

		log.Debugf("Successfuly stored document '%s'", d.GetId())
		return nil
	}
}

