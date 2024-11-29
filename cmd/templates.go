package main

import (
	"bytes"
	"html/template"
	"net/http"
)

type pageData map[string]any

func render(w http.ResponseWriter, path string, data pageData) {
	t, err := template.ParseFiles(path, "./assets/templates/main.layout.html")

	if err != nil {
		http.Error(w, "Cannot load page", 500)
		return
	}

	buffer := new(bytes.Buffer)
	err = t.Execute(buffer, data)
	if err != nil {
		http.Error(w, "Cannot load page", 500)
		return
	}

	buffer.WriteTo(w)

}
