package author_test

import (
	"testing"

	"library-be/app/domain/dto"
	"library-be/app/repository"
	"library-be/app/service"
	"library-be/config"
	"library-be/test"
)

func TestAuthorServiceFindAll(t *testing.T) {
	env := config.InitEnvironmentWithFiles("../../.env.testing")
	db := test.InitTestingDatabase(env)
	repo := repository.NewAuthorRepository(db)
	service := service.NewAuthorService(repo)

	if _, err := service.Save(dto.AuthorDto{
		FirstName: "John",
		LastName:  "Doe",
		Biography: "Bio 1",
	}); err != nil {
		t.Fatalf("failed to create author: %v", err)
	}

	if _, err := service.Save(dto.AuthorDto{
		FirstName: "Jane",
		LastName:  "Smith",
		Biography: "Bio 2",
	}); err != nil {
		t.Fatalf("failed to create author: %v", err)
	}

	authors := service.FindAll()
	if len(authors) != 2 {
		t.Fatalf("expected 2 authors, got %d", len(authors))
	}
}

func TestAuthorServiceFind(t *testing.T) {
	env := config.InitEnvironmentWithFiles("../../.env.testing")
	db := test.InitTestingDatabase(env)
	repo := repository.NewAuthorRepository(db)
	service := service.NewAuthorService(repo)

	created, err := service.Save(dto.AuthorDto{
		FirstName: "Alice",
		LastName:  "Walker",
		Biography: "Bio",
	})
	if err != nil {
		t.Fatalf("failed to create author: %v", err)
	}

	result, err := service.Find(created.ID.String())
	if err != nil {
		t.Fatalf("expected find to succeed: %v", err)
	}
	if result.ID == nil || result.ID.String() != created.ID.String() {
		t.Fatalf("expected author ID %s, got %v", created.ID.String(), result.ID)
	}
}

func TestAuthorServiceSaveUpdateAndDelete(t *testing.T) {
	env := config.InitEnvironmentWithFiles("../../.env.testing")
	db := test.InitTestingDatabase(env)
	repo := repository.NewAuthorRepository(db)
	service := service.NewAuthorService(repo)

	created, err := service.Save(dto.AuthorDto{
		FirstName: "Mary",
		LastName:  "Shelley",
		Biography: "Original",
	})
	if err != nil {
		t.Fatalf("failed to create author: %v", err)
	}

	updated, err := service.Save(dto.AuthorDto{
		ID:        created.ID,
		FirstName: created.FirstName,
		LastName:  "Austin",
		Biography: "Updated",
	})
	if err != nil {
		t.Fatalf("failed to update author: %v", err)
	}
	if updated.LastName != "Austin" {
		t.Fatalf("expected updated last name, got %s", updated.LastName)
	}

	if err := service.Delete(created.ID.String()); err != nil {
		t.Fatalf("failed to delete author: %v", err)
	}

	if _, err := service.Find(created.ID.String()); err == nil {
		t.Fatal("expected error when fetching deleted author")
	}
}
