package main

import (
	"net/http"
	"text/template"
)

func (app *app) Home(w http.ResponseWriter, r *http.Request) {

	posts, err := app.postService.ListPosts()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Be aware this will be parse the file on every request - there is a better way!!!
	t, err := template.ParseFiles("./assets/templates/home.page.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	t.Execute(w, map[string]any{"Posts": posts})
}

func (app *app) createPost(w http.ResponseWriter, r *http.Request) {

	// Be aware this will be parse the file on every request - there is a better way!!!
	t, err := template.ParseFiles("./assets/templates/post.create.page.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	t.Execute(w, nil)
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
