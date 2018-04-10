package models

// Tag represents a tag for projects
type Tag struct {
	TagName   string `json:"tagName"`
	ProjectID int    `json:"projectID"`
}

// AllTags returns all tags from all projects
func (db *DB) AllTags() ([]*Tag, error) {
	rows, err := db.Query(`select * from all_tags()`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tags := []*Tag{}
	for rows.Next() {
		var tag Tag
		if err := rows.Scan(
			&tag.TagName,
		); err != nil {
			return nil, err
		}
		tags = append(tags, &tag)
	}

	return tags, nil
}

// CreateTag creates a user defined tag
func (db *DB) CreateTag(projectID int, tagName string) error {
	_, err := db.Exec(`
		select create_tag($1, $2)
	`, projectID, tagName)
	return err
}

// AllProjectTags returns all tags that belonging to the project
func (db *DB) AllProjectTags(projectID int) ([]*Tag, error) {
	rows, err := db.Query(`select * from get_project_s_tags($1)`, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tags := []*Tag{}
	for rows.Next() {
		var tag Tag
		if err := rows.Scan(
			&tag.TagName,
		); err != nil {
			return nil, err
		}
		tags = append(tags, &tag)
	}

	return tags, nil
}
