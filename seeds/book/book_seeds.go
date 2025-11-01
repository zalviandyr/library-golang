package book

import (
	"library-be/app/domain/dao"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func BookSeeds(db *gorm.DB, count int) {
	// fetch publisher
	var publishers []dao.Publisher
	if err := db.Find(&publishers).Error; err != nil {
		panic(err)
	}
	if len(publishers) == 0 {
		panic("Publisher is empty")
	}

	// fetch authors
	var authors []dao.Author
	if err := db.Find(&authors).Error; err != nil {
		panic(err)
	}
	if len(authors) == 0 {
		panic("Author is empty")
	}

	books := make([]dao.Book, count)
	for i := range books {
		selectedPublisher := &publishers[rand.Intn(len(publishers))]
		selectedAuthor := getSelectedAuthors(authors)

		books[i] = dao.Book{
			Title:       faker.Name(),
			Description: faker.Paragraph(),
			ISBN:        fakeISBN(),
			PublishedAt: time.Now(),
			PublisherID: selectedPublisher.ID,
			Publisher:   selectedPublisher,
			Author:      selectedAuthor,
		}
	}

	if err := db.Create(&books).Error; err != nil {
		panic(err)
	}
}

func fakeISBN() string {
	digits, _ := faker.RandomInt(0, 9, 13) // 13 angka, tiap angka 0–9
	var b strings.Builder
	for _, d := range digits {
		b.WriteString(strconv.Itoa(d))
	}
	return b.String()
}

func getSelectedAuthors(authors []dao.Author) []dao.Author {
	nAuthors := rand.Intn(3) + 1
	selected := make([]dao.Author, 0, nAuthors)
	usedID := map[uuid.UUID]struct{}{}

	for len(selected) < nAuthors && len(selected) < len(authors) {
		candidate := authors[rand.Intn(len(authors))]
		if candidate.ID == nil {
			continue
		}

		if _, ok := usedID[*candidate.ID]; ok {
			continue
		}

		usedID[*candidate.ID] = struct{}{}
		selected = append(selected, candidate)
	}

	return selected
}
