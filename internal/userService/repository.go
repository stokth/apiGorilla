package userService

import "gorm.io/gorm"

type UsersRepository interface {
	// CreateTask - Передаем в функцию task типа Task их orm.go
	// возвращаем созданный Task и ошибку
	CreateUser(task Users) (Users, error)
	// GetAllTasks - Возвращаем массив из всех задач в БД и ошибку
	GetAllUsers() ([]Users, error)
	// UpdateTaskByID - Передаем id и Task, возвращаем обновленный Task
	// и ошибку
	UpdateUserByID(id uint, user Users) (Users, error)
	// DeleteTaskByID - Передаем id для удаления, возвращаем только ошибку
	DeleteUserByID(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user Users) (Users, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return Users{}, result.Error
	}
	return user, nil
}

func (r *userRepository) GetAllUsers() ([]Users, error) {
	var users []Users
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) UpdateUserByID(id uint, user Users) (Users, error) {
	result := r.db.Model(&Users{}).Where("id =?", id).Update("email", user.Email).Update("password", user.Password)
	if result.Error != nil {
		return Users{}, result.Error
	}
	return user, nil
}

func (r *userRepository) DeleteUserByID(id uint) error {
	return r.db.Delete(&Users{}, id).Error
}
