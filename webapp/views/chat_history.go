package views

import (
	"fmt"
	"html/template"
	"log-api/lib/logger"
	"log-api/models"
	"net/http"

	"github.com/russross/blackfriday/v2"
)

type ChatHistory struct {
	ChatMessages []models.ChatMessage
}

func RenderChatHistory(w http.ResponseWriter, chatMessages []models.ChatMessage) {
	content := ChatHistory{
		ChatMessages: chatMessages,
	}
	tmpl := template.Must(template.New("chat_history.html").Funcs(template.FuncMap{"markDown": markDowner}).ParseFiles("templates/chat_history.html"))
	err := tmpl.ExecuteTemplate(w, "chat_history.html", content)

	if err != nil {
		logger.Errorf("Error rendering chat history: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func markDowner(args ...interface{}) template.HTML {
	s := blackfriday.Run([]byte(fmt.Sprint(args...)))
	return template.HTML(s)
}
