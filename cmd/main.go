package main

import (
	"os"

	"github.com/szymon676/jobguru/internal/app"
)

func main() {
	port := os.Getenv("PORT")
	app.SetupApp(port)
}
