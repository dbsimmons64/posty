package main

import (
	"net/http"
)

func (app *app) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", app.Home)
	mux.HandleFunc("GET /posts/create", app.createPost)
	mux.HandleFunc("POST /posts/create", app.storePost)
	mux.HandleFunc("GET /hello", app.Hello)

	return mux
}
