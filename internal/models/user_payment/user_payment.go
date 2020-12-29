package user_payment

import (
	"github.com/zerodays/sistem-payments/internal/database"
	"time"
)

type UserPayment struct {
	ID     int    `db:"id" json:"id"`
	UserID string `db:"user_id" json:"user_id"`

	Amount    int       `db:"amount" json:"amount"`
	DatePayed time.Time `db:"date_payed" json:"date_payed"`
}

// Insert inserts user payment into database.
func (p *UserPayment) Insert() error {
	insert := `INSERT INTO user_payment (user_id, amount, date_payed) VALUES ($1, $2, $3) RETURNING *`
	return database.DB.Get(p, insert, p.UserID, p.Amount, p.DatePayed)
}

// Update updates user payment.
func (p *UserPayment) Update() error {
	update := `UPDATE user_payment SET amount=$2, date_payed=$3 WHERE id=$1 RETURNING *`
	return database.DB.Get(p, update, p.ID, p.Amount, p.DatePayed)
}

// Delete deletes user payment.
func (p *UserPayment) Delete() error {
	del := `DELETE FROM user_payment WHERE id=$1`
	_, err := database.DB.Exec(del, p.ID)
	return err
}

// ForID gets user payment for specified ID.
func ForID(id int) (*UserPayment, error) {
	p := &UserPayment{}
	query := `SELECT * FROM user_payment WHERE id=$1`
	err := database.DB.Get(p, query, id)

	return p, err
}

// All returns all payments.
func All() ([]*UserPayment, error) {
	payments := make([]*UserPayment, 0)
	query := `SELECT * FROM user_payment ORDER BY date_payed DESC`

	err := database.DB.Select(&payments, query)
	return payments, err
}

// Total returns sum of all user payments
func Total() (int, error) {
	query := `SELECT SUM(amount) FROM user_payment`
	var amount int

	err := database.DB.Get(&amount, query)
	return amount, err
}
