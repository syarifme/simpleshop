package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `gorm:"unique" json:"username"`
	Email    string `gorm:"unique_index" json:"email"`
	Password string `json:"password"`
	Picture  string `json:"picture"`
	Status   string `json:"status"`
}
