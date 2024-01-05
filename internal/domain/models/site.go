package models

import (
	"time"

	"gorm.io/gorm"
)

type Site struct {
	ID         int64          `json:"id"`
	Name       string         `json:"name"`
	AccessTime int64          `json:"access_time"`
	State      string         `json:"state"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-"`
}
