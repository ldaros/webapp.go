package controllers

import (
	"log-api/views"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	views.RenderIndex(w)
}
