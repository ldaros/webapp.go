package main

import (
	"log"
	"log-api/controllers"
	"log-api/db"
	"log-api/lib/logger"
	"net/http"
)

func main() {
	logStore := db.NewLogStoreJson()

	mux := http.NewServeMux()

	mux.HandleFunc("/api/log", controllers.LogAPIHandler(logStore))
	mux.HandleFunc("/views/log_list", controllers.LogListHandler())
	mux.HandleFunc("/", controllers.HomePageHandler)
	mux.HandleFunc("/logs", controllers.LogsPageHandler)
	mux.HandleFunc("/chat", controllers.ChatPageHandler)
	mux.HandleFunc("/views/chat_history", controllers.ChatHistoryHandler())
	mux.HandleFunc("/views/chat_send", controllers.ChatSendHandler())

	// Serve static files from the /assets directory
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	logger.Info("Starting server on :3000")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
