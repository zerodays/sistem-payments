package router

import (
	"github.com/zerodays/sistem-payments/internal/handle"
	"net/http"
)

// apiRoutes returns routes. It is in function instead of a variable to allow
// lazy loading.
func apiRoutes() []Route {
	return []Route{
		{
			Name:           "payments",
			Path:           "/payments",
			AuthorizedOnly: true,
			GET:            http.HandlerFunc(handle.PaymentsHandle),
			POST:           http.HandlerFunc(handle.CreatePaymentHandle),
		},
		{
			Name:           "payments_overdue",
			Path:           "/payments/overdue",
			AuthorizedOnly: true,
			GET:            http.HandlerFunc(handle.OverdueHandle),
		},
		{
			Name:           "payment",
			Path:           "/payments/{id}",
			AuthorizedOnly: true,
			PUT:            http.HandlerFunc(handle.UpdatePaymentHandle),
			DELETE:         http.HandlerFunc(handle.DeletePaymentHandle),
		},

		{
			Name:           "user_payments",
			Path:           "/user_payments",
			AuthorizedOnly: true,
			GET:            http.HandlerFunc(handle.UserPaymentsHandle),
			POST:           http.HandlerFunc(handle.CreateUserPaymentHandle),
		},
		{
			Name:           "user_payment",
			Path:           "/user_payments/{id}",
			AuthorizedOnly: true,
			PUT:            http.HandlerFunc(handle.UpdateUserPaymentHandle),
			DELETE:         http.HandlerFunc(handle.DeleteUserPaymentHandle),
		},

		{
			Name:           "payments_per_projects",
			Path:           "/payments/projects",
			AuthorizedOnly: true,
			GET:            http.HandlerFunc(handle.PaymentsForProjectsHandle),
		},
		{
			Name:           "payments_per_project",
			Path:           "/payments/projects/{id}",
			AuthorizedOnly: true,
			GET:            http.HandlerFunc(handle.PaymentsForProjectHandle),
		},

		{
			Name: "fault_tolerance_switch",
			Path: "/fault",
			GET:  http.HandlerFunc(handle.FaultToleranceHandle),
		},
	}
}
