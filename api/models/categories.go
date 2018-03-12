package models

// Category represents one category for projects
type Category struct {
	Name string `json:"name"`
}

// AllCategories gets all categories in db
func (db *DB) AllCategories() ([]*Category, error) {
	rows, err := db.Query(`SELECT name FROM categories`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := []*Category{}
	for rows.Next() {
		var category Category
		if err := rows.Scan(
			&category.Name,
		); err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}

	return categories, nil
}

// CreateCategory creates a category
// requires admin access
func (db *DB) CreateCategory(name string) error {
	_, err := db.Exec(`
        INSERT INTO categories (name)
        VALUES($1)
    `, name)
	return err
}
