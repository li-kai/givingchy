package models

import "github.com/lib/pq"

// Statistic represents one Statistic for projects
type Statistic map[string]interface{}

type donor struct {
	Username      string  `json:"username"`
	TotalDonation float64 `json:"totalDonation"`
	Image         string  `json:"image"`
}

// AllStatistics gets all statistics in db such as min, max, avg, count
func (db *DB) AllStatistics() ([]*Statistic, error) {
	stats := []*Statistic{}

	var arr []float64
	if err := db.QueryRow(`
        select * from get_users_quarter_donations()
    `).Scan(pq.Array(&arr)); err != nil {
		return nil, err
	}
	userState := Statistic{
		"name":   "users",
		"values": arr,
	}
	stats = append(stats, &userState)

	rows, err := db.Query(`
		select * from top_donors()
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*donor{}
	for rows.Next() {
		var u donor
		if err := rows.Scan(
			&u.Username,
			&u.TotalDonation,
			&u.Image,
		); err != nil {
			return nil, err
		}
		users = append(users, &u)
	}
	donors := Statistic{
		"name":   "top_donors",
		"values": users,
	}
	stats = append(stats, &donors)

	return stats, nil
}
