package service

import (
	"errors"
	"library-be/app/domain/dao"
	"library-be/app/domain/dto"
	"library-be/app/repository"

	"github.com/google/uuid"
)

type BookService struct {
	authorRepo *repository.AuthorRepository
	repo       *repository.BookRepository
}

func NewBookService(
	authorRepo *repository.AuthorRepository,
	repo *repository.BookRepository,
) *BookService {
	return &BookService{
		authorRepo: authorRepo,
		repo:       repo,
	}
}

func (s *BookService) FindAll() []dao.Book {
	return s.repo.FindAll()
}

func (s *BookService) Find(id string) (*dao.Book, error) {
	return s.repo.Find(id)
}

func (s *BookService) Save(data dto.BookDto) (*dao.Book, error) {
	seen := make(map[uuid.UUID]struct{})
	uniqueIDs := make([]uuid.UUID, 0, len(data.AuthorIDs))
	for _, id := range data.AuthorIDs {
		if id == uuid.Nil {
			continue
		}
		if _, ok := seen[id]; ok {
			continue
		}
		seen[id] = struct{}{}
		uniqueIDs = append(uniqueIDs, id)
	}

	var authors []dao.Author
	for _, authorId := range uniqueIDs {
		if author, _ := s.authorRepo.Find(authorId.String()); author != nil {
			authors = append(authors, *author)
		}
	}

	if len(authors) == 0 {
		return nil, errors.New("author cannot empty or invalid author id")
	}

	return s.repo.Save(data, authors)
}

func (s *BookService) Delete(id string) error {
	return s.repo.Delete(id)
}
