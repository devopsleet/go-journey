package main

import (
	"fmt"
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slicing containing
// "Hello from Sanippetbox" as teh response body
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

// Add a snippetView handler function
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet"))

}

// Add a snippetCreate handler function
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating  a new snippet ...."))
}

func main() {
	fmt.Println("Hello World!")

	// Use the http.NewServeMux() function to initialize a new serveMux, the
	// register the home funcation as the handler for the "/" URL pattern
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// Print a log message to say that the server is starting
	log.Print("starting server on: 4000")

	// Start a new web server
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
