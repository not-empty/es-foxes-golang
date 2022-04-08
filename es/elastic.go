package es

import (
	"github.com/olivere/elastic/v7"
)

func GetESClient(url string) (*elastic.Client, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(false),
	)

	if err != nil {
		return nil, err
	}

	return client, nil
}
