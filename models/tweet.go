package models

import "gorm.io/gorm"

type Tweet struct {
	gorm.Model
	Message string `json:"message" gorm:"type:text;not null"`
	UserID  uint   `json:"user_id"`

	//foreign keys
	User *User `json:"-"`
}
