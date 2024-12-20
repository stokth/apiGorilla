package userService

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
	UserID uint   `json:"user_id"`
	User   Users  `gorm:"foreignKey:UserID"`
}

type Users struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Tasks    []Task `gorm:"foreignKey:UserID"`
}
