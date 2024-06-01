package utils

import(
	"net/http"
	"html/template"
)

func RenderTemplate(w http.ResponseWriter, asciiArt string) {
	tmpl := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>ASCII Art</title>
</head>
<body>
    <h1>Generated ASCII Art</h1>
    <pre>{{.}}</pre>
    <a href="/">Back</a>
</body>
</html>
`
	t, err := template.New("ascii").Parse(tmpl)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	t.Execute(w, asciiArt)
}