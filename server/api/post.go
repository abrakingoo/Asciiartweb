package api

import (
	"asciiartweb/server/art"
	"html/template"
	"net/http"
	"os"
)

type artResult struct {
	Result string
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)
		os.Exit(0)
	}

	userInput := r.FormValue("user_input")
	banner := r.FormValue("banner")

	asciiArt := art.Art(userInput, banner)

	data := artResult {
		Result : asciiArt,
	}

	tmpl := template.Must(template.ParseFiles("templates/ascii-art.html"))

	tmpl.Execute(w, data)
}
