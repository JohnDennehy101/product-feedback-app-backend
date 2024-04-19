package main

import "net/http"

func (app *application) root(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Hello World..."))
}
