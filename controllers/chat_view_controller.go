package controllers

import (
	"log-api/lib/logger"
	"log-api/models"
	"log-api/views"
	"net/http"
)

func ChatHistoryHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")

		logger.Infof("Chat history requested for chat ID: %s", id)

		messages := []models.ChatMessage{
			{Username: "Agent", IsAgent: true, Message: "Hello, how can I help you today?", Time: "11:09 PM"},
		}

		views.RenderChatHistory(w, messages)
	}
}

func ChatSendHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the form data
		r.ParseForm()

		// Get the chat ID
		id := r.Form.Get("id")

		// Get message from the form
		message := r.Form.Get("message")

		logger.Infof("Chat message received for %s: %s", id, message)

		messages := []models.ChatMessage{
			{Username: "Agent", IsAgent: true, Message: "Hello, how can I help you today?", Time: "11:09 PM"},
			{Username: "Customer", IsAgent: false, Message: message, Time: "11:10 PM"},
		}

		views.RenderChatHistory(w, messages)
	}
}
