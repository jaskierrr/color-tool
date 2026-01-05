package main

import (
	"log"

	"github.com/jaskierrr/color-tool/internal"
)

func main() {
	err := app.Run()
	if err != nil {
		log.Fatalf("Failed launch app. Error: %v", err)
	}
}
