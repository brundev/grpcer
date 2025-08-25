package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	*sql.DB
}

type Item struct {
	ID          int
	Name        string
	Description string
}

func New(dataSourceName string) (*DB, error) {
	// Open the database file. It will be created if it doesn't exist.
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}

	// Ping the database to verify the connection is alive.
	if err = db.Ping(); err != nil {
		return nil, err
	}

	// SQL statement to create our table.
	// "IF NOT EXISTS" prevents an error if the table already exists.
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS items (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"name" TEXT,
		"description" TEXT
	);`

	// Execute the SQL statement.
	_, err = db.Exec(createTableSQL)
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func (db *DB) InsertItem(name string, description string) error {
	// Use a prepared statement to prevent SQL injection.
	insertSQL := "INSERT INTO items(name, description) VALUES (?, ?)"
	_, err := db.Exec(insertSQL, name, description)
	return err
}

func (db *DB) GetItems() ([]Item, error) {
	rows, err := db.Query("SELECT id, name, description FROM items ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create a slice to hold the results.
	var items []Item

	// Loop through the returned rows.
	for rows.Next() {
		var item Item
		// Scan the values from the current row into the 'item' struct.
		if err := rows.Scan(&item.ID, &item.Name, &item.Description); err != nil {
			return nil, err
		}
		// Add the item to our slice.
		items = append(items, item)
	}

	// Check for any errors that occurred during the iteration.
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}
