package views

type MenuItem struct {
	URL    string
	Name   string
	Active bool
}

func GetMenuItems(current string) []MenuItem {
	menuItems := []MenuItem{
		{URL: "/", Name: "Home", Active: false},
		{URL: "/logs", Name: "Logs", Active: false},
		{URL: "/chat", Name: "Chat", Active: false},
	}

	for i, item := range menuItems {
		if item.URL == current {
			menuItems[i].Active = true
		}
	}

	return menuItems
}
