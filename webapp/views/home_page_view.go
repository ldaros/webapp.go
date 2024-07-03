package views

import (
	"html/template"
	"net/http"
)

type HomePageContent struct {
	Title           string
	MenuItems       []MenuItem
	BreadcrumbItems []BreadcrumbItem
}

func RenderHomePage(w http.ResponseWriter) {
	content := HomePageContent{
		Title:           "Home",
		MenuItems:       GetMenuItems("/"),
		BreadcrumbItems: GetBreadcrumbItems("Home", "/"),
	}

	tmpl := template.Must(template.ParseFiles(
		"templates/base.html",
		"templates/menu.html",
		"templates/breadcrumb.html",
		"templates/home.html"))
	tmpl.Execute(w, content)
}
