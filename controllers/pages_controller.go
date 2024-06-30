package controllers

import (
	"log-api/views"
	"net/http"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	views.RenderHomePage(w)
}

func LogsPageHandler(w http.ResponseWriter, r *http.Request) {
	views.RenderLogsPage(w)
}
