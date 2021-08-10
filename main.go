package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	// .env
	"os"

	"github.com/joho/godotenv"

	// db

	"github.com/alecchendev/go-spotify-social/db"
)

var templateDir = "templates/"
var templateFilenames = []string{
	templateDir + "index.html",
	templateDir + "bookmarks.html",
	templateDir + "explore.html",
	templateDir + "settings.html",
	templateDir + "profile.html",
}

var templates = template.Must(template.ParseFiles(templateFilenames...))

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := homeData{
		Title: "Spotify Social",
	}
	renderTemplate(w, "index", data)
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	pathVars := mux.Vars(r)
	uri := pathVars["uri"]

	// Retrieve profile data
	// ...
	data := profileData{
		Uri: uri,
	}
	renderTemplate(w, "profile", data)
}

func bookmarksHandler(w http.ResponseWriter, r *http.Request) {
	data := bookmarksData{}
	renderTemplate(w, "bookmarks", data)
}

func exploreHandler(w http.ResponseWriter, r *http.Request) {
	data := exploreData{}
	renderTemplate(w, "explore", data)
}

func settingsHandler(w http.ResponseWriter, r *http.Request) {
	data := settingsData{}
	renderTemplate(w, "settings", data)
}

func main() {

	// Connect to db
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file.")
	}
	dbUri := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	collectionName := os.Getenv("COLLECTION_NAME")
	client, ctx := db.InitializeDBClient(dbUri, dbName, collectionName)
	defer client.Disconnect(ctx)

	// Serve pages
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/bookmarks", bookmarksHandler)
	r.HandleFunc("/explore", exploreHandler)
	r.HandleFunc("/settings", settingsHandler)
	r.HandleFunc("/{uri}", profileHandler)
	http.Handle("/", r)

	// // Start server
	port := ":8080"
	fmt.Println("Serving at port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
