package repository

import (
	"errors"
	"library-be/app/domain/dao"
	"library-be/app/domain/dto"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (repo *AuthRepository) FindByUsername(username string) (*dao.User, error) {
	user := dao.User{
		Username: username,
	}

	if tx := repo.db.
		Find(&user); tx.RowsAffected == 0 {
		return nil, errors.New("username not found")
	}

	return &user, nil
}

func (repo *AuthRepository) Create(data dto.RegisterDto) error {
	user := dao.User{
		Username: data.Username,
		Password: data.Password,
	}

	if tx := repo.db.Create(&user); tx.Error != nil {
		return tx.Error
	}

	return nil
}
