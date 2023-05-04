package main

import (
	"fmt"
	"log"

	"github.com/szymon676/jobguru/api/handlers"
	"github.com/szymon676/jobguru/service"
	"github.com/szymon676/jobguru/storage"

	"github.com/szymon676/jobguru/api/routes"
)

var dsn = fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable", "postgres", "1234", "jobguru")

func main() {
	store, err := storage.NewPostgreStorage(dsn)
	if err != nil {
		log.Fatal(err)
	}

	jobStore := storage.NewPostgreJobStorage(store)
	userStore := storage.NewPostgreUserStorage(store)
	jobService := service.NewJobService(jobStore)
	userService := service.NewUserService(userStore)

	jh := handlers.NewJobHandler(jobService)
	uh := handlers.NewUsersHandler(userService)

	err = routes.SetupRoutes(":3000", uh, jh)
	if err != nil {
		log.Fatal(err)
	}
}
