package controllers

import (
	"log-api/db"
	"log-api/lib/logger"
	"log-api/views"
	"net/http"
)

func LogListHandler() http.HandlerFunc {
	logStore := db.NewLogStoreJson()

	return func(w http.ResponseWriter, r *http.Request) {
		logs, err := logStore.GetAll()
		if err != nil {
			logger.Error("Failed to get logs: ", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		views.RenderLogList(w, logs)
	}
}
