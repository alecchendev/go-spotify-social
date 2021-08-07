package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Spotify Social</h1>")
	})
	port := ":8080"
	fmt.Println("Serving at port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
