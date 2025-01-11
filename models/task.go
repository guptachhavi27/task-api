package models

import (
	"time"
)

type Task struct {
    ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    Title       string    `gorm:"not null" json:"title" binding:"required"`
    Description string    `gorm:"not null" json:"description" binding:"required"`
    Status      string    `gorm:"type:task_status;default:'pending'" json:"status"`  
    CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
