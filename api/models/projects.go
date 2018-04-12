package models

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

// Project represents one project
type Project struct {
	ID             int       `json:"id"`
	Title          string    `json:"title"`
	UserID         int       `json:"userId"`
	Category       string    `json:"category"`
	Description    string    `json:"description"`
	Verified       bool      `json:"verified"`
	Tags           []string  `json:"tags"`
	Image          string    `json:"image"`
	AmountRaised   float64   `json:"amountRaised"`
	AmountRequired float64   `json:"amountRequired"`
	StartTime      time.Time `json:"startTime"`
	EndTime        time.Time `json:"endTime"`
}

func readProjects(rows *sql.Rows, err error) ([]*Project, error) {
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	projects := []*Project{}
	for rows.Next() {
		var project Project
		if err := rows.Scan(
			&project.ID,
			&project.Title,
			&project.UserID,
			&project.Category,
			&project.Description,
			&project.Verified,
			pq.Array(&project.Tags),
			&project.Image,
			&project.AmountRaised,
			&project.AmountRequired,
			&project.StartTime,
			&project.EndTime,
		); err != nil {
			return nil, err
		}
		projects = append(projects, &project)
	}
	return projects, nil
}

// AllProjects gets all projects in db
func (db *DB) AllProjects(numPerPage int, pageIdx int) ([]*Project, error) {
	return readProjects(db.Query(`
		select * from all_projects($1, $2)
		`, numPerPage, pageIdx))
}

// SearchProjects does a full text search on all projects in db
func (db *DB) SearchProjects(searchTerm string, numPerPage int, pageIdx int) ([]*Project, error) {
	return readProjects(db.Query(`
		select * from search_projects($1, $2, $3)
		`, searchTerm, numPerPage, pageIdx))
}

// CreateProject creates a funding project
func (db *DB) CreateProject(
	title string,
	userID int,
	category string,
	description string,
	image string,
	amountRequired float64,
	endTime time.Time,
) (int, error) {
	id := 0
	err := db.QueryRow(`
        select create_project($1, $2, $3, $4, $5, $6, $7)
	`, title, userID, category, description, image, amountRequired, endTime,
	).Scan(&id)
	return id, err
}

// ReplaceProject Replaces a funding project
func (db *DB) ReplaceProject(
	id int,
	title string,
	userID int,
	category string,
	description string,
	image string,
	verified bool,
	amountRequired float64,
	endTime time.Time,
) error {
	_, err := db.Exec(`
        select update_project($1, $2, $3, $4, $5, $6, $7, $8, $9)
    `, id,
		title,
		userID,
		category,
		description,
		image,
		verified,
		amountRequired,
		endTime,
	)
	return err
}

// GetProject returns all relevant data pertaining to the project
func (db *DB) GetProject(id int) (*Project, error) {
	var project Project
	err := db.QueryRow(`
        select * from get_project($1)
    `, id,
	).Scan(
		&project.ID,
		&project.Title,
		&project.UserID,
		&project.Category,
		&project.Description,
		&project.Verified,
		pq.Array(&project.Tags),
		&project.Image,
		&project.AmountRaised,
		&project.AmountRequired,
		&project.StartTime,
		&project.EndTime,
	)
	return &project, err
}
