package app

import (
	"io"
	"io/ioutil"

	"github.com/google/uuid"
)

// Represents application document
type Document struct {
	reader  io.Reader
	id      string
	content []byte
}

// Creates a new document
func NewDocument(id string, r io.Reader) *Document {
	doc := &Document{
		id:     id,
		reader: r,
	}

	// TODO: find a better idea
	doc.content, _ = ioutil.ReadAll(r)

	if doc.id == "" {
		doc.id = getNewId()
	}

	return doc
}

func (d *Document) GetId() string {
	return d.id
}

func (d *Document) Content() []byte {
	return d.content
}

func getNewId() string {
	return uuid.New().String()
}
