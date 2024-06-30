package views

import (
	"html/template"
	"log-api/models"
	"net/http"
)

type LogViewContent struct {
	LogGroups []LogGroup
}

type LogGroup struct {
	SequenceID string
	Logs       []LogRender
}

type LogRender struct {
	IsLeader bool
	Log      models.Log
}

func RenderLogList(w http.ResponseWriter, logs []models.Log) {
	content := LogViewContent{}
	var currentGroup *LogGroup

	for _, log := range logs {
		if currentGroup == nil || log.SequenceID != currentGroup.SequenceID {
			if currentGroup != nil {
				content.LogGroups = append(content.LogGroups, *currentGroup)
			}
			currentGroup = &LogGroup{SequenceID: log.SequenceID}
		}
		isFirst := currentGroup != nil && len(currentGroup.Logs) == 0
		currentGroup.Logs = append(currentGroup.Logs, LogRender{IsLeader: isFirst, Log: log})
	}
	if currentGroup != nil {
		content.LogGroups = append(content.LogGroups, *currentGroup)
	}

	tmpl := template.Must(template.ParseFiles("templates/log_list.html"))
	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, content)
}
