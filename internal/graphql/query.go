package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/zerodays/sistem-payments/internal/models/payment"
	"github.com/zerodays/sistem-payments/internal/models/user_payment"
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"payments": &graphql.Field{
				Type:        graphql.NewList(paymentType),
				Description: "Get payments list",
				Args: graphql.FieldConfigArgument{
					"payed": &graphql.ArgumentConfig{
						Type: graphql.Boolean,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					payed, ok := p.Args["payed"].(bool)
					if ok {
						return payment.Payed(payed)
					} else {
						return payment.All()
					}
				},
			},
			"payment": &graphql.Field{
				Type:        paymentType,
				Description: "Get payment for ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return payment.ForID(p.Args["id"].(int))
				},
			},
			"user_payments": &graphql.Field{
				Type:        graphql.NewList(userPaymentType),
				Description: "Get user payments list",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return user_payment.All()
				},
			},
			"user_payment": &graphql.Field{
				Type:        userPaymentType,
				Description: " Get user payment for ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return user_payment.ForID(p.Args["id"].(int))
				},
			},
		},
	},
)
