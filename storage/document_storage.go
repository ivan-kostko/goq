package storage

import (
	"io"

	"github.com/ivan-kostko/goq/app"
	"github.com/pkg/errors"
)

type DocumentStorer func(d *app.Document) error

type DocumentGetter func(identifier string) (*app.Document, error)

type Storer interface {
	Store(identifier string, content []byte) error
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
		err = errors.Wrapf(err, "Failed to extract document '%s'", identifier)
		if err != nil {
			return nil, err
		}

		log.Debugf("Creating new document from storage response '%#v', '%#v'", identifier, reader)
		return app.NewDocument(identifier, reader), nil
	}
}

func NewDocumentStorer(log Logger, storage Storer) DocumentStorer {
	return func(d *app.Document) error {

		log.Infof("Persisting document '%s'", d.GetId())

		log.Debugf("Persisting document '%#v'", d)
		err := storage.Store(d.GetId(), d.Content())
		err = errors.Wrapf(err, "Failed to persist document '%s'", d.GetId())
		if err != nil {
			return err
		}

		log.Debugf("Successfuly stored document '%s'", d.GetId())
		return nil
	}
}
