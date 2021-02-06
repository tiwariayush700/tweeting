package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Action struct {
	gorm.Model
	Message string         `json:"message" gorm:"type:text;not null"`
	Status  ActionStatus   `json:"status" binding:"required,oneof=pending approved rejected" gorm:"type:text;check:status = 'pending' or status = 'approved' or status = 'rejected';not null"`
	Body    datatypes.JSON `json:"token"`
}

type ActionStatus string
