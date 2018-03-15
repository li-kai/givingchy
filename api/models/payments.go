package models

import (
	"database/sql"
	"time"
)

// Payment represents one user's transaction
type Payment struct {
	ID        int       `json:"id"`
	UserID    int       `json:"userId"`
	ProjectID int       `json:"projectId"`
	PaidAt    time.Time `json:"paidAt"`
	Amount    float64   `json:"amount"`
}

func readPayments(rows *sql.Rows, err error) ([]*Payment, error) {
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	payments := []*Payment{}
	for rows.Next() {
		var payment Payment
		if err := rows.Scan(
			&payment.ID,
			&payment.UserID,
			&payment.ProjectID,
			&payment.PaidAt,
			&payment.Amount,
		); err != nil {
			return nil, err
		}
		payments = append(payments, &payment)
	}
	return payments, nil
}

// AllPayments gets all payments in db
func (db *DB) AllPayments() ([]*Payment, error) {
	return readPayments(db.Query(`
        SELECT id, user_id, project_id, moment, amount
        FROM payments`))
}

// AllProjectPayments gets all payments in db related to project
func (db *DB) AllProjectPayments(projectID int) ([]*Payment, error) {
	return readPayments(db.Query(`
        SELECT id, user_id, project_id, moment, amount
        FROM payments
        WHERE project_id = $1
    `, projectID))
}

// AllUserPayments gets all payments in db related to user
func (db *DB) AllUserPayments(userID int) ([]*Payment, error) {
	return readPayments(db.Query(`
        SELECT id, user_id, project_id, moment, amount
        FROM payments
        WHERE user_id = $1
    `, userID))
}

// CreatePayment creates a payment given user and project ids
func (db *DB) CreatePayment(userID int, projectID int, amount float64) (int, error) {
	id := 0
	err := db.QueryRow(`
        INSERT INTO payments (user_id, project_id, amount)
        VALUES($1, $2, $3)
        RETURNING id
    `, userID, projectID, amount,
	).Scan(&id)
	return id, err
}

// UpdatePayment updates a payment given id and new updated text
func (db *DB) UpdatePayment(paymentID int, amount float64) error {
	_, err := db.Exec(`
        UPDATE payments
        SET amount = $1
        WHERE id = $2
    `, amount, paymentID,
	)
	return err
}

// DeletePayment deletes a payment given id
func (db *DB) DeletePayment(paymentID int) error {
	_, err := db.Exec(`
        DELETE FROM payments
        WHERE id = $1
    `, paymentID,
	)
	return err
}
