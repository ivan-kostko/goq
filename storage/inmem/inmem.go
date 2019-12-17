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

func NewStorer(log Logger) Storer {
	return func(identifier string, content []byte) error {
		log.Infof("Storing '%s'", identifier)

		globalCache.Store(identifier, content)

		return nil
	}
}

func NewGetter(log Logger) Getter {
	return func(identifier string) (io.Reader, error) {
		log.Infof("Getting '%s'", identifier)

		content, ok := globalCache.Load(identifier)
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
