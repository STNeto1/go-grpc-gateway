package database

import (
	"__user/ent"
	"context"
	_ "github.com/lib/pq"
	"log"
)

func InitDB() *ent.Client {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=gw_users password=postgres sslmode=disable")
	if err != nil {
		log.Fatalln("error opening connection", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalln("feailed creating schema", err)
	}

	return client
}
