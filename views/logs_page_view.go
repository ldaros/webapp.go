package views

import (
	"html/template"
	"net/http"
)

type LogsPageContent struct {
	Title           string
	MenuItems       []MenuItem
	BreadcrumbItems []BreadcrumbItem
}

func RenderLogsPage(w http.ResponseWriter) {
	content := LogsPageContent{
		Title: "Home",
		MenuItems: []MenuItem{
			{URL: "/", Name: "Home", Active: false},
			{URL: "/logs", Name: "Logs", Active: true},
		},
		BreadcrumbItems: []BreadcrumbItem{
			{URL: "#", Name: "General", Active: false},
			{URL: "/logs", Name: "Logs", Active: true},
		},
	}

	tmpl := template.Must(template.ParseFiles(
		"templates/base.html",
		"templates/menu.html",
		"templates/breadcrumb.html",
		"templates/logs_page.html"))
	tmpl.Execute(w, content)
}
