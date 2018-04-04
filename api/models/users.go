package models

// User represents one user
type User struct {
	ID             int       `json:"id"`
	Email          string    `json:"email"`
	Username       string    `json:"username"`
	TotalDonation  float64   `json:"totalDonation"`
	Image          string    `json:"image"`
	IsAdmin        bool      `json:"isAdmin"`
	// Password       string    `json:"password,omitempty"`
}

// AllUsers gets all users in db
func (db *DB) AllUsers() ([]*User, error) {
	rows, err := db.Query(`
		select * from all_users()
		`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*User{}
	for rows.Next() {
		var u User
		if err := rows.Scan(
			&u.ID,
			&u.Email,
			&u.Username,
			&u.TotalDonation,
			&u.Image,
			&u.IsAdmin,
		); err != nil {
			return nil, err
		}
		users = append(users, &u)
	}

	return users, nil
}

// GetUser returns the particular user given email and password
func (db *DB) GetUser(email string, password string) (*User, error) {
	var u User
	err := db.QueryRow(`
        select * from get_user($1, $2)
    `, email, password,
	).Scan(			
		&u.ID,
		&u.Email,
		&u.Username,
		&u.TotalDonation,
		&u.Image,
		&u.IsAdmin,
	)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// CreateUser creates a non-admin user given email and password
func (db *DB) CreateUser(
    email string, 
    password string,
    username string,
    image string,
) (int, error) {
	id := 0
	err := db.QueryRow(`
		select create_user($1, $2, $3, $4)	
		`,email, password, username, image
	).Scan(&id)
	return id, err
}
