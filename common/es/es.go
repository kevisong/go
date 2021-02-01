package es

import (
	"context"
	"fmt"

	"github.com/olivere/elastic"
	log "github.com/sirupsen/logrus"
)

// Document interface
type Document interface {
	Index() string
	Type() string
	ID() string
}

var defaultClient *elastic.Client

// Init initializes default client
func Init(hosts []string, enableAuth bool, username, password string) (err error) {
	if enableAuth {
		defaultClient, err = elastic.NewClient(
			elastic.SetBasicAuth(username, password),
			elastic.SetSniff(false),
			elastic.SetURL(hosts...),
		)
	} else {
		defaultClient, err = elastic.NewClient(
			elastic.SetSniff(false),
			elastic.SetURL(hosts...),
		)
	}
	if err != nil {
		return err
	}
	return nil
}

func checkExists(index string) (bool, error) {
	exists, err := defaultClient.IndexExists(index).Do(context.Background())
	if err != nil {
		return false, err
	}
	return exists, nil
}

// UpdateSettings update index settings
func UpdateSettings(index string, setting map[string]interface{}) (uIndex *elastic.IndicesPutSettingsResponse, err error) {

	// check if index exists
	exists, err := checkExists(index)
	if err != nil {
		return nil, err
	}

	if exists {
		uIndex, err = defaultClient.IndexPutSettings(index).BodyJson(setting).Do(context.Background())
		if err != nil {
			return nil, err
		}
	}

	return uIndex, nil

}

// CheckAndCreate creates index if not exists
func CheckAndCreate(index string, bodyString string) (cIndex *elastic.IndicesCreateResult, err error) {

	// check if index exists
	exists, err := checkExists(index)
	if err != nil {
		return nil, err
	}

	if !exists {
		cIndex, err = defaultClient.CreateIndex(index).BodyString(bodyString).Do(context.Background())
		if err != nil {
			return nil, err
		}
	}

	return cIndex, nil

}

// CheckAndDelete deletes index if exists
func CheckAndDelete(index string) (dIndex *elastic.IndicesDeleteResponse, err error) {

	// check if index exists
	exists, err := checkExists(index)
	if err != nil {
		return nil, err
	}

	if exists {
		dIndex, err = defaultClient.DeleteIndex(index).Do(context.Background())
		if err != nil {
			return nil, err
		}
	}

	return dIndex, nil

}

func index(index string, docs []Document, count int) {
	bulkRequest := defaultClient.Bulk()
	for i := 0; i < count; i++ {
		bulkRequest = bulkRequest.Add(
			elastic.NewBulkIndexRequest().Index(index).Type(docs[i].Type()).Id(docs[i].ID()).Doc(docs[i]),
		)
	}
	if bulkRequest.NumberOfActions() != 0 {
		resp, err := bulkRequest.Do(context.Background())
		if err != nil {
			log.Error(err)
		}
		if resp != nil && resp.Errors {
			for _, failed := range resp.Failed() {
				log.Error(failed.Error)
			}
		}
	}
	bulkRequest.Reset()
}

// Index index documents to elasticsearch
func Index(idx string, docs []Document) error {

	// check if index exists
	exists, err := checkExists(idx)
	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("index %s doesn't exists", idx)
	}

	for len(docs) > 0 {
		if len(docs) >= 10000 {
			index(idx, docs, 10000)
			docs = docs[10000:]
		} else {
			index(idx, docs, len(docs))
		}
	}

	return nil

}
