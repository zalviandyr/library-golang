package author_test

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

func TestAuthorRepositoryFindAll(t *testing.T) {
	env := config.InitEnvironmentWithFiles("../../.env.testing")
	db := test.InitTestingDatabase(env)
	repo := repository.NewAuthorRepository(db)

	if _, err := repo.Save(dto.AuthorDto{
		FirstName: "John",
		LastName:  "Doe",
		Biography: "Bio 1",
	}); err != nil {
		t.Fatalf("failed to seed author: %v", err)
	}

	if _, err := repo.Save(dto.AuthorDto{
		FirstName: "Jane",
		LastName:  "Smith",
		Biography: "Bio 2",
	}); err != nil {
		t.Fatalf("failed to seed author: %v", err)
	}

	authors := repo.FindAll()
	if len(authors) != 2 {
		t.Fatalf("expected 2 authors, got %d", len(authors))
	}
}

func TestAuthorRepositoryFind(t *testing.T) {
	env := config.InitEnvironmentWithFiles("../../.env.testing")
	db := test.InitTestingDatabase(env)
	repo := repository.NewAuthorRepository(db)

	saved, err := repo.Save(dto.AuthorDto{
		FirstName: "Alice",
		LastName:  "Walker",
		Biography: "Bio",
	})
	if err != nil {
		t.Fatalf("failed to seed author: %v", err)
	}

	publisher := dao.Publisher{
		Name:    "Awesome Publisher",
		Address: "123 Street",
		Phone:   "08123456789",
		Website: "https://publisher.test",
	}
	if err := db.Create(&publisher).Error; err != nil {
		t.Fatalf("failed to create publisher: %v", err)
	}

	book := dao.Book{
		Title:       "Sample Book",
		Description: "A nice book",
		ISBN:        "ISBN-123",
		PublishedAt: time.Now(),
		PublisherID: publisher.ID,
	}
	if err := db.Create(&book).Error; err != nil {
		t.Fatalf("failed to create book: %v", err)
	}

	if err := db.Model(&book).Association("Author").Append(saved); err != nil {
		t.Fatalf("failed to associate author and book: %v", err)
	}

	result, err := repo.Find(saved.ID.String())
	if err != nil {
		t.Fatalf("expected to find author, got error: %v", err)
	}
	if result.ID == nil || result.ID.String() != saved.ID.String() {
		t.Fatalf("expected to get author with ID %s, got %v", saved.ID.String(), result.ID)
	}
	if len(result.Book) != 1 {
		t.Fatalf("expected associated book to be preloaded, got %d records", len(result.Book))
	}
}

func TestAuthorRepositoryFindNotFound(t *testing.T) {
	env := config.InitEnvironmentWithFiles("../../.env.testing")
	db := test.InitTestingDatabase(env)
	repo := repository.NewAuthorRepository(db)

	_, err := repo.Find(uuid.NewString())
	if err == nil {
		t.Fatal("expected error when author does not exist")
	}

	if err.Error() != "record not found" {
		t.Fatalf("expected 'record not found' error, got %v", err)
	}
}

func TestAuthorRepositorySaveUpdate(t *testing.T) {
	env := config.InitEnvironmentWithFiles("../../.env.testing")
	db := test.InitTestingDatabase(env)
	repo := repository.NewAuthorRepository(db)

	saved, err := repo.Save(dto.AuthorDto{
		FirstName: "Mary",
		LastName:  "Shelley",
		Biography: "Original bio",
	})
	if err != nil {
		t.Fatalf("failed to create author: %v", err)
	}

	updated, err := repo.Save(dto.AuthorDto{
		ID:        saved.ID,
		FirstName: saved.FirstName,
		LastName:  "Austin",
		Biography: "Updated bio",
	})
	if err != nil {
		t.Fatalf("failed to update author: %v", err)
	}

	if updated.LastName != "Austin" {
		t.Fatalf("expected last name to be updated, got %s", updated.LastName)
	}
	if updated.Biography != "Updated bio" {
		t.Fatalf("expected biography to be updated, got %s", updated.Biography)
	}
}

func TestAuthorRepositoryDelete(t *testing.T) {
	env := config.InitEnvironmentWithFiles("../../.env.testing")
	db := test.InitTestingDatabase(env)
	repo := repository.NewAuthorRepository(db)

	saved, err := repo.Save(dto.AuthorDto{
		FirstName: "Leo",
		LastName:  "Tolstoy",
		Biography: "Writer",
	})
	if err != nil {
		t.Fatalf("failed to create author: %v", err)
	}

	if err := repo.Delete(saved.ID.String()); err != nil {
		t.Fatalf("failed to delete author: %v", err)
	}

	if _, err := repo.Find(saved.ID.String()); err == nil {
		t.Fatal("expected error when fetching deleted author")
	}
}
