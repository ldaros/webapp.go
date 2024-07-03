package controllers

import (
	"log-api/views"
	"net/http"

	"github.com/google/uuid"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	views.RenderHomePage(w)
}

func LogsPageHandler(w http.ResponseWriter, r *http.Request) {
	views.RenderLogsPage(w)
}

func ChatPageHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		id = uuid.New().String()
		http.Redirect(w, r, "/chat?id="+id, http.StatusSeeOther)
		return
	}

	views.RenderChatPage(w)
}
