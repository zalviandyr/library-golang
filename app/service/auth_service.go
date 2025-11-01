package service

import (
	"errors"
	"library-be/app/domain/dto"
	"library-be/app/pkg"
	"library-be/app/repository"
)

type AuthService struct {
	repo *repository.AuthRepository
}

func NewAuthService(
	repo *repository.AuthRepository,
) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) Login(data dto.LoginDto) error {
	user, err := s.repo.FindByUsername(data.Username)
	if err != nil {
		return err
	}

	if isValid, _ := pkg.CheckPasswordHash(data.Password, user.Password); !isValid {
		return errors.New("invalid password")
	}

	return nil
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
