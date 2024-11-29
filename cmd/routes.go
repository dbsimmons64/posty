package main

import (
	"net/http"
)

func (app *app) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", app.Home)
	mux.HandleFunc("GET /posts/create", app.createPost)
	mux.HandleFunc("POST /posts/create", app.storePost)

	// Serve static files
	fileserver := http.FileServer(http.Dir("./assets/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileserver))

	return mux
}
