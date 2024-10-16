package models

import "time"

type UserModel struct {
	ID        uint   `gorm:"primaryKey"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Username  string `gorm:"unique;not null"`
	DateOfBirth string
	CreatedAt time.Time
	UpdatedAt time.Time
}
// TableName overrides the table name used by GORM
func (UserModel) TableName() string {
	return "users"
}