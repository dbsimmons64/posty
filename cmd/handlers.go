package main

import (
	"net/http"
)

func (app *app) Home(w http.ResponseWriter, r *http.Request) {

	posts, err := app.postService.ListPosts()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	render(w, "./assets/templates/home.page.html", pageData{"Posts": posts})
}

func (app *app) createPost(w http.ResponseWriter, r *http.Request) {

	render(w, "./assets/templates/post.create.page.html", nil)
}

func (app *app) storePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	err = app.postService.CreatePost(
		r.PostForm.Get("title"),
		r.PostForm.Get("content"),
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
