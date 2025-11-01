package repository

import (
	"errors"
	"library-be/app/domain/dao"
	"library-be/app/domain/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{
		db: db,
	}
}

func (repo *AuthorRepository) FindAll() []dao.Author {
	var authors []dao.Author

	repo.db.Find(&authors)
	return authors
}

func (repo *AuthorRepository) Find(id string) (*dao.Author, error) {
	uuid := uuid.MustParse(id)
	author := new(dao.Author)
	author.ID = &uuid

	if tx := repo.db.
		Preload("Book").
		Find(author); tx.RowsAffected == 0 {
		return nil, errors.New("record not found")
	}

	return author, nil
}

func (repo *AuthorRepository) Save(data dto.AuthorDto) (*dao.Author, error) {
	author := dao.Author{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Biography: data.Biography,
	}

	if data.ID == nil {
		if tx := repo.db.Create(&author); tx.Error != nil {
			return nil, tx.Error
		}
	} else {
		author.ID = data.ID

		if tx := repo.db.Updates(&author); tx.Error != nil {
			return nil, tx.Error
		}
	}

	return repo.Find(author.ID.String())
}

func (repo *AuthorRepository) Delete(id string) error {
	author, err := repo.Find(id)
	if err != nil {
		return err
	}

	if tx := repo.db.Delete(author); tx.Error != nil {
		return tx.Error
	}

	return nil
}
