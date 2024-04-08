package utils

import (
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

func CreateESClient() (esClient *elasticsearch.Client, err error) {

	clientES, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Printf("Error creating the client: %s", err)
	} else {
		log.Println(clientES.Info())
	}

	return clientES, err
}
