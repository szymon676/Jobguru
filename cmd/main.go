package main

import (
	"os"

	"github.com/szymon676/jobguru/internal/app"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	}
	dsn := os.Getenv("DSN")
	app.SetupApp(port, dsn)
}
