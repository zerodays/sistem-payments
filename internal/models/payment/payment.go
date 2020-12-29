package payment

import (
	"github.com/zerodays/sistem-payments/internal/database"
	"gopkg.in/guregu/null.v4"
	"time"
)

type Payment struct {
	ID      int `db:"id" json:"id"`
	PhaseID int `db:"phase_id" json:"phase_id"`

	Amount int  `db:"amount" json:"amount"`
	Payed  bool `db:"payed" json:"payed"`

	DateDue   time.Time `db:"date_due" json:"date_due"`
	DatePayed null.Time `db:"date_payed" json:"date_payed"`
}

// Insert inserts payment p into database.
func (p *Payment) Insert() error {
	insert := `INSERT INTO payments (phase_id, amount, date_due) VALUES ($1, $2, $3) RETURNING *`
	return database.DB.Get(p, insert, p.PhaseID, p.Amount, p.DateDue)
}

// Update updates payment.
func (p *Payment) Update() error {
	update := `UPDATE payments SET amount=$2, payed=$3, date_due=$4, date_payed=$5 WHERE id=$1 RETURNING *`
	return database.DB.Get(p, update, p.ID, p.Amount, p.Payed, p.DatePayed, p.DatePayed)
}

// Delete deletes payment.
func (p *Payment) Delete() error {
	del := `DELETE FROM payments WHERE id=$1`
	_, err := database.DB.Exec(del, p.ID)
	return err
}

// ForID returns payment for specified ID.
func ForID(id int) (*Payment, error) {
	p := &Payment{}
	query := `SELECT * FROM payments WHERE id=$1`
	err := database.DB.Get(p, query, id)
	return p, err
}

// All returns all payments in database
func All() ([]*Payment, error) {
	payments := make([]*Payment, 0)
	query := `SELECT * FROM payments ORDER BY date_due DESC`

	err := database.DB.Select(&payments, query)
	return payments, err
}

// Payed returns payments that have payed parameter set to `payed`.
func Payed(payed bool) ([]*Payment, error) {
	payments := make([]*Payment, 0)
	query := `SELECT * FROM payments WHERE payed=$1 ORDER BY date_due DESC`

	err := database.DB.Select(&payments, query, payed)
	return payments, err
}

// Overdue returns payments that are overdue.
func Overdue() ([]*Payment, error) {
	payments := make([]*Payment, 0)
	query := `SELECT * FROM payments WHERE payed=FALSE AND date_due <= $1 ORDER BY date_due`

	err := database.DB.Select(&payments, query, time.Now())
	return payments, err
}

// Total returns sum of all payments in database.
func Total() (int, error) {
	query := `SELECT SUM(amount) FROM payments`
	var amount int

	err := database.DB.Get(&amount, query)
	return amount, err
}

// TotalPayed returns sum of all payed payments in database.
func TotalPayed() (int, error) {
	query := `SELECT SUM(amount) FROM payments WHERE payed=TRUE`
	var amount int

	err := database.DB.Get(&amount, query)
	return amount, err
}

// TotalOverdue return sum of all overdue payments in database.
func TotalOverdue() (int, error) {
	query := `SELECT SUM(amount) FROM payments WHERE payed=FALSE AND date_due <= $1`
	var amount int

	err := database.DB.Get(&amount, query, time.Now())
	return amount, err
}
