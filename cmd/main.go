package main

import (
	"log"

	"github.com/beadex/espresso/lib/backend"
	"github.com/beadex/espresso/lib/gui"
)

func main() {
	log.Println("Starting Espresso...")

	// Initialize backend
	b := backend.Initialize()

	// Initialize GUI
	g := gui.Initialize(b)

	// Run the application
	if err := g.Run(); err != nil {
		log.Fatalf("Application failed: %v", err)
	}
}
