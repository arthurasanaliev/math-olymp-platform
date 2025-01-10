package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/jackc/pgx/v4"
)

const (
	dbUser = "postgres"
	dbHost = "localhost"
	dbPort = "5432"
	dbName = "math_olymp"
)

// connectToDB connects to the database
func connectToDB() (*pgx.Conn, error) {
	var password string
	fmt.Print("Enter PostgreSQL password: ")
	_, err := fmt.Scan(&password)
	if err != nil {
		return nil, fmt.Errorf("unable to read password: %v", err)
	}

	passwordStr := strings.TrimSpace(password)
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, passwordStr, dbHost, dbPort, dbName)
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}
	fmt.Println("Successfully connected to the database")
	return conn, nil
}

// createUsersTable creates the users table if it doesn't exist
func createUsersTable(conn *pgx.Conn) error {
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(50) NOT NULL,
			password VARCHAR(255) NOT NULL
		);
	`

	_, err := conn.Exec(context.Background(), createTableQuery)
	if err != nil {
		return fmt.Errorf("unable to create users table: %v", err)
	}
	fmt.Println("Users table created (if not already existing)")
	return nil
}

// createUser inserts a user into the users table
func createUser(conn *pgx.Conn, username, password string) error {
	insertQuery := `
		INSERT INTO users (username, password)
		VALUES ($1, $2);
	`

	_, err := conn.Exec(context.Background(), insertQuery, username, password)
	if err != nil {
		return fmt.Errorf("unable to insert user: %v", err)
	}
	fmt.Printf("User %s created successfully!\n", username)
	return nil
}

// getUsers retrieves all users from the database
func getUsers(conn *pgx.Conn) error {
	rows, err := conn.Query(context.Background(), "SELECT id, username FROM users")
	if err != nil {
		return fmt.Errorf("unable to fetch users: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var username string
		if err := rows.Scan(&id, &username); err != nil {
			return fmt.Errorf("failed to scan user: %v", err)
		}
		fmt.Printf("ID: %d, Username: %s\n", id, username)
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("rows error: %v", err)
	}

	return nil
}

// main is a driver function
func main() {
	conn, err := connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	if err := createUsersTable(conn); err != nil {
		log.Fatal(err)
	}

	if err := createUser(conn, "elnazar", "0000"); err != nil {
		log.Fatal(err)
	}

	if err := getUsers(conn); err != nil {
		log.Fatal(err)
	}
}
