package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Spotify Social</h1>")
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	pathVars := mux.Vars(r)
	uri := pathVars["uri"]
	fmt.Fprintf(w, "<h1>Profile: %v</h1>", uri)
}

func bookmarksHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Bookmarks</h1>")
}

func exploreHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Explore</h1>")
}

func settingsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Settings</h1>")
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/bookmarks", bookmarksHandler)
	r.HandleFunc("/explore", exploreHandler)
	r.HandleFunc("/settings", settingsHandler)
	r.HandleFunc("/{uri}", profileHandler)
	http.Handle("/", r)

	port := ":8080"
	fmt.Println("Serving at port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
