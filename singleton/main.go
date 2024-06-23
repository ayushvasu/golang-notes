package main

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

// Singleton struct to hold the database connection pool
type Database struct {
	ConnectionPool *sql.DB
}

var (
	dbInstance *Database
	once       sync.Once
)

// GetInstance returns the singleton instance of Database
func GetInstance(maxOpenConns, maxIdleConns int) *Database {
	once.Do(func() {
		// Initialize the singleton instance
		dbInstance = &Database{}
		dbInstance.initConnectionPool(maxOpenConns, maxIdleConns)
	})
	return dbInstance
}

// initConnectionPool initializes the PostgreSQL connection pool
func (db *Database) initConnectionPool(maxOpenConns, maxIdleConns int) {
	// Replace with your PostgreSQL connection parameters
	connectionString := "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"

	// Open a database connection
	dbConn, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	// Check if the connection is established
	err = dbConn.Ping()
	if err != nil {
		panic(err)
	}

	// Set connection pool configuration
	dbConn.SetMaxOpenConns(maxOpenConns)
	dbConn.SetMaxIdleConns(maxIdleConns)

	// Set the connection pool in the singleton instance
	db.ConnectionPool = dbConn
}

// Example function to execute a query
func (db *Database) QueryExample() {
	rows, err := db.ConnectionPool.Query("SELECT username, full_name as fullName FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Process rows
	for rows.Next() {
		// Scan rows into variables
		var username string
		var fullName string
		err := rows.Scan(&username, &fullName)
		if err != nil {
			panic(err)
		}
		fmt.Printf("username: %s, Name: %s\n", username, fullName)
	}
}

func main() {
	// Get the singleton instance with specified connection pool sizes
	db := GetInstance(20, 10) // Example: maxOpenConns = 20, maxIdleConns = 10

	// Example usage
	db.QueryExample()
}
