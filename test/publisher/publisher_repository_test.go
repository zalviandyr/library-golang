package publisher_test

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

func TestPublisherRepositoryFindAll(t *testing.T) {
	env := config.InitEnvironment("../../.env.testing")
	db := test.InitTestingDatabase(env)
	repo := repository.NewPublisherRepository(db)

	if _, err := repo.Save(dto.PublisherDto{
		Name:    "Publisher One",
		Address: "Address One",
		Phone:   "08111111111",
		Website: "https://one.test",
	}); err != nil {
		t.Fatalf("failed to seed publisher: %v", err)
	}

	if _, err := repo.Save(dto.PublisherDto{
		Name:    "Publisher Two",
		Address: "Address Two",
		Phone:   "08222222222",
		Website: "https://two.test",
	}); err != nil {
		t.Fatalf("failed to seed publisher: %v", err)
	}

	publishers := repo.FindAll()
	if len(publishers) != 2 {
		t.Fatalf("expected 2 publishers, got %d", len(publishers))
	}
}

func TestPublisherRepositoryFind(t *testing.T) {
	env := config.InitEnvironment("../../.env.testing")
	db := test.InitTestingDatabase(env)
	repo := repository.NewPublisherRepository(db)

	saved, err := repo.Save(dto.PublisherDto{
		Name:    "Publisher",
		Address: "Street 1",
		Phone:   "08123456789",
		Website: "https://publisher.test",
	})
	if err != nil {
		t.Fatalf("failed to seed publisher: %v", err)
	}

	book := dao.Book{
		Title:       "Sample Book",
		Description: "A nice book",
		ISBN:        "ISBN-123",
		PublishedAt: time.Now(),
		PublisherID: saved.ID,
	}
	if err := db.Create(&book).Error; err != nil {
		t.Fatalf("failed to create book: %v", err)
	}

	result, err := repo.Find(saved.ID.String())
	if err != nil {
		t.Fatalf("expected to find publisher, got error: %v", err)
	}
	if result.ID == nil || result.ID.String() != saved.ID.String() {
		t.Fatalf("expected publisher ID %s, got %v", saved.ID.String(), result.ID)
	}
	if len(result.Book) != 1 {
		t.Fatalf("expected associated book to be preloaded, got %d records", len(result.Book))
	}
}

func TestPublisherRepositoryFindNotFound(t *testing.T) {
	env := config.InitEnvironment("../../.env.testing")
	db := test.InitTestingDatabase(env)
	repo := repository.NewPublisherRepository(db)

	_, err := repo.Find(uuid.NewString())
	if err == nil {
		t.Fatal("expected error when publisher does not exist")
	}

	if err.Error() != "record not found" {
		t.Fatalf("expected 'record not found' error, got %v", err)
	}
}

func TestPublisherRepositorySaveUpdate(t *testing.T) {
	env := config.InitEnvironment("../../.env.testing")
	db := test.InitTestingDatabase(env)
	repo := repository.NewPublisherRepository(db)

	saved, err := repo.Save(dto.PublisherDto{
		Name:    "Initial Publisher",
		Address: "Old Address",
		Phone:   "08123456789",
		Website: "https://initial.test",
	})
	if err != nil {
		t.Fatalf("failed to create publisher: %v", err)
	}

	updated, err := repo.Save(dto.PublisherDto{
		ID:      saved.ID,
		Name:    "Updated Publisher",
		Address: "New Address",
		Phone:   "08987654321",
		Website: "https://updated.test",
	})
	if err != nil {
		t.Fatalf("failed to update publisher: %v", err)
	}

	if updated.Name != "Updated Publisher" {
		t.Fatalf("expected name to be updated, got %s", updated.Name)
	}
	if updated.Address != "New Address" {
		t.Fatalf("expected address to be updated, got %s", updated.Address)
	}
	if updated.Phone != "08987654321" {
		t.Fatalf("expected phone to be updated, got %s", updated.Phone)
	}
	if updated.Website != "https://updated.test" {
		t.Fatalf("expected website to be updated, got %s", updated.Website)
	}
}

func TestPublisherRepositoryDelete(t *testing.T) {
	env := config.InitEnvironment("../../.env.testing")
	db := test.InitTestingDatabase(env)
	repo := repository.NewPublisherRepository(db)

	saved, err := repo.Save(dto.PublisherDto{
		Name:    "To Be Deleted",
		Address: "Delete Address",
		Phone:   "08123450000",
		Website: "https://delete.test",
	})
	if err != nil {
		t.Fatalf("failed to create publisher: %v", err)
	}

	if err := repo.Delete(saved.ID.String()); err != nil {
		t.Fatalf("failed to delete publisher: %v", err)
	}

	if _, err := repo.Find(saved.ID.String()); err == nil {
		t.Fatal("expected error when fetching deleted publisher")
	}
}
