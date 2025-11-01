package book_test

import (
	"testing"
	"time"

	"github.com/google/uuid"

	"library-be/app/domain/dao"
	"library-be/app/domain/dto"
	"library-be/app/repository"
	"library-be/config"
	"library-be/test"
)

func TestBookRepositoryFindAll(t *testing.T) {
	env := config.InitEnvironmentWithFiles("../../.env.testing")
	db := test.InitTestingDatabase(env)

	authorRepo := repository.NewAuthorRepository(db)
	publisherRepo := repository.NewPublisherRepository(db)
	bookRepo := repository.NewBookRepository(db)

	author, err := authorRepo.Save(dto.AuthorDto{
		FirstName: "John",
		LastName:  "Doe",
		Biography: "Bio 1",
	})
	if err != nil {
		t.Fatalf("failed to create author: %v", err)
	}

	publisher, err := publisherRepo.Save(dto.PublisherDto{
		Name:    "Publisher One",
		Address: "Address One",
		Phone:   "08111111111",
		Website: "https://publisher.one",
	})
	if err != nil {
		t.Fatalf("failed to create publisher: %v", err)
	}

	bookData := dto.BookDto{
		Title:       "Book One",
		Description: "Description One",
		ISBN:        "ISBN-001",
		PublishedAt: time.Now(),
		PublisherID: publisher.ID,
	}

	if _, err := bookRepo.Save(bookData, []dao.Author{*author}); err != nil {
		t.Fatalf("failed to create book: %v", err)
	}

	bookData.Title = "Book Two"
	bookData.Description = "Description Two"
	bookData.ISBN = "ISBN-002"
	bookData.PublishedAt = time.Now().Add(time.Hour)

	if _, err := bookRepo.Save(bookData, []dao.Author{*author}); err != nil {
		t.Fatalf("failed to create second book: %v", err)
	}

	books := bookRepo.FindAll()
	if len(books) != 2 {
		t.Fatalf("expected 2 books, got %d", len(books))
	}
}

func TestBookRepositoryFind(t *testing.T) {
	env := config.InitEnvironmentWithFiles("../../.env.testing")
	db := test.InitTestingDatabase(env)

	authorRepo := repository.NewAuthorRepository(db)
	publisherRepo := repository.NewPublisherRepository(db)
	bookRepo := repository.NewBookRepository(db)

	author, err := authorRepo.Save(dto.AuthorDto{
		FirstName: "Alice",
		LastName:  "Walker",
		Biography: "Author bio",
	})
	if err != nil {
		t.Fatalf("failed to create author: %v", err)
	}

	publisher, err := publisherRepo.Save(dto.PublisherDto{
		Name:    "Publisher",
		Address: "Street 1",
		Phone:   "08123456789",
		Website: "https://publisher.test",
	})
	if err != nil {
		t.Fatalf("failed to create publisher: %v", err)
	}

	bookData := dto.BookDto{
		Title:       "Sample Book",
		Description: "A nice book",
		ISBN:        "ISBN-123",
		PublishedAt: time.Now(),
		PublisherID: publisher.ID,
	}
	saved, err := bookRepo.Save(bookData, []dao.Author{*author})
	if err != nil {
		t.Fatalf("failed to create book: %v", err)
	}

	result, err := bookRepo.Find(saved.ID.String())
	if err != nil {
		t.Fatalf("expected to find book, got error: %v", err)
	}
	if result.ID == nil || result.ID.String() != saved.ID.String() {
		t.Fatalf("expected book ID %s, got %v", saved.ID.String(), result.ID)
	}
	if result.Publisher == nil || result.Publisher.ID == nil || result.Publisher.ID.String() != publisher.ID.String() {
		t.Fatalf("expected publisher to be preloaded")
	}
	if len(result.Author) != 1 || result.Author[0].ID == nil || result.Author[0].ID.String() != author.ID.String() {
		t.Fatalf("expected author to be preloaded")
	}
}

func TestBookRepositoryFindNotFound(t *testing.T) {
	env := config.InitEnvironmentWithFiles("../../.env.testing")
	db := test.InitTestingDatabase(env)
	repo := repository.NewBookRepository(db)

	_, err := repo.Find(uuid.NewString())
	if err == nil {
		t.Fatal("expected error when book does not exist")
	}

	if err.Error() != "record not found" {
		t.Fatalf("expected 'record not found' error, got %v", err)
	}
}

func TestBookRepositorySaveUpdate(t *testing.T) {
	env := config.InitEnvironmentWithFiles("../../.env.testing")
	db := test.InitTestingDatabase(env)

	authorRepo := repository.NewAuthorRepository(db)
	publisherRepo := repository.NewPublisherRepository(db)
	bookRepo := repository.NewBookRepository(db)

	authorOne, err := authorRepo.Save(dto.AuthorDto{
		FirstName: "Mary",
		LastName:  "Shelley",
		Biography: "Author one",
	})
	if err != nil {
		t.Fatalf("failed to create author: %v", err)
	}

	authorTwo, err := authorRepo.Save(dto.AuthorDto{
		FirstName: "Jane",
		LastName:  "Austen",
		Biography: "Author two",
	})
	if err != nil {
		t.Fatalf("failed to create second author: %v", err)
	}

	publisher, err := publisherRepo.Save(dto.PublisherDto{
		Name:    "Publisher",
		Address: "Street 1",
		Phone:   "08123456789",
		Website: "https://publisher.test",
	})
	if err != nil {
		t.Fatalf("failed to create publisher: %v", err)
	}

	bookData := dto.BookDto{
		Title:       "Original Book",
		Description: "Original description",
		ISBN:        "ISBN-123",
		PublishedAt: time.Now(),
		PublisherID: publisher.ID,
	}
	saved, err := bookRepo.Save(bookData, []dao.Author{*authorOne})
	if err != nil {
		t.Fatalf("failed to create book: %v", err)
	}

	bookData.ID = saved.ID
	bookData.Title = "Updated Book"
	bookData.Description = "Updated description"
	bookData.ISBN = "ISBN-456"
	bookData.PublishedAt = time.Now().Add(24 * time.Hour)

	updated, err := bookRepo.Save(bookData, []dao.Author{*authorTwo})
	if err != nil {
		t.Fatalf("failed to update book: %v", err)
	}

	if updated.Title != "Updated Book" {
		t.Fatalf("expected title to be updated, got %s", updated.Title)
	}
	if len(updated.Author) != 1 || updated.Author[0].ID == nil || updated.Author[0].ID.String() != authorTwo.ID.String() {
		t.Fatalf("expected authors to be replaced")
	}
}

func TestBookRepositoryDelete(t *testing.T) {
	env := config.InitEnvironmentWithFiles("../../.env.testing")
	db := test.InitTestingDatabase(env)

	authorRepo := repository.NewAuthorRepository(db)
	publisherRepo := repository.NewPublisherRepository(db)
	bookRepo := repository.NewBookRepository(db)

	author, err := authorRepo.Save(dto.AuthorDto{
		FirstName: "Leo",
		LastName:  "Tolstoy",
		Biography: "Author bio",
	})
	if err != nil {
		t.Fatalf("failed to create author: %v", err)
	}

	publisher, err := publisherRepo.Save(dto.PublisherDto{
		Name:    "Publisher",
		Address: "Street 1",
		Phone:   "08123456789",
		Website: "https://publisher.test",
	})
	if err != nil {
		t.Fatalf("failed to create publisher: %v", err)
	}

	bookData := dto.BookDto{
		Title:       "Book To Delete",
		Description: "Delete description",
		ISBN:        "ISBN-DELETE",
		PublishedAt: time.Now(),
		PublisherID: publisher.ID,
	}
	saved, err := bookRepo.Save(bookData, []dao.Author{*author})
	if err != nil {
		t.Fatalf("failed to create book: %v", err)
	}

	if err := bookRepo.Delete(saved.ID.String()); err != nil {
		t.Fatalf("failed to delete book: %v", err)
	}

	if _, err := bookRepo.Find(saved.ID.String()); err == nil {
		t.Fatal("expected error when fetching deleted book")
	}
}
