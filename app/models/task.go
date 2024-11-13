package models

import "gorm.io/gorm"

type Task struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      uint   `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	gorm.Model
}

func (Task) TableName() string {
	return "tasks"
}
