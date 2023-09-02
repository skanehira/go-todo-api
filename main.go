package main

import (
	"context"
	"log"
	"net/http"
	"todo/ent"
	"todo/ent/ogent"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open(dialect.SQLite, "file:todo.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	if err := client.Schema.Create(context.Background(), schema.WithAtlas(true)); err != nil {
		log.Fatal(err)
	}

	h := handler{
		OgentHandler: ogent.NewOgentHandler(client),
		client:       client,
	}

	server, err := ogent.NewServer(h)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("listening on port %s", ":8080")
	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatal(err)
	}
}
