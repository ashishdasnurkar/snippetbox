package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello form Snippetbox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Showing a snippet..."))
}

func createNewSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating a new snippet..."))
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/showsnippet", showSnippet)
	mux.HandleFunc("/createnewsnippet", createNewSnippet)
	log.Println("Starting server on port 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}