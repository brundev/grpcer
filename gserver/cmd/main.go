package main

import (
	"gserver/internal/database" // Make sure to replace 'your-project' with your actual module name from go.mod
	"log"
)

// DataStore is the interface that defines the database operations our application needs.
// It is defined in the 'main' package because this is where it is consumed.
// This decouples our main logic from the specific database implementation.
type DataStore interface {
	InsertItem(name string, description string) error
	GetItems() ([]database.Item, error)
}

func main() {
	// The database package provides a function to create a new, concrete implementation.
	// We receive a concrete type (*database.DB) that satisfies our DataStore interface.
	db, err := database.New("storage.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close() // Ensure the database connection is closed when main exits.

	// We can now pass our concrete db object to any function that expects a DataStore interface.
	runApp(*db)

}

func runApp(database database.DB) {

}
