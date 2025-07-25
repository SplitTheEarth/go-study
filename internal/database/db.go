package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// InitDB initializes the database connection and creates necessary tables
func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite3", "./internal/studyapp.db")
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	// Test the connection
	if err = DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	// Create tables
	if err = createTables(); err != nil {
		return fmt.Errorf("failed to create tables: %w", err)
	}

	fmt.Println("Database initialized successfully")
	return nil
}

// createTables creates all necessary database tables
func createTables() error {
	createUsersTableSQL := `CREATE TABLE IF NOT EXISTS users (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"username" TEXT NOT NULL UNIQUE,
		"email" TEXT NOT NULL UNIQUE,
		"password_hash" TEXT NOT NULL,
		"score" INTEGER NOT NULL DEFAULT 0
	);`

	createDecksTableSQL := `CREATE TABLE IF NOT EXISTS decks (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"name" TEXT NOT NULL,
		"description" TEXT NOT NULL,
		"created_at" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`

	createQuestionsTableSQL := `CREATE TABLE IF NOT EXISTS questions (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"deck_id" INTEGER NOT NULL,
		"question_text" TEXT NOT NULL,
		"answer" TEXT NOT NULL,
		"options" TEXT NOT NULL, -- Can be NULL if it's not a multiple-choice question
		"created_at" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (deck_id) REFERENCES decks(id)
	);`

	fmt.Println("Creating users table...")
	_, err := DB.Exec(createUsersTableSQL)
	if err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}
	fmt.Println("Users table created successfully")

	fmt.Println("Creating decks table...")
	_, err = DB.Exec(createDecksTableSQL)
	if err != nil {
		return fmt.Errorf("failed to create decks table: %w", err)
	}
	fmt.Println("Decks table created successfully")

	fmt.Println("Creating questions table...")
	_, err = DB.Exec(createQuestionsTableSQL)
	if err != nil {
		return fmt.Errorf("failed to create questions table: %w", err)
	}
	fmt.Println("Questions table created successfully")

	return nil
}

// CloseDB closes the database connection
func CloseDB() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
