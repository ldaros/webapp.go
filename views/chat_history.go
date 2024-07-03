package views

import (
	"html/template"
	"log-api/models"
	"net/http"
)

type ChatHistory struct {
	ChatMessages []models.ChatMessage
}

func RenderChatHistory(w http.ResponseWriter, chatMessages []models.ChatMessage) {
	content := ChatHistory{
		ChatMessages: chatMessages,
	}

	tmpl := template.Must(template.ParseFiles(
		"templates/chat_history.html"))
	tmpl.Execute(w, content)
}
