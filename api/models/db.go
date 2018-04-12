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
	AllStatistics() ([]*Statistic, error)

	AllProjects(numPerPage, pageIdx int) ([]*Project, error)
	SearchProjects(searchTerm string, numPerPage, pageIdx int) ([]*Project, error)
	CreateProject(
		title string,
		userID int,
		category string,
		description string,
		image string,
		amountRequired float64,
		endTime time.Time,
	) (int, error)
	ReplaceProject(
		id int,
		title string,
		userID int,
		category string,
		description string,
		image string,
		verified bool,
		amountRequired float64,
		endTime time.Time,
	) error
	GetProject(id int) (*Project, error)

	AllUsers(numPerPage, pageIdx int) ([]*User, error)
	GetUser(email string, password string) (*User, error)
	CreateUser(email string, password string, username string, image string) (*User, error)

	AllCategories() ([]*Category, error)
	CreateCategory(name string) error

	AllTags() ([]*Tag, error)
	CreateTag(projectID int, tagName string) error
	AllProjectTags(projectID int) ([]*Tag, error)

	AllPayments(numPerPage, pageIdx int) ([]*Payment, error)
	AllProjectPayments(projectID int, numPerPage, pageIdx int) ([]*Payment, error)
	AllUserPayments(userID int, numPerPage, pageIdx int) ([]*Payment, error)
	CreatePayment(userID int, projectID int, amount float64) (int, error)
	UpdatePayment(paymentID int, amount float64) error
	DeletePayment(paymentID int) error

	AllComments(numPerPage, pageIdx int) ([]*Comment, error)
	AllProjectComments(projectID, numPerPage, pageIdx int) ([]*Comment, error)
	AllUserComments(userID, numPerPage, pageIdx int) ([]*Comment, error)
	CreateComment(userID int, projectID int, content string) (int, error)
	UpdateComment(commentID int, content string) error
	DeleteComment(commentID int) error
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

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	// Try to connect every second
	for i := 0; i < 10; i++ {
		log.Printf("Attempting connection: %d", i)
		err = db.Ping()
		if err == nil {
			return &DB{db}, nil
		}

		time.Sleep(time.Second * time.Duration(i))
	}
	return nil, err
}
