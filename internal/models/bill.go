package models

import "gorm.io/gorm"

type Bill struct {
	gorm.Model
	Type  int
	Owner string
	name  string
	cost  int
	state int
}
