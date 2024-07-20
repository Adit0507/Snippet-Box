package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// file server serves files out of the "./ui/static/" directory.
	fileServer := http.FileServer(http.Dir("./ui/static"))

	// mux.Handle() registers the file server as the file handler for all
	// URL paths that start with "/static". For matching paths, we strip
	// "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting server on 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}