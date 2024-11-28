package web

import (
	"github.com/dbsimmons64/cmd/web/"
	"net/http"
)

func (app *app) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", app.Home)
	mux.HandleFunc("GET /posts/create", app.createPost)
	mux.HandleFunc("POST /posts/create", app.storePost)

	return mux
}
