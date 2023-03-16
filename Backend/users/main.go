package main

import (
	"flag"

	_ "github.com/lib/pq"
	"github.com/szymon676/job-guru/users/api"
	"github.com/szymon676/job-guru/users/storage"
)

func main() {
	dsn := flag.String("dsn", "host=localhost user=postgres password=1234 dbname=jobguru-users port=5432 sslmode=disable", "dsn to connect to the database")
	flag.Parse()

	_, err := storage.NewDatabase("postgres", *dsn)
	if err != nil {
		panic(err)
	}

	ps, err := storage.NewPostgresStorage()
	if err != nil {
		panic(err)
	}

	server := api.NewApiServer(":5000", ps)
	server.Run()
}
