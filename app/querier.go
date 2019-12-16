package app

import (
	"bytes"
	"fmt"

	"github.com/PuerkitoBio/goquery"

	"github.com/pkg/errors"
)

type QueryResult []*Document

type Querier func(string) (QueryResult, error)

type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Error(args ...interface{})
}

func NewQuerier(d *Document, log Logger) Querier {

	doc, parceError := goquery.NewDocumentFromReader(d.reader)
	parceError = errors.Wrapf(parceError, "Failed to parce document '%s'", d.GetId())
	if parceError != nil {
		log.Error(parceError)
	}

	return func(query string) (QueryResult, error) {
		log.Infof("Running query '%s' on '%s' document", query, d.GetId())
		if parceError != nil {
			return nil, parceError
		}

		result := make([]*Document, 0)

		doc.Find(query).Each(func(i int, s *goquery.Selection) {
			result = append(result, NewDocument(fmt.Sprintf("%s/%s/%v", d.GetId(), query, i), bytes.NewBufferString(s.Text())))
		})

		return result, nil
	}
}
