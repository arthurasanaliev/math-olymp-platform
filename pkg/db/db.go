package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"strings"
)

const (
	dbUser = "postgres"
	dbHost = "localhost"
	dbPort = "5432"
	dbName = "math_olymp"
)

// DB represents database connection
type DB struct {
	Conn *pgx.Conn
}

// NewDB returns a new DB connection
func NewDB() (*DB, error) {
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
	return &DB{Conn: conn}, nil
}

// Close closes the database connection
func (db *DB) Close() error {
	return db.Conn.Close(context.Background())
}
