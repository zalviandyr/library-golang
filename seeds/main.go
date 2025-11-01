package main

import (
	"library-be/config"
	"library-be/seeds/author"
	"library-be/seeds/book"
	"library-be/seeds/publisher"
)

func main() {
	count := 5
	env := config.InitEnvironment()
	db := config.InitDatabase(env)

	author.AuthorSeeds(db, count)
	publisher.PublisherSeeds(db, count)
	book.BookSeeds(db, count)
}
