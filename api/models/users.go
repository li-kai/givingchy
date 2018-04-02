package models

// User represents one user
type User struct {
	ID             int       `json:"id"`
	Email          string    `json:"email"`
	Username       string    `json:"username"`
	TotalDonation  float64   `json:"totalDonation"`
	MobileNumber   string    `json:"mobileNumber"`
	Address        string    `json:"address"`
	Occupation     string    `json:"occupation"`
	Image          string    `json:"image"`
	Motto          string    `json:"motto"`
	IsAdmin        bool      `json:"isAdmin"`
	BankAccount    string    `json:"bankAccount"`
	Birthday       time.Time `json:"birthday"`
	Sex            string    `json:"sex"`
	Password       string    `json:"password,omitempty"`
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
			&u.MobileNumber,
			&u.Address,
			&u.Occupation,
			&u.Image,
			&u.Motto,
			&u.IsAdmin,
			&u.BankAccount,
			&u.Birthday,
			&u.Sex,
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
		&u.MobileNumber,
		&u.Address,
		&u.Occupation,
		&u.Image,
		&u.Motto,
		&u.IsAdmin,
		&u.BankAccount,
		&u.Birthday,
		&u.Sex,
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
    mobile_number string,
    address string,
    occupation string,
    image string,
    motto string,
    bank_account string,
    birthday time.Time,
    sex string,
) (int, error) {
	id := 0
	err := db.QueryRow(`
		select create_user($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)	
		`,email, password, username, mobile_number, address, occupation, 
			image, motto, bank_account, birthday, sex
	).Scan(&id)
	return id, err
}
