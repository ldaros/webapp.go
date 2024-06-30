package models

import "time"

type Log struct {
	ID         int       `json:"id"`
	SequenceID string    `json:"sequence_id"`
	Name       string    `json:"name"`
	Input      string    `json:"input"`
	Output     string    `json:"output"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	Tokens     int       `json:"tokens"`
}
