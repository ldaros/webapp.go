package views

type BreadcrumbItem struct {
	URL    string
	Name   string
	Active bool
}

func GetBreadcrumbItems(name string, url string) []BreadcrumbItem {
	return []BreadcrumbItem{
		{URL: "#", Name: "General", Active: false},
		{URL: url, Name: name, Active: true},
	}
}
