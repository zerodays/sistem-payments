package handle

import (
	"encoding/json"
	"github.com/zerodays/sistem-payments/internal/handle/errors"
	"github.com/zerodays/sistem-payments/internal/logger"
	"github.com/zerodays/sistem-payments/internal/models/user_payment"
	"github.com/zerodays/sistem-payments/internal/util"
	"net/http"
	"time"
)

type newUserPaymentRequest struct {
	UserID    string    `json:"user_id"`
	Amount    int       `json:"amount"`
	DatePayed time.Time `json:"date_payed"`
}

type updateUserPaymentRequest struct {
	Amount    int       `json:"amount"`
	DatePayed time.Time `json:"date_payed"`
}

// UserPaymentsHandle returns all user payments.
func UserPaymentsHandle(w http.ResponseWriter, _ *http.Request) {
	payments, err := user_payment.All()
	if err != nil {
		logger.Log.Warn().Err(err).Send()
		errors.Response(w, errors.DatabaseError)
		return
	}

	res, _ := json.Marshal(payments)
	_, _ = w.Write(res)
}

// CreateUserPaymentHandle creates new user payment.
func CreateUserPaymentHandle(w http.ResponseWriter, r *http.Request) {
	req := &newUserPaymentRequest{}
	if !util.ParseJSON(w, r, req) {
		return
	}

	p := &user_payment.UserPayment{
		UserID:    req.UserID,
		Amount:    req.Amount,
		DatePayed: req.DatePayed,
	}

	err := p.Insert()
	if err != nil {
		logger.Log.Warn().Err(err).Send()
		errors.Response(w, errors.DatabaseError)
		return
	}

	res, _ := json.Marshal(p)
	_, _ = w.Write(res)
}

// UpdateUserPaymentHandle updates user payment.
func UpdateUserPaymentHandle(w http.ResponseWriter, r *http.Request) {
	req := &updateUserPaymentRequest{}
	if !util.ParseJSON(w, r, req) {
		return
	}

	id, ok := util.IdFromRequest(w, r)
	if !ok {
		return
	}

	p, err := user_payment.ForID(id)
	if err != nil {
		logger.Log.Warn().Err(err).Send()
		errors.Response(w, errors.DatabaseError)
		return
	}

	p.DatePayed = req.DatePayed
	p.Amount = req.Amount
	err = p.Update()
	if err != nil {
		logger.Log.Warn().Err(err).Send()
		errors.Response(w, errors.DatabaseError)
		return
	}

	res, _ := json.Marshal(p)
	_, _ = w.Write(res)
}

// DeleteUserPaymentHandle deletes user payment for specified id.
func DeleteUserPaymentHandle(w http.ResponseWriter, r *http.Request) {
	id, ok := util.IdFromRequest(w, r)
	if !ok {
		return
	}

	p, err := user_payment.ForID(id)
	if err != nil {
		logger.Log.Warn().Err(err).Send()
		errors.Response(w, errors.DatabaseError)
		return
	}

	err = p.Delete()
	if err != nil {
		logger.Log.Warn().Err(err).Send()
		errors.Response(w, errors.DatabaseError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
