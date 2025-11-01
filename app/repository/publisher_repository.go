package repository

import (
	"errors"
	"library-be/app/domain/dao"
	"library-be/app/domain/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PublisherRepository struct {
	db *gorm.DB
}

func NewPublisherRepository(db *gorm.DB) *PublisherRepository {
	return &PublisherRepository{
		db: db,
	}
}

func (repo *PublisherRepository) FindAll() []dao.Publisher {
	var publishers []dao.Publisher

	repo.db.Find(&publishers)
	return publishers
}

func (repo *PublisherRepository) Find(id string) (*dao.Publisher, error) {
	uuid := uuid.MustParse(id)
	publisher := new(dao.Publisher)
	publisher.ID = &uuid

	if tx := repo.db.
		Preload("Book").
		Find(publisher); tx.RowsAffected == 0 {
		return nil, errors.New("record not found")
	}

	return publisher, nil
}

func (repo *PublisherRepository) Save(data dto.PublisherDto) (*dao.Publisher, error) {
	publisher := dao.Publisher{
		Name:    data.Name,
		Address: data.Address,
		Phone:   data.Phone,
		Website: data.Website,
	}

	if data.ID == nil {
		if tx := repo.db.Create(&publisher); tx.Error != nil {
			return nil, tx.Error
		}
	} else {
		publisher.ID = data.ID

		if tx := repo.db.Updates(&publisher); tx.Error != nil {
			return nil, tx.Error
		}
	}

	return repo.Find(publisher.ID.String())
}

func (repo *PublisherRepository) Delete(id string) error {
	publisher, err := repo.Find(id)
	if err != nil {
		return err
	}

	if tx := repo.db.Delete(publisher); tx.Error != nil {
		return tx.Error
	}

	return nil
}
