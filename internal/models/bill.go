package models

import "gorm.io/gorm"

type Bill struct {
	gorm.Model
	Type  int    `json:"type" gorm:"not null"`
	Owner string `json:"owner" gorm:"index:idx_owner;not null"`
	Name  string `json:"name" gorm:"not null"`
	Cost  int    `json:"cost" gorm:"not null"`
	State int    `json:"state" gorm:"not null"`
}
