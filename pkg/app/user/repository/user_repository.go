package repository

import (
	"errors"
	"gorm-in-go/common/error_handler"
	"gorm-in-go/pkg/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserById(id string) (*models.User, error)
	UpdateUser(id string, user *models.User) error
	DeleteUser(id string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetUserById(id string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, error_handler.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(id string, user *models.User) error {
	var existingUser models.User
	if err := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&existingUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return error_handler.ErrUserNotFound
		}
		return err
	}

	existingUser.Password = user.Password
	existingUser.Username = user.Username

	return r.db.Model(&existingUser).Updates(user).Error
}

func (r *userRepository) DeleteUser(id string) error {
	if err := r.db.Where("id = ? AND deleted_at IS NULL", id).Delete(&models.User{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return error_handler.ErrUserNotFound
		}
		return err
	}
	return nil
}
