package inmem

import (
	"bytes"
	"io"
	"sync"

	"github.com/pkg/errors"
)

var (
	globalCache = new(sync.Map)
)

type Storer func(identifier string, content []byte) error

type Getter func(identifier string) (io.Reader, error)

type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Error(args ...interface{})
}

func NewStorer(log Logger, storage interface{ Store(key, value interface{}) }) Storer {
	return func(identifier string, content []byte) error {
		log.Infof("Storing '%s'", identifier)

		storage.Store(identifier, content)

		return nil
	}
}

func NewGetter(log Logger, storage interface {
	Load(key interface{}) (value interface{}, ok bool)
}) Getter {
	return func(identifier string) (io.Reader, error) {
		log.Infof("Getting '%s'", identifier)

		content, ok := storage.Load(identifier)
		if content == nil || !ok {
			return nil, errors.New("No content found")
		}

		log.Debugf("Casting content '%s' to bytes", identifier)
		if x, ok := content.([]byte); ok {
			return bytes.NewReader(x), nil
		}

		return nil, errors.New("Failed to cast content")
	}
}

func (f Storer) Store(identifier string, content []byte) error {
	return f(identifier, content)
}

func (f Getter) Get(identifier string) (io.Reader, error) {
	return f(identifier)
}
