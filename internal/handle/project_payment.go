package handle

import (
	"encoding/json"
	"github.com/zerodays/sistem-payments/internal/handle/errors"
	"github.com/zerodays/sistem-payments/internal/logger"
	"github.com/zerodays/sistem-payments/internal/models/payment"
	"github.com/zerodays/sistem-payments/internal/models/project"
	"github.com/zerodays/sistem-payments/internal/util"
	"net/http"
)

type paymentsForProjectResponse struct {
	*project.Project
	Payments []*payment.Payment `json:"payments"`
}

// PaymentsForProjectsHandle returns payments per project
func PaymentsForProjectsHandle(w http.ResponseWriter, _ *http.Request) {
	projects, err := project.All()
	if err != nil {
		logger.Log.Warn().Err(err).Send()
		errors.Response(w, errors.UnknownError)
		return
	}

	resp := make([]paymentsForProjectResponse, len(projects))
	for i, pr := range projects {
		payments, err := payment.ForProjectID(pr.ID)
		if err != nil {
			logger.Log.Warn().Err(err).Send()
			errors.Response(w, errors.DatabaseError)
			return
		}

		resp[i] = paymentsForProjectResponse{
			Project:  pr,
			Payments: payments,
		}
	}

	res, _ := json.Marshal(resp)
	_, _ = w.Write(res)
}

// PaymentsForProjectHandle returns payments for specified project.
func PaymentsForProjectHandle(w http.ResponseWriter, r *http.Request) {
	id, ok := util.IdFromRequest(w, r)
	if !ok {
		return
	}

	payments, err := payment.ForProjectID(id)
	if err != nil {
		logger.Log.Warn().Err(err).Send()
		errors.Response(w, errors.DatabaseError)
		return
	}

	res, _ := json.Marshal(payments)
	_, _ = w.Write(res)
}
