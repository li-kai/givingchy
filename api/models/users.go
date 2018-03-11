package models

// User represents one user
type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"isAdmin"`
}

// AllUsers gets all users in db
func (db *DB) AllUsers() ([]*User, error) {
	rows, err := db.Query(`
        SELECT id, email, is_admin
        FROM users`)
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
        SELECT id, email, is_admin FROM users
        WHERE email = $1
        AND password = crypt($2, password)
    `, email, password,
	).Scan(&u.ID, &u.Email, &u.IsAdmin)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// CreateUser creates a non-admin user given email and password
func (db *DB) CreateUser(email string, password string) (int, error) {
	id := 0
	err := db.QueryRow(`
        INSERT INTO users (email, password)
        VALUES($1, crypt($2, gen_salt('bf', 8)))
        RETURNING id
    `, email, password,
	).Scan(&id)
	return id, err
}
