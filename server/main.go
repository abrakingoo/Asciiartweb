package main

import (
	"html/template"
	"net/http"
	"fmt"
	"asciiartweb/server/api"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			tmpl := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Ascii Art Web</title>
</head>
<body>
    <h1>Ascii Art Web</h1>
    <form action="/ascii-art" method="post">
        <label for="user_input">Text Input</label>
        <input type="text" name="user_input" id="user_input"><br>
        <h4>Change Banner File</h4>
        <fieldset>
            <legend>Banner Options</legend>
            <input type="radio" name="banner" id="standard" value="standard" checked>
            <label for="standard1">Standard</label><br>
            <input type="radio" name="banner" id="shadow" value="shadow">
            <label for="shadow">Shadow</label><br>
            <input type="radio" name="banner" id="thinkertoy" value="thinkertoy">
            <label for="thinkertoy">Thinkertoy</label><br>
        </fieldset>
        <input type="submit" value="Create Art">
    </form>
    <div id="response"></div>
</body>
</html>
`
			t, _ := template.New("index").Parse(tmpl)
			t.Execute(w, nil)
		} else {
			api.PostHandler(w, r)
		}
	})

	fmt.Println("Server running on http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
