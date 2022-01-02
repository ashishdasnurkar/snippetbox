package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	// without the exact path match, "/" is a subtree pattern and acts as a catch-all
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello form Snippetbox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", 405)
		return
	}
	// retrieve id from query string for eg. /snippet?id=5
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	//w.Write([]byte("Showing a snippet..."))
	fmt.Fprintf(w, "Dispaying a specific snippet with ID: %d ...", id)
}

func createNewSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(405) // must be called before w.Write
		w.Write([]byte("Method not allowed")) // w.Write sends 200 status code
		return
	}
	w.Write([]byte("Creating a new snippet..."))
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createNewSnippet)
	log.Println("Starting server on port 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}