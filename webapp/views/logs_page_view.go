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
		Title:           "Logs",
		MenuItems:       GetMenuItems("/logs"),
		BreadcrumbItems: GetBreadcrumbItems("Logs", "/logs"),
	}

	tmpl := template.Must(template.ParseFiles(
		"templates/base.html",
		"templates/menu.html",
		"templates/breadcrumb.html",
		"templates/logs_page.html"))
	tmpl.Execute(w, content)
}
