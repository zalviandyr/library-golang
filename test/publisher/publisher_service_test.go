package publisher_test

import (
	"testing"

	"library-be/app/domain/dto"
	"library-be/app/repository"
	"library-be/app/service"
	"library-be/config"
	"library-be/test"
)

func TestPublisherServiceFindAll(t *testing.T) {
	env := config.InitEnvironment("../../.env.testing")
	db := test.InitTestingDatabase(env)
	repo := repository.NewPublisherRepository(db)
	service := service.NewPublisherService(repo)

	if _, err := service.Save(dto.PublisherDto{
		Name:    "Publisher One",
		Address: "Address One",
		Phone:   "08111111111",
		Website: "https://one.test",
	}); err != nil {
		t.Fatalf("failed to create publisher: %v", err)
	}

	if _, err := service.Save(dto.PublisherDto{
		Name:    "Publisher Two",
		Address: "Address Two",
		Phone:   "08222222222",
		Website: "https://two.test",
	}); err != nil {
		t.Fatalf("failed to create publisher: %v", err)
	}

	publishers := service.FindAll()
	if len(publishers) != 2 {
		t.Fatalf("expected 2 publishers, got %d", len(publishers))
	}
}

func TestPublisherServiceFind(t *testing.T) {
	env := config.InitEnvironment("../../.env.testing")
	db := test.InitTestingDatabase(env)
	repo := repository.NewPublisherRepository(db)
	service := service.NewPublisherService(repo)

	created, err := service.Save(dto.PublisherDto{
		Name:    "Publisher",
		Address: "Street 1",
		Phone:   "08123456789",
		Website: "https://publisher.test",
	})
	if err != nil {
		t.Fatalf("failed to create publisher: %v", err)
	}

	result, err := service.Find(created.ID.String())
	if err != nil {
		t.Fatalf("expected find to succeed: %v", err)
	}
	if result.ID == nil || result.ID.String() != created.ID.String() {
		t.Fatalf("expected publisher ID %s, got %v", created.ID.String(), result.ID)
	}
}

func TestPublisherServiceSaveUpdateAndDelete(t *testing.T) {
	env := config.InitEnvironment("../../.env.testing")
	db := test.InitTestingDatabase(env)
	repo := repository.NewPublisherRepository(db)
	service := service.NewPublisherService(repo)

	created, err := service.Save(dto.PublisherDto{
		Name:    "Initial Publisher",
		Address: "Old Address",
		Phone:   "08123456789",
		Website: "https://initial.test",
	})
	if err != nil {
		t.Fatalf("failed to create publisher: %v", err)
	}

	updated, err := service.Save(dto.PublisherDto{
		ID:      created.ID,
		Name:    "Updated Publisher",
		Address: "New Address",
		Phone:   "08987654321",
		Website: "https://updated.test",
	})
	if err != nil {
		t.Fatalf("failed to update publisher: %v", err)
	}
	if updated.Name != "Updated Publisher" {
		t.Fatalf("expected updated name, got %s", updated.Name)
	}

	if err := service.Delete(created.ID.String()); err != nil {
		t.Fatalf("failed to delete publisher: %v", err)
	}

	if _, err := service.Find(created.ID.String()); err == nil {
		t.Fatal("expected error when fetching deleted publisher")
	}
}
