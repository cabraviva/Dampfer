package database

import (
	"Dampfer/utils"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// Initialize the database connection and create the users table if it doesn't exist
func init() {
	var err error

	DB, err = sql.Open("sqlite3", "./Dampfer.db")
	if err != nil {
		utils.Log.Panic("Failed to open database: ", err)
		panic(err)
	}

	// Create the users table with permissions
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS users (
        username TEXT PRIMARY KEY,
        password_hash TEXT NOT NULL,
        permission TEXT NOT NULL CHECK (permission IN ('system-admin', 'admin', 'insight'))
    )`)
	if err != nil {
		utils.Log.Panic("Failed to create users table: ", err)
		panic(err)
	}
}
