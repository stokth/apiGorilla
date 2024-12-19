package userService

import (
	"apiGorilla/internal/taskService"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Email    string             `json:"email"`
	Password string             `json:"password"`
	Tasks    []taskService.Task `gorm:"foreignKey:UserID"`
}
