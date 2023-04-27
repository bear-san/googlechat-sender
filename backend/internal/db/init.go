package db

import (
	"context"
	"fmt"
	"github.com/bear-san/googlechat-sender/backend/ent"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var Client *ent.Client

func init() {
	var err error
	Client, err = ent.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			"5432",
			os.Getenv("DB_USER"),
			"google-chat-sender",
			os.Getenv("DB_PASSWORD"),
		),
	)

	if err != nil {
		panic(err)
	}

	if err := Client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
