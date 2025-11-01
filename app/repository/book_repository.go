package repository

import (
	"errors"
	"library-be/app/domain/dao"
	"library-be/app/domain/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (repo *BookRepository) FindAll() []dao.Book {
	var books []dao.Book

	repo.db.Find(&books)
	return books
}

func (repo *BookRepository) Find(id string) (*dao.Book, error) {
	uuid := uuid.MustParse(id)
	book := new(dao.Book)
	book.ID = &uuid

	if tx := repo.db.
		Preload("Author").
		Preload("Publisher").
		Find(book); tx.RowsAffected == 0 {
		return nil, errors.New("record not found")
	}

	return book, nil
}

func (repo *BookRepository) Save(data dto.BookDto, authors []dao.Author) (*dao.Book, error) {
	book := dao.Book{
		Title:       data.Title,
		Description: data.Description,
		ISBN:        data.ISBN,
		PublishedAt: data.PublishedAt,
		PublisherID: data.PublisherID,
		Author:      authors,
	}

	if data.ID == nil {
		if tx := repo.db.Create(&book); tx.Error != nil {
			return nil, tx.Error
		}
	} else {
		book.ID = data.ID

		// update book
		if tx := repo.db.Updates(&book); tx.Error != nil {
			return nil, tx.Error
		}

		// replace author
		if err := repo.db.Model(&book).Association("Author").Replace(authors); err != nil {
			return nil, err
		}
	}

	return repo.Find(book.ID.String())
}

func (repo *BookRepository) Delete(id string) error {
	book, err := repo.Find(id)
	if err != nil {
		return err
	}

	if tx := repo.db.Delete(book); tx.Error != nil {
		return tx.Error
	}

	return nil
}
