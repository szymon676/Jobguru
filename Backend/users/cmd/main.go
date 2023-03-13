package main

import (
	"flag"

	_ "github.com/lib/pq"
	"github.com/szymon676/job-guru/users/internal/database"
	"github.com/szymon676/job-guru/users/internal/handlers"
)

func main() {
	dsn := flag.String("dsn", "host=localhost user=postgres password=1234 dbname=jobguru-users port=5432 sslmode=disable", "dsn to connect to the database")
	flag.Parse()

	if _, err := database.NewDatabase("postgres", *dsn); err != nil {
		panic(err)
	}

	server := handlers.NewApiServer(":5000")
	server.Run()
}
