package views

import (
	"html/template"
	"log-api/models"
	"net/http"
)

func RenderLogList(w http.ResponseWriter, logs []models.Log) {
	data := struct {
		Logs []models.Log
	}{
		Logs: logs,
	}

	tmpl := template.Must(template.ParseFiles("templates/log_list.html"))
	tmpl.Execute(w, data)
	w.Header().Set("Content-Type", "text/html")
}
