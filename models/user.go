package models

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"firstName" validate:"required" binding:"required" gorm:"type:text;not null"`
	LastName  string `json:"lastName" validate:"required" binding:"required" gorm:"type:text;not null"`
	Email     string `json:"email" binding:"required" gorm:"uniqueIndex"`
	Role      string `json:"role" binding:"required,oneof=ADMIN USER SUPER_ADMIN" validate:"required,oneof=ADMIN USER" gorm:"type:text;check:role = 'ADMIN' or role = 'USER' or role = 'SUPER_ADMIN';not null"`
	Phone     string `json:"phone" binding:"required" gorm:"type:text"`
	Password  string `json:"password" binding:"required" validate:"required" gorm:"type:text"`
}

type UserRole string

const (
	UserRoleAdmin      = UserRole("ADMIN")
	UserRoleEmployee   = UserRole("USER")
	UserRoleSuperAdmin = UserRole("SUPER_ADMIN")
)

type UserLoginJWTClaims struct {
	Authorized bool   `json:"authorized"`
	Id         uint   `json:"id"`
	Role       string `json:"role"`
	jwt.StandardClaims
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required" validate:"required"`
}

type UserResponse struct {
	gorm.Model
	FirstName string `json:"firstName" validate:"required" binding:"required" gorm:"type:text;not null"`
	LastName  string `json:"lastName" validate:"required" binding:"required" gorm:"type:text;not null"`
	Email     string `json:"email" binding:"required" gorm:"uniqueIndex"`
	Role      string `json:"role" binding:"required,oneof=ADMIN USER" validate:"required,oneof=ADMIN USER" gorm:"type:text;check:role = 'ADMIN' or role = 'USER';not null"`
	Phone     string `json:"phone" binding:"required" gorm:"type:text"`
}
