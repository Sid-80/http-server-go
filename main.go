package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Content     string `json:"Content"`
}

type Articles []Article

func getAllArticles(w http.ResponseWriter, r *http.Request) {

	articles := Articles{
		Article{Title: "Hello", Description: "Description", Content: "Content"},
		Article{Title: "Hello 2", Description: "Description 2", Content: "Content 2"},
		Article{Title: "Hello 3", Description: "Description 3", Content: "Content 3"},
	}
	json.NewEncoder(w).Encode(articles)

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", getAllArticles)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
