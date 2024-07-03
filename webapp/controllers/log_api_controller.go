package controllers

import (
	"encoding/json"
	"log-api/db"
	"log-api/lib/logger"
	"log-api/models"
	"net/http"
	"time"
)

type CreateLogRequest struct {
	SequenceID string `json:"sequence_id"`
	Name       string `json:"name"`
	Input      string `json:"input"`
	Output     string `json:"output"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	Tokens     int    `json:"tokens"`
}

type CreateLogResponse struct {
	ID int `json:"id"`
}

func LogAPIHandler(logStore db.LogStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			CreateLog(w, r)
		case "GET":
			ListLogs(w)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}
}

func CreateLog(w http.ResponseWriter, r *http.Request) {
	logStore := db.NewLogStoreJson()

	var req CreateLogRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	logger.Debug("CreateLogRequest: ", req)

	// parse time
	startTime, err := time.Parse(time.RFC3339, req.StartTime)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	endTime, err := time.Parse(time.RFC3339, req.EndTime)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	log := models.Log{
		SequenceID: req.SequenceID,
		Name:       req.Name,
		Input:      req.Input,
		Output:     req.Output,
		StartTime:  startTime,
		EndTime:    endTime,
		Tokens:     req.Tokens,
	}

	log, err = logStore.Insert(log)
	if err != nil {
		logger.Error("Failed to insert log: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	logger.Infof("Log created with ID: %d", log.ID)

	// Return the log as JSON
	resp := CreateLogResponse{ID: log.ID}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		logger.Error("Failed to marshal response: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResp)
}

func ListLogs(w http.ResponseWriter) {
	logStore := db.NewLogStoreJson()

	logs, err := logStore.GetAll()
	if err != nil {
		logger.Error("Failed to get logs: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	jsonLogs, err := json.Marshal(logs)
	if err != nil {
		logger.Error("Failed to marshal logs: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonLogs)
}
