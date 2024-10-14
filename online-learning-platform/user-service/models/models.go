package models

import (
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `json:"-"`
}

// TableName overrides the table name used by GORM
func (UserModel) TableName() string {
	return "users"
}