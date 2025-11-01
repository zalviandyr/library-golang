package test

import (
	"fmt"
	"library-be/app/domain/dao"
	"library-be/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitTestingDatabase(env config.Environment) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		env.DB_HOST, env.DB_USER, env.DB_PASSWORD, env.DB_DATABASE, env.DB_PORT,
	)
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic("Failed to connect database")
	}

	// create uuid_generate_v4()
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	resetDatabase(db)

	// Migrate the schema
	db.AutoMigrate(
		&dao.Author{},
		&dao.Publisher{},
		&dao.Book{},
	)

	return db
}

func resetDatabase(db *gorm.DB) {
	tables, err := db.Migrator().GetTables()
	if err != nil {
		log.Printf("⚠️ Failed to get tables: %v\n", err)
		return
	}

	for _, table := range tables {
		if err := db.Migrator().DropTable(table); err != nil {
			log.Printf("⚠️ Failed to drop table %s: %v\n", table, err)
		}
	}

	log.Println("🧹 All tables dropped successfully.")
}
