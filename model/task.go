package model

import "time"

type Task struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Description string `json:"description"`
	Status      string `json:"status"`
	createdAt   time.Time
	updatedAt   time.Time
}