package main

import (
	"log"

	_ "github.com/lib/pq"

	"github.com/iskaa02/taskkit-server/api"
	"github.com/iskaa02/taskkit-server/ent"
)

func main() {
	client, err := ent.Open("postgres", "postgres://postgres:password@localhost:5432?sslmode=disable")
	if err != nil {
		log.Fatal("cannot open db", err)
	}
	if err != nil {
		log.Fatal(err)
	}
	s := api.NewServer(client)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
