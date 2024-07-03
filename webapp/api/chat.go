package api

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type ApiChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type ApiChatRequest struct {
	Message string `json:"message"`
}

const (
	chatUrl = "http://localhost:3001/api/chat"
)

func GetChatMessages(id string) ([]ApiChatMessage, error) {
	resp, err := http.Get(chatUrl + "?id=" + id)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var messages []ApiChatMessage
	err = json.NewDecoder(resp.Body).Decode(&messages)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func PostChatMessage(id string, message string) ([]ApiChatMessage, error) {
	request := ApiChatRequest{Message: message}
	jsonRequest, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(chatUrl+"?id="+id, "application/json", bytes.NewBuffer(jsonRequest))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var messages []ApiChatMessage
	err = json.NewDecoder(resp.Body).Decode(&messages)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
