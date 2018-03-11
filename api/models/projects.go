package models

// Project represents one project
type Project struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	Verified    bool   `json:"verified"`
	Category    string `json:"category`
	UserID      int    `json:"userId"`
}

// AllProjects gets all projects in db
func (db *DB) AllProjects() ([]*Project, error) {
	rows, err := db.Query(`
        SELECT id, title, description, start_time, end_time,
               verified, category, user_id
        FROM projects`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	projects := []*Project{}
	for rows.Next() {
		var p Project
		if err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Description,
			&p.StartTime,
			&p.EndTime,
			&p.Verified,
			&p.Category,
			&p.UserID,
		); err != nil {
			return nil, err
		}
		projects = append(projects, &p)
	}

	return projects, nil
}

/*
func (p *project) getProject(db *sql.DB) error {
	return db.QueryRow("SELECT name, price FROM projects WHERE id=$1",
		p.ID).Scan(&p.Name, &p.Price)
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
		p.Name, p.Price).Scan(&p.ID)
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
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}

	return projects, nil
}
*/
