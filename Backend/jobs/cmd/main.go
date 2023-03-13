package main

import (
	"flag"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/szymon676/job-guru/jobs/internal/database"
	"github.com/szymon676/job-guru/jobs/internal/handlers"
)

func main() {
	dsn := flag.String("dsn", "host=localhost user=postgres password=1234 dbname=jobguru-jobs port=5432 sslmode=disable", "dsn to connect to the database")
	flag.Parse()

	if _, err := database.NewDatabase("postgres", *dsn); err != nil {
		panic(err)
	}

	server := handlers.NewApiServer(":1337")
	server.Run()
}