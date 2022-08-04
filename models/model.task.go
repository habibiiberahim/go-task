package models

import "time"

type ModelTask struct {
	ID        uint `json:"id"`
	Description     string `json:"description"`
	IsDone     bool `json:"is_done"`
	Assignn    string `json:"assignn"`
	Deadline time.Time `json:"deadline"`
}