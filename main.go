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
	// if current url path is not "/", then http.NotFound()
	// sends a 404 response to the client
	if r.URL.Path != "/"{
		http.NotFound(w, r)
		return 
	}

	w.Write([]byte("Hello from SnippetBox"))
}

func snippetView(w http.ResponseWriter, r*http.Request){
	w.Write([]byte("Display a specific snippet"))
}

func snippetCreate(w http.ResponseWriter, r*http.Request){
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// servemux stores a mapping b/w predefined	URL paths of app. 
	// and the corresponding handlers
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)	// home function is registered as the handler for "/" URL pattern
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting server on :4000")
	// ListenAndServe starts a new server
	err := http.ListenAndServe(":4000", mux)
	// if http.ListenAndServe returns an error, log.Fatal() is used
	// log.Fatal is used to log the error message and exit
	log.Fatal(err)
}