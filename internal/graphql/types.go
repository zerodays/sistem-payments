package graphql

import "github.com/graphql-go/graphql"

var paymentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Payments",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"phase_id": &graphql.Field{
				Type: graphql.Int,
			},
			"amount": &graphql.Field{
				Type: graphql.Int,
			},
			"payed": &graphql.Field{
				Type: graphql.Boolean,
			},
			"date_due": &graphql.Field{
				Type: graphql.DateTime,
			},
			"date_payed": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)

var userPaymentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "UserPayments",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"user_id": &graphql.Field{
				Type: graphql.String,
			},
			"amount": &graphql.Field{
				Type: graphql.Int,
			},
			"date_payed": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)
