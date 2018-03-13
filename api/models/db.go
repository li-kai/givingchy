package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	// Invoke pq's init method
	_ "github.com/lib/pq"
)

// Datastore defines all methods over models
type Datastore interface {
	AllProjects() ([]*Project, error)
	CreateProject(
		title string,
		description string,
		amountRequired float64,
		endTime time.Time,
		category string,
		userID int,
	) (int, error)
	GetProject(id string) (*Project, error)

	AllUsers() ([]*User, error)
	GetUser(email string, password string) (*User, error)
	CreateUser(email string, password string) (int, error)
	AllCategories() ([]*Category, error)
	CreateCategory(name string) error
}

// DB Wraps sql db for usage
type DB struct {
	*sql.DB
}

// NewDB initialises a new db instance
func NewDB() (*DB, error) {
	connectionString :=
		fmt.Sprintf("host=postgres sslmode=disable user=%s password=%s dbname=%s",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"),
		)

	var db *sql.DB
	var err error
	// Try to connect every second
	for i := 0; i < 30; i++ {
		log.Printf("Attempting connection: %d", i)
		db, err = sql.Open("postgres", connectionString)
		if err == nil {
			break
		}

		time.Sleep(time.Second)
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
