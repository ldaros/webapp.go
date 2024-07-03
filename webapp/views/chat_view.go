package views

import (
	"html/template"
	"net/http"
)

type ChatPageContent struct {
	Title           string
	MenuItems       []MenuItem
	BreadcrumbItems []BreadcrumbItem
}

func RenderChatPage(w http.ResponseWriter) {
	content := ChatPageContent{
		Title:           "Chat",
		MenuItems:       GetMenuItems("/chat"),
		BreadcrumbItems: GetBreadcrumbItems("Chat", "/chat"),
	}

	tmpl := template.Must(template.ParseFiles(
		"templates/base.html",
		"templates/menu.html",
		"templates/breadcrumb.html",
		"templates/chat_page.html"))
	tmpl.Execute(w, content)
}
