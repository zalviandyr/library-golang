package author

import (
	"library-be/app/domain/dao"

	"github.com/go-faker/faker/v4"
	"gorm.io/gorm"
)

func AuthorSeeds(db *gorm.DB, count int) {
	authors := make([]dao.Author, count)

	for i := range authors {
		authors[i] = dao.Author{
			FirstName: faker.FirstName(),
			LastName:  faker.LastName(),
			Biography: faker.Paragraph(),
		}
	}

	if err := db.Create(&authors).Error; err != nil {
		panic(err)
	}
}
