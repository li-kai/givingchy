package models

import (
	"database/sql"
	"time"
)

// Comment represents one comment
type Comment struct {
	ID        int       `json:"id"`
	UserID    int       `json:"userId"`
	ProjectID int       `json:"projectId"`
	CreatedAt time.Time `json:"createdAt"`
	Content   string    `json:"content"`
}

func readComments(rows *sql.Rows, err error) ([]*Comment, error) {
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	comments := []*Comment{}
	for rows.Next() {
		var comment Comment
		if err := rows.Scan(
			&comment.ID,
			&comment.UserID,
			&comment.ProjectID,
			&comment.CreatedAt,
			&comment.Content,
		); err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}
	return comments, nil
}

// AllComments gets all comments in db
func (db *DB) AllComments(numPerPage int, pageIdx int) ([]*Comment, error) {
	return readComments(db.Query(`
		select * from all_comments($1, $2)
		`, numPerPage, pageIdx))
}

// AllProjectComments gets all comments in db related to project
func (db *DB) AllProjectComments(projectID int, numPerPage int, pageIdx int) ([]*Comment, error) {
	return readComments(db.Query(`
        select * from all_project_comments($1, $2, $3)
    `, projectID, numPerPage, pageIdx))
}

// AllUserComments gets all comments in db related to user
func (db *DB) AllUserComments(userID int, numPerPage int, pageIdx int) ([]*Comment, error) {
	return readComments(db.Query(`
        select * from all_user_comments($1, $2, $3)
    `, userID, numPerPage, pageIdx))
}

// CreateComment creates a comment given user and project ids
func (db *DB) CreateComment(userID int, projectID int, content string) (int, error) {
	id := 0
	err := db.QueryRow(`
        select create_comment($1, $2, $3)
    `, userID, projectID, content,
	).Scan(&id)
	return id, err
}

// UpdateComment updates a comment given id and new updated text
func (db *DB) UpdateComment(commentID int, content string) error {
	_, err := db.Exec(`
        select update_comment($1, $2)
    `, content, commentID,
	)
	return err
}

// DeleteComment deletes a comment given id
func (db *DB) DeleteComment(commentID int) error {
	_, err := db.Exec(`
        select delete_commet($1)
    `, commentID,
	)
	return err
}
