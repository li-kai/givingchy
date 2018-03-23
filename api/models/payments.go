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
		select * from all_payments()
		`))
}

// AllProjectPayments gets all payments in db related to project
func (db *DB) AllProjectPayments(projectID int) ([]*Payment, error) {
	return readPayments(db.Query(`
        select * from all_project_payments($1)
    `, projectID))
}

// AllUserPayments gets all payments in db related to user
func (db *DB) AllUserPayments(userID int) ([]*Payment, error) {
	return readPayments(db.Query(`
        select * from all_user_payments($1)
    `, userID))
}

// CreatePayment creates a payment given user and project ids
func (db *DB) CreatePayment(userID int, projectID int, amount float64) (int, error) {
	id := 0
	err := db.QueryRow(`
        select create_payment($1, $2, $3)
    `, userID, projectID, amount,
	).Scan(&id)
	return id, err
}

// UpdatePayment updates a payment given id and new updated text
func (db *DB) UpdatePayment(paymentID int, amount float64) error {
	_, err := db.Exec(`
        select update_payment($1, $2)
    `, amount, paymentID,
	)
	return err
}

// DeletePayment deletes a payment given id
func (db *DB) DeletePayment(paymentID int) error {
	_, err := db.Exec(`
        select delete_payment($1)
    `, paymentID,
	)
	return err
}
