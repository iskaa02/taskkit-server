package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/iskaa02/taskkit-server/api"
	"github.com/iskaa02/taskkit-server/db/models"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:password@localhost:5432?sslmode=disable")
	if err != nil {
		log.Fatal("cannot open db", err)
	}
	_, err = models.Prepare(context.Background(), db)
	if err != nil {
		log.Fatal(err)
	}
	s := api.NewServer(db)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
