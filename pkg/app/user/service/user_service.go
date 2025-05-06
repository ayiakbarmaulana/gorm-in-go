package service

import (
	"gorm-in-go/pkg/app/user/repository"
	"gorm-in-go/pkg/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(createUserRequest *CreateUserRequest) error
	GetUserById(id string) (*models.User, error)
	UpdateUser(id string, updateUserRequest *UpdateUserRequest) error
	DeleteUser(id string) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(createUserRequest *CreateUserRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createUserRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		ID:       uuid.New().String(),
		Email:    createUserRequest.Email,
		Password: string(hashedPassword),
		Username: createUserRequest.Username,
	}

	return s.userRepo.CreateUser(&user)
}

func (s *userService) GetUserById(id string) (*models.User, error) {
	return s.userRepo.GetUserById(id)
}

func (s *userService) UpdateUser(id string, updateUserRequest *UpdateUserRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updateUserRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := models.User{
		Password: string(hashedPassword),
		Username: updateUserRequest.Username,
	}

	return s.userRepo.UpdateUser(id, &user)
}

func (s *userService) DeleteUser(id string) error {
	return s.userRepo.DeleteUser(id)
}

type CreateUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Username string `json:"username" validate:"required,min=3"`
}

type UpdateUserRequest struct {
	Password string `json:"password" validate:"required,min=8"`
	Username string `json:"username" validate:"required,min=3"`
}
