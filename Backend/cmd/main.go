package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/szymon676/jobguru/api/handlers"
	"github.com/szymon676/jobguru/service"
	"github.com/szymon676/jobguru/storage"

	"github.com/szymon676/jobguru/api/routes"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

var dsn = fmt.Sprintf("host=jobguru-db user=%s password=%s dbname=%s port=5432 sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

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
