package main

import (
	"database/sql"
	"log"
	"net/http"

	repositories "github.com/dbsimmons64/posty/repos"
	"github.com/dbsimmons64/posty/services"
	_ "github.com/mattn/go-sqlite3"
)

type app struct {
	postService services.PostService
}

func main() {

	// Connect to the SQLITE3 database
	db, err := sql.Open("sqlite3", "database/app.db")
	if err != nil {
		log.Fatal(err)
	}

	// Ensure we close any database connection
	// Probably not needed for sqlite.
	defer db.Close()

	repo := repositories.PostRepositoryDB{DB: db}
	service := services.PostServiceImpl{Repo: repo}
	app := app{postService: service}

	srv := http.Server{
		Addr:    ":8080",
		Handler: app.routes(),
	}

	log.Println("Listening on port 8080")
	srv.ListenAndServe()
}
