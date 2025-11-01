package service

import (
	"library-be/app/domain/dao"
	"library-be/app/domain/dto"
	"library-be/app/repository"
)

type AuthorService struct {
	repo *repository.AuthorRepository
}

func NewAuthorService(
	repo *repository.AuthorRepository,
) *AuthorService {
	return &AuthorService{
		repo: repo,
	}
}

func (s *AuthorService) FindAll() []dao.Author {
	return s.repo.FindAll()
}

func (s *AuthorService) Find(id string) (*dao.Author, error) {
	return s.repo.Find(id)
}

func (s *AuthorService) Save(data dto.AuthorDto) (*dao.Author, error) {
	return s.repo.Save(data)
}

func (s *AuthorService) Delete(id string) error {
	return s.repo.Delete(id)
}
