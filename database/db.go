package database

import (
	"Dampfer/utils"
	"database/sql"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// Initialize the database connection and create the necessary tables
func init() {
	var err error

	// Get the home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		utils.Log.Panic("Error retrieving home directory: ", err)
		panic(err)
	}

	// Define the Dampfer folder and Dampfer.db file path
	dbDir := filepath.Join(homeDir, "Dampfer")
	dbFilePath := filepath.Join(dbDir, "Dampfer.db")

	// Create the directory if it doesn't exist
	err = os.MkdirAll(dbDir, 0755)
	if err != nil {
		utils.Log.Panic("Error creating database directory: ", err)
		panic(err)
	}

	// Open the SQLite database
	DB, err = sql.Open("sqlite3", dbFilePath)
	if err != nil {
		utils.Log.Panic("Failed to open database: ", err)
		panic(err)
	}

	// Create the necessary tables
	createUserTable()
	createAuthTable()
}

// createUserTable creates the users table if it doesn't already exist
func createUserTable() {
	// Create the users table with permission field
	query := `CREATE TABLE IF NOT EXISTS users (
        username TEXT PRIMARY KEY,
        password_hash TEXT NOT NULL,
        permission TEXT NOT NULL CHECK (permission IN ('system-admin', 'admin', 'insight'))
    )`
	_, err := DB.Exec(query)
	if err != nil {
		utils.Log.Panic("Failed to create users table: ", err)
		panic(err)
	}
}

// Creates table that stores jwt secret
func createAuthTable() {
	query := `CREATE TABLE IF NOT EXISTS auth (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        secret_key BLOB NOT NULL
    )`
	_, err := DB.Exec(query)
	if err != nil {
		utils.Log.Panic("Failed to create auth table: ", err)
		panic(err)
	}
}
