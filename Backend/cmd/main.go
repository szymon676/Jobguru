package main

import (
	"log"

	"github.com/szymon676/jobguru/api/handlers"
	"github.com/szymon676/jobguru/storage"

	"github.com/szymon676/jobguru/api/routes"
)

func main() {
	store, err := storage.NewPostgreStorage("some dsn, soon...")
	if err != nil {
		log.Fatal(err)
	}

	js := storage.NewPostgreJobStorage(store)
	us := storage.NewPostgreUserStorage(store)
	jh := handlers.NewJobHandler(js)
	ah := handlers.NewAuthHandler(us)

	err = routes.SetupRoutes(":4000", ah, jh)
	if err != nil {
		log.Fatal(err)
	}
}
