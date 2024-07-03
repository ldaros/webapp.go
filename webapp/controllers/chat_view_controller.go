package controllers

import (
	"log-api/api"
	"log-api/lib/logger"
	"log-api/models"
	"log-api/views"
	"net/http"
)

func ChatHistoryHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")

		logger.Infof("Chat history requested for chat ID: %s", id)

		messages, err := api.GetChatMessages(id)
		if err != nil {
			logger.Errorf("Error getting chat messages: %s", err)
			http.Error(w, "Error getting chat messages", http.StatusInternalServerError)
			return
		}

		views.RenderChatHistory(w, ConvertToViewMessages(messages))
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

		messages, err := api.PostChatMessage(id, message)
		if err != nil {
			logger.Errorf("Error posting chat message: %s", err)
			http.Error(w, "Error posting chat message", http.StatusInternalServerError)
			return
		}

		views.RenderChatHistory(w, ConvertToViewMessages(messages))
	}
}

func ConvertToViewMessages(messages []api.ApiChatMessage) []models.ChatMessage {
	var viewMessages []models.ChatMessage

	for _, message := range messages {
		role := "Agent"
		if message.Role == "user" {
			role = "User"
		}

		viewMessages = append(viewMessages, models.ChatMessage{
			Username: role,
			IsAgent:  message.Role == "assistant",
			Message:  message.Content,
			Time:     "11:09 PM",
		})
	}

	return viewMessages
}
