package model

import "github.com/jinzhu/gorm"

// Info struct
type Infos struct {
	gorm.Model
	Company string `json:"company"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Msg     string `json:"msg"`
}