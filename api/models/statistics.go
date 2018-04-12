package models

// Statistic represents one Statistic for projects
type Statistic struct {
	Name           string  `json:"name"`
	Min            float64 `json:"min"`
	Max            float64 `json:"max"`
	Avg            float64 `json:"avg"`
	FirstQuartile  float64 `json:"firstQuartile"`
	SecondQuartile float64 `json:"secondQuartile"`
	ThirdQuartile  float64 `json:"thirdQuartile"`
}

// AllStatistics gets all statistics in db such as min, max, avg, count
// func (db *DB) AllStatistics() ([]*Statistic, error) {
// 	rows, err := db.Query(`select * from all_categories()`)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	categories := []*Statistic{}
// 	for rows.Next() {
// 		var Statistic Statistic
// 		if err := rows.Scan(
// 			&Statistic.Name,
// 			&Statistic.ProjNum,
// 		); err != nil {
// 			return nil, err
// 		}
// 		categories = append(categories, &Statistic)
// 	}

// 	return categories, nil
// }
