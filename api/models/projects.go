package models

import (
	"time"
)

// Project represents one project
type Project struct {
	ID             int       `json:"id"`
	Title          string    `json:"title"`
	UserID         int       `json:"userId"`
	Category       string    `json:"category"`
	Description    string    `json:"description"`
	Verified       bool      `json:"verified"`
	Image          string    `json:"image"`
	AmountRaised   float64   `json:"amountRaised"`
	AmountRequired float64   `json:"amountRequired"`
	StartTime      time.Time `json:"startTime"`
	EndTime        time.Time `json:"endTime"`
}

// AllProjects gets all projects in db
func (db *DB) AllProjects() ([]*Project, error) {
	rows, err := db.Query(`
		select * from all_projects()
		`)
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
	`, title, userID, category, description, image, amountRequired, endTime
	).Scan(&id)
	return id, err
}

// GetProject returns all relevant data pertaining to the project
func (db *DB) GetProject(id string) (*Project, error) {
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
		&project.Image,
		&project.AmountRaised,
		&project.AmountRequired,
		&project.StartTime,
		&project.EndTime,
	)
	return &project, err
}

/*
func (p *project) getProject(db *sql.DB) error {
	return db.QueryRow("SELECT name, price FROM projects WHERE id=$1",
		p.ID).Scan(&project.Name, &project.Price)
}

func (p *project) updateProject(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE projects SET name=$1, price=$2 WHERE id=$3",
			p.Name, p.Price, p.ID)

	return err
}

func (p *project) deleteProject(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM projects WHERE id=$1", p.ID)

	return err
}

func (p *project) createProject(db *sql.DB) error {
	return db.QueryRow(
		"INSERT INTO projects(name, price) VALUES($1, $2) RETURNING id",
		p.Name, p.Price).Scan(&project.ID)
}

func getProjects(db *sql.DB, start, count int) ([]project, error) {
	rows, err := db.Query(
		"SELECT id, name,  price FROM projects LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	projects := []project{}

	for rows.Next() {
		var p project
		if err := rows.Scan(&project.ID, &project.Name, &project.Price); err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}

	return projects, nil
}
*/
