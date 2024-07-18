package main

import (
	"log"
	"net/http"
)

// this function writes a byte slice containing
// "Hello from SnippetBox" as the response body
// http.ResponseWriter provides methods for assembling a HTTP response and sending to the user,
// http.Request parameter is a pointer to a struct which holds information about the current request 
func home(w http.ResponseWriter, r*http.Request) {
	w.Write([]byte("Hello from SnippetBox"))
}

func main() {
	// servemux stores a mapping b/w predefined	URL paths of app. 
	// and the corresponding handlers
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)	// home function is registered as the handler for "/" URL pattern

	log.Print("Starting server on :4000")
	// ListenAndServe starts a new server
	err := http.ListenAndServe(":4000", mux)
	// if http.ListenAndServe returns an error, log.Fatal() is used
	// log.Fatal is used to log the error message and exit
	log.Fatal(err)
}