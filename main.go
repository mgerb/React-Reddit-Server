package main

import (
    "github.com/julienschmidt/httprouter"
    "net/http"
)

func init() {
	router := httprouter.New()
	router.ServeFiles("/dist/*filepath", http.Dir("./dist/"))
	router.NotFound = http.HandlerFunc(fileHandler("./dist/index.html"))
	http.Handle("/", router)
}

func fileHandler(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path)
	}
}