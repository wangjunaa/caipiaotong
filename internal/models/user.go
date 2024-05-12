package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username string    `json:"username" gorm:"not null"`
	Phone    string    `json:"phone" gorm:"index:idx_phone;unique;not null"`
	Password string    `json:"password" gorm:"not null"`
	LoginAt  time.Time `json:"loginAt"`
	IsAdmin  bool      `json:"isAdmin"`
}

func (user *User) TableName() string {
	return "users"
}

//type SafeUserData struct {
//	Username string
//	Phone    string
//	LoginAt  time.Time
//	CreateAt time.Time
//	UpdateAt time.Time
//}

//func (user *User) ToSafeUserData() SafeUserData {
//	safeData := SafeUserData{
//		Username: user.Username,
//		Phone:    user.Phone,
//		LoginAt:  user.LoginAt,
//		CreateAt: user.CreatedAt,
//		UpdateAt: user.UpdatedAt,
//	}
//	return safeData
//}
