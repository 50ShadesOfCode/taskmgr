package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Task :
type Task struct {
	gorm.Model
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"desc"`
	StartTime   time.Time `json:"startTime"`
	FinishTime  time.Time `json:"finishTime"`
	IsCompleted bool      `json:"isCompleted"`
}

//TableName :
func (Task) TableName() string {
	return "tasks"
}
