package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	cmd "github.com/ravisantoshgudimetla/aggregation-parser/cmd"
)

func handler(w http.ResponseWriter, r *http.Request) {
	url := os.Args[1]
	jobs := cmd.Parse(url)
	parsedTemplate, err := template.ParseFiles("templates/job_status.html")
	if err != nil {
		log.Fatalf("error parsing template %v", err)
	}
	if err := parsedTemplate.Execute(w, jobs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
