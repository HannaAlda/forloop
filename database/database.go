package database

import (
	"context"
	"flag"

	"github.com/tikv/client-go/v2/rawkv"
)

func Connect() (*rawkv.Client, error) {
	addres := flag.String("pd", "127.0.0.1:2379", "pd address")
	// client, err := txnkv.NewClient([]string{*addres})
	// if err != nil {
	// 	return nil, err
	// }

	client, err := rawkv.NewClientWithOpts(context.TODO(), []string{*addres})
	if err != nil {
		return nil, err
	}

	return client, nil

}
