package model

import "github.com/jinzhu/gorm"

// User struct
type Users struct {
	gorm.Model
	Username string `gorm:"UNIQUE;NOT NULL" json:"username"`
	Email    string `gorm:"UNIQUE" json:"email"`
	Password string `gorm:"NOT NULL" json:"password"`
}
