package models

import (
	"time"
)

type Bill struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt" gorm:"index"`
	Type      int       `json:"type" gorm:"not null"`
	Owner     string    `json:"owner" gorm:"index:idx_owner;not null"`
	Name      string    `json:"name" gorm:"not null"`
	Cost      int       `json:"cost" gorm:"index:idx_cost;not null"`
	State     int       `json:"state" gorm:"not null"`
}
