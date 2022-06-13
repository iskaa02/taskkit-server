package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/iskaa02/taskkit-server/api"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/test?sslmode=disable")
	if err != nil {
		log.Fatal("cannot open db", err)
	}
	s := api.NewServer(db)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
