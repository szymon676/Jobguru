package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/szymon676/jobguru/api/handlers"
	"github.com/szymon676/jobguru/api/service"
	"github.com/szymon676/jobguru/storage"

	"github.com/szymon676/jobguru/api/routes"
)

var (
	dsn = os.Getenv("DATABASE_URL")
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
}

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
