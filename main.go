package main

import (
    "github.com/julienschmidt/httprouter"
    "net/http"
)

func init() {
	r := httprouter.New()
	r.ServeFiles("/*filepath", http.Dir("./dist/"))
	r.NotFound = http.HandlerFunc(fileHandler("./dist/index.html"))

	http.Handle("/", r)
}

func fileHandler(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path)
	}
}