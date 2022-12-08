package database

import (
	"__product/ent"
	"context"
	_ "github.com/lib/pq"
	"log"
)

func InitDB() *ent.Client {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=gw_products password=postgres sslmode=disable")
	if err != nil {
		log.Fatalln("error opening connection", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalln("failed creating schema", err)
	}

	return client
}
