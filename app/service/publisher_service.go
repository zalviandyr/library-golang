package service

import (
	"library-be/app/domain/dao"
	"library-be/app/domain/dto"
	"library-be/app/repository"
)

type PublisherService struct {
	repo *repository.PublisherRepository
}

func NewPublisherService(
	repo *repository.PublisherRepository,
) *PublisherService {
	return &PublisherService{
		repo: repo,
	}
}

func (s *PublisherService) FindAll() []dao.Publisher {
	return s.repo.FindAll()
}

func (s *PublisherService) Find(id string) (*dao.Publisher, error) {
	return s.repo.Find(id)
}

func (s *PublisherService) Save(data dto.PublisherDto) (*dao.Publisher, error) {
	return s.repo.Save(data)
}

func (s *PublisherService) Delete(id string) error {
	return s.repo.Delete(id)
}
