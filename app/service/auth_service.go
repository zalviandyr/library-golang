package service

import (
	"errors"
	"library-be/app/domain/dto"
	"library-be/app/pkg"
	"library-be/app/repository"
	"library-be/config"
	"time"
)

type AuthService struct {
	repo *repository.AuthRepository
	env  config.Environment
}

func NewAuthService(
	repo *repository.AuthRepository,
	env config.Environment,
) *AuthService {
	return &AuthService{
		repo: repo,
		env:  env,
	}
}

func (s *AuthService) Login(data dto.LoginDto) (string, error) {
	user, err := s.repo.FindByUsername(data.Username)
	if err != nil {
		return "", err
	}

	if isValid, _ := pkg.CheckPasswordHash(data.Password, user.Password); !isValid {
		return "", errors.New("invalid password")
	}

	// generate token
	token, err := pkg.GenerateJWT(user.ID.String(), s.env.JWT_SECRET, time.Hour*24)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) Register(data dto.RegisterDto) error {
	if user, _ := s.repo.FindByUsername(data.Username); user != nil {
		return errors.New("username already taken")
	}

	if data.Password != data.ConfirmPassword {
		return errors.New("password and confirm_password must same")
	}

	password, err := pkg.HashPassword(data.Password)
	if err != nil {
		return err
	}
	data.Password = password

	return s.repo.Create(data)
}
