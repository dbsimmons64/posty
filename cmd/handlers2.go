package main

import (
	"fmt"
	"net/http"
)

func (app *app) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}
