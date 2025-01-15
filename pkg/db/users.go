package db

import (
	"context"
    "golang.org/x/crypto/bcrypt"
	"fmt"
)

// CheckUserExists checks if a user exists in the database
func (db *DB) CheckUserExists(username string) (bool, error) {
	query := `SELECT COUNT(*) FROM users WHERE username = $1`

	var count int
	err := db.Conn.QueryRow(context.Background(), query, username).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("database query failed: %v", err)
	}

	return count > 0, nil
}

// CheckUserPassword checks if a user's password is correct
func (db *DB) CheckUserPassword(username, password string) (bool, error) {
	query := `SELECT password FROM users WHERE username = $1`

	var storedPassword string
	err := db.Conn.QueryRow(context.Background(), query, username).Scan(&storedPassword)
	if err != nil {
		return false, fmt.Errorf("database query failed: %v", err)
	}
    err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
    if err != nil {
        return false, nil
    }

	return true, nil
}

// InsertUser inserts a new user into the database
func (db *DB) InsertUser(username, hashedPassword string) error {
	query := `INSERT INTO users (username, password) VALUES ($1, $2)`

	_, err := db.Conn.Exec(context.Background(), query, username, hashedPassword)
	if err != nil {
		return fmt.Errorf("unable to insert user: %v", err)
	}

	return nil
}
