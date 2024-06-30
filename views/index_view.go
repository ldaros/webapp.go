package views

import (
	"html/template"
	"net/http"
)

func RenderIndex(w http.ResponseWriter) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}
