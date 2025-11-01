package book_test

import (
	"testing"
	"time"

	"github.com/google/uuid"

	"library-be/app/domain/dto"
	"library-be/app/repository"
	"library-be/app/service"
	"library-be/config"
	"library-be/test"
)

func TestBookServiceFindAll(t *testing.T) {
	env := config.InitEnvironmentWithFiles("../../.env.testing")
	db := test.InitTestingDatabase(env)

	authorRepo := repository.NewAuthorRepository(db)
	publisherRepo := repository.NewPublisherRepository(db)
	bookRepo := repository.NewBookRepository(db)
	bookService := service.NewBookService(authorRepo, bookRepo)

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
		AuthorIDs:   []uuid.UUID{*author.ID},
	}

	if _, err := bookService.Save(bookData); err != nil {
		t.Fatalf("failed to create book: %v", err)
	}

	bookData.Title = "Book Two"
	bookData.Description = "Description Two"
	bookData.ISBN = "ISBN-002"
	bookData.PublishedAt = time.Now().Add(time.Hour)

	if _, err := bookService.Save(bookData); err != nil {
		t.Fatalf("failed to create second book: %v", err)
	}

	books := bookService.FindAll()
	if len(books) != 2 {
		t.Fatalf("expected 2 books, got %d", len(books))
	}
}

func TestBookServiceFind(t *testing.T) {
	env := config.InitEnvironmentWithFiles("../../.env.testing")
	db := test.InitTestingDatabase(env)

	authorRepo := repository.NewAuthorRepository(db)
	publisherRepo := repository.NewPublisherRepository(db)
	bookRepo := repository.NewBookRepository(db)
	bookService := service.NewBookService(authorRepo, bookRepo)

	author, err := authorRepo.Save(dto.AuthorDto{
		FirstName: "Alice",
		LastName:  "Walker",
		Biography: "Bio",
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
		AuthorIDs:   []uuid.UUID{*author.ID},
	}
	created, err := bookService.Save(bookData)
	if err != nil {
		t.Fatalf("failed to create book: %v", err)
	}

	result, err := bookService.Find(created.ID.String())
	if err != nil {
		t.Fatalf("expected find to succeed: %v", err)
	}
	if result.ID == nil || result.ID.String() != created.ID.String() {
		t.Fatalf("expected book ID %s, got %v", created.ID.String(), result.ID)
	}
}

func TestBookServiceSaveUpdate(t *testing.T) {
	env := config.InitEnvironmentWithFiles("../../.env.testing")
	db := test.InitTestingDatabase(env)

	authorRepo := repository.NewAuthorRepository(db)
	publisherRepo := repository.NewPublisherRepository(db)
	bookRepo := repository.NewBookRepository(db)
	bookService := service.NewBookService(authorRepo, bookRepo)

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
		AuthorIDs:   []uuid.UUID{*authorOne.ID},
	}
	saved, err := bookService.Save(bookData)
	if err != nil {
		t.Fatalf("failed to create book: %v", err)
	}

	bookData.ID = saved.ID
	bookData.Title = "Updated Book"
	bookData.Description = "Updated description"
	bookData.ISBN = "ISBN-456"
	bookData.PublishedAt = time.Now().Add(24 * time.Hour)
	bookData.AuthorIDs = []uuid.UUID{*authorTwo.ID}

	updated, err := bookService.Save(bookData)
	if err != nil {
		t.Fatalf("failed to update book: %v", err)
	}
	if updated.Title != "Updated Book" {
		t.Fatalf("expected updated title, got %s", updated.Title)
	}
	if len(updated.Author) != 1 || updated.Author[0].ID == nil || updated.Author[0].ID.String() != authorTwo.ID.String() {
		t.Fatalf("expected authors to be replaced")
	}
}

func TestBookServiceSaveInvalidAuthors(t *testing.T) {
	env := config.InitEnvironmentWithFiles("../../.env.testing")
	db := test.InitTestingDatabase(env)

	authorRepo := repository.NewAuthorRepository(db)
	publisherRepo := repository.NewPublisherRepository(db)
	bookRepo := repository.NewBookRepository(db)
	bookService := service.NewBookService(authorRepo, bookRepo)

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
		Title:       "Invalid Book",
		Description: "Should fail",
		ISBN:        "ISBN-FAIL",
		PublishedAt: time.Now(),
		PublisherID: publisher.ID,
		AuthorIDs:   []uuid.UUID{},
	}
	if _, err := bookService.Save(bookData); err == nil {
		t.Fatal("expected error when providing no valid authors")
	}
}

func TestBookServiceDelete(t *testing.T) {
	env := config.InitEnvironmentWithFiles("../../.env.testing")
	db := test.InitTestingDatabase(env)

	authorRepo := repository.NewAuthorRepository(db)
	publisherRepo := repository.NewPublisherRepository(db)
	bookRepo := repository.NewBookRepository(db)
	bookService := service.NewBookService(authorRepo, bookRepo)

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
		AuthorIDs:   []uuid.UUID{*author.ID},
	}
	saved, err := bookService.Save(bookData)
	if err != nil {
		t.Fatalf("failed to create book: %v", err)
	}

	if err := bookService.Delete(saved.ID.String()); err != nil {
		t.Fatalf("failed to delete book: %v", err)
	}

	if _, err := bookService.Find(saved.ID.String()); err == nil {
		t.Fatal("expected error when fetching deleted book")
	}
}
