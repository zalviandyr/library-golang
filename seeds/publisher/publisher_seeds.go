package publisher

import (
	"library-be/app/domain/dao"

	"github.com/go-faker/faker/v4"
	"gorm.io/gorm"
)

func PublisherSeeds(db *gorm.DB, count int) {
	publishers := make([]dao.Publisher, count)

	for i := range publishers {
		address := faker.GetRealAddress()

		publishers[i] = dao.Publisher{
			Name:    faker.Name(),
			Address: address.Address,
			Phone:   faker.Phonenumber(),
			Website: faker.URL(),
		}
	}

	if err := db.Create(&publishers).Error; err != nil {
		panic(err)
	}
}
