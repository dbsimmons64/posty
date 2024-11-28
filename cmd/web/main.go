package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/dbsimmons64/posty/internal/models/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

type app struct {
	posts *sqlite.PostModel
}

func main() {
	// Connect to the SQLITE3 database
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}

	// Ensure we close any database connection
	// Probably not needed for sqlite.
	defer db.Close()

	app := app{
		posts: &sqlite.PostModel{
			DB: db,
		},
	}

	srv := http.Server{
		Addr:    ":8080",
		Handler: app.routes(),
	}

	log.Println("Listening on port 8080")
	srv.ListenAndServe()
}
