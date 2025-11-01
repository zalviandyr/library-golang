package config

import (
	"fmt"
	"library-be/app/domain/dao"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDatabase(env Environment) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		env.DB_HOST, env.DB_USER, env.DB_PASSWORD, env.DB_DATABASE, env.DB_PORT,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Failed to connect database")
	}

	// create uuid_generate_v4()
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	// Migrate the schema
	db.AutoMigrate(
		&dao.Author{},
		&dao.Publisher{},
		&dao.Book{},
	)

	return db
}
