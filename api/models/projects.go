package models

import (
	"time"
)

// Project represents one project
type Project struct {
	ID             int       `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	StartTime      time.Time `json:"startTime"`
	EndTime        time.Time `json:"endTime"`
	Verified       bool      `json:"verified"`
	AmountRequired float64   `json:"amountRequired"`
	AmountRaised   float64   `json:"amountRaised"`
	Category       string    `json:"category"`
	UserID         int       `json:"userId"`
}

// AllProjects gets all projects in db
func (db *DB) AllProjects() ([]*Project, error) {
	rows, err := db.Query(`
        WITH total_payments AS (
            SELECT user_id, sum(amount) AS raised
            FROM payments
            GROUP BY user_id
        )
        SELECT p.id, p.title, p.description,
               p.start_time, p.end_time,
               p.amount_required, COALESCE(f.raised, 0) as amount_raised,
               p.verified, p.category, p.user_id
        FROM projects p LEFT OUTER JOIN total_payments f
        ON p.user_id = f.user_id`)
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
			&project.Description,
			&project.StartTime,
			&project.EndTime,
			&project.AmountRequired,
			&project.AmountRaised,
			&project.Verified,
			&project.Category,
			&project.UserID,
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
	description string,
	amountRequired float64,
	endTime time.Time,
	category string,
	userID int,
) (int, error) {
	id := 0
	err := db.QueryRow(`
        INSERT INTO projects (title, description, amount_required, end_time, category, user_id)
        VALUES($1, $2, $3, $4, $5, $6)
        RETURNING id
    `, title, description, amountRequired, endTime, category, userID,
	).Scan(&id)
	return id, err
}

// GetProject returns all relevant data pertaining to the project
func (db *DB) GetProject(id string) (*Project, error) {
	var project Project
	err := db.QueryRow(`
        SELECT p.id, p.title, p.description,
               p.start_time, p.end_time,
               p.amount_required, COALESCE(SUM(f.amount), 0) as amount_raised,
               p.verified, p.category, p.user_id
        FROM projects p LEFT OUTER JOIN payments f
        ON p.user_id = f.user_id
        WHERE p.id = $1
        GROUP BY p.id, p.user_id
    `, id,
	).Scan(
		&project.ID,
		&project.Title,
		&project.Description,
		&project.StartTime,
		&project.EndTime,
		&project.AmountRequired,
		&project.AmountRaised,
		&project.Verified,
		&project.Category,
		&project.UserID,
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
