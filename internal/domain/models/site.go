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

type GetSiteInput struct {
	Name                string `json:"name"`
	IsMaximumAccessTime bool   `json:"is_maximum_access_time"`
	IsMinimumAccessTime bool   `json:"is_minimum_access_time"`
}

type GetSiteOutput struct {
	Name       string `json:"name"`
	State      string `json:"state"`
	AccessTime int64  `json:"access_time"`
}
