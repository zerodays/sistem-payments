package handle

import (
	"encoding/json"
	"github.com/zerodays/sistem-payments/internal/handle/errors"
	"github.com/zerodays/sistem-payments/internal/logger"
	"github.com/zerodays/sistem-payments/internal/models/payment"
	"github.com/zerodays/sistem-payments/internal/util"
	"gopkg.in/guregu/null.v4"
	"net/http"
	"strconv"
	"time"
)

type newPaymentRequest struct {
	PhaseID int       `json:"phase_id"`
	Amount  int       `json:"amount"`
	DateDue time.Time `json:"date_due"`
}

type updatePaymentRequest struct {
	Amount int  `json:"amount"`
	Payed  bool `json:"payed"`

	DateDue   time.Time `json:"date_due"`
	DatePayed null.Time `json:"date_payed"`
}

func listResponse(w http.ResponseWriter, payments []*payment.Payment, err error) {
	if err != nil {
		logger.Log.Warn().Err(err).Send()
		errors.Response(w, errors.DatabaseError)
		return
	}

	res, _ := json.Marshal(payments)
	_, _ = w.Write(res)
}

// PaymentsHandle lists all payments.
// If GET parameter `payed` is supplied, only payments
// that have payed status same as `payed` are returned.
func PaymentsHandle(w http.ResponseWriter, r *http.Request) {
	var (
		payments []*payment.Payment
		err      error
	)

	payedStr := r.URL.Query().Get("payed")

	if payedStr == "" {
		payments, err = payment.All()
	} else {
		payed, err := strconv.ParseBool(payedStr)
		if err != nil {
			payments, err = payment.All()
		} else {
			payments, err = payment.Payed(payed)
		}
	}

	listResponse(w, payments, err)
}

// OverdueHandle lists all overdue payments.
func OverdueHandle(w http.ResponseWriter, _ *http.Request) {
	payments, err := payment.Overdue()
	listResponse(w, payments, err)
}

// CreatePaymentHandle creates new payment.
func CreatePaymentHandle(w http.ResponseWriter, r *http.Request) {
	req := &newPaymentRequest{}
	if !util.ParseJSON(w, r, req) {
		return
	}

	p := &payment.Payment{
		PhaseID: req.PhaseID,
		Amount:  req.Amount,
		DateDue: req.DateDue,
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

// UpdatePaymentHandle updates payment.
func UpdatePaymentHandle(w http.ResponseWriter, r *http.Request) {
	req := &updatePaymentRequest{}
	if !util.ParseJSON(w, r, req) {
		return
	}

	if req.Payed && !req.DatePayed.Valid {
		req.DatePayed.SetValid(time.Now())
	}

	id, ok := util.IdFromRequest(w, r)
	if !ok {
		return
	}

	p, err := payment.ForID(id)
	if err != nil {
		logger.Log.Warn().Err(err).Send()
		errors.Response(w, errors.DatabaseError)
		return
	}

	p.Amount = req.Amount
	p.Payed = req.Payed
	p.DateDue = req.DateDue
	p.DatePayed = req.DatePayed
	err = p.Update()
	if err != nil {
		logger.Log.Warn().Err(err).Send()
		errors.Response(w, errors.DatabaseError)
		return
	}

	res, _ := json.Marshal(p)
	_, _ = w.Write(res)
}

// DeletePaymentHandle deletes payment for specified ID.
func DeletePaymentHandle(w http.ResponseWriter, r *http.Request) {
	id, ok := util.IdFromRequest(w, r)
	if !ok {
		return
	}

	p, err := payment.ForID(id)
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
