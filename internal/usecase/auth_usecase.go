package usecase

import (
	"be-sakoola/internal/config"
	"be-sakoola/internal/dto/request"
	"be-sakoola/models"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Register(input request.RegisterRequest) (*models.User, error)
	Login(req request.LoginRequest) (*models.User, error)
}

type authUsecase struct{}

func NewAuthUsecase() AuthUsecase {
	return &authUsecase{}
}

func (uc *authUsecase) Register(input request.RegisterRequest) (*models.User, error) {
	var existing models.User
	if err := config.DB.Where("email = ?", input.Email).First(&existing).Error; err == nil {
		return nil, errors.New("email already registered")
	}

	hPass, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	var userRole models.Role
	if err := config.DB.Where("LOWER(name) = ?", input.Role).First(&userRole).Error; err != nil {
		return nil, errors.New("role tidak terdaftar")
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hPass),
		Roles:    []models.Role{userRole},
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (uc *authUsecase) Login(input request.LoginRequest) (*models.User, error) {
	var user models.User
	if err := config.DB.Preload("Roles").Where("email = ?", input.Email).First(&user).Error; err != nil {
		return nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	return &user, nil
}
