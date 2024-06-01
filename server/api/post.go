package api

import (
	"asciiartweb/server/art"
	"asciiartweb/server/utils"
	"net/http"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)
		return
	}

	userInput := r.FormValue("user_input")
	banner := r.FormValue("banner")

	asciiArt := art.Art(userInput, banner)

	utils.RenderTemplate(w, asciiArt)
}
