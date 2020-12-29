package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/zerodays/sistem-payments/internal/models/payment"
	"time"
)

var createPayment = &graphql.Field{
	Type:        paymentType,
	Description: "Create new payment",
	Args: graphql.FieldConfigArgument{
		"phase_id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"amount": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"date_due": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.DateTime),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		pym := &payment.Payment{
			PhaseID: p.Args["phase_id"].(int),
			Amount:  p.Args["amount"].(int),
			DateDue: p.Args["date_due"].(time.Time),
		}

		err := pym.Insert()
		if err == nil {
			return pym, nil
		} else {
			return nil, err
		}
	},
}

var updatePayment = &graphql.Field{
	Type:        paymentType,
	Description: "Update payment",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"amount": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"payed": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
		"date_due": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.DateTime),
		},
		"date_payed": &graphql.ArgumentConfig{
			Type: graphql.DateTime,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		pym, err := payment.ForID(p.Args["id"].(int))
		if err != nil {
			return nil, err
		}

		pym.Amount = p.Args["amount"].(int)
		pym.Payed = p.Args["payed"].(bool)
		pym.DateDue = p.Args["date_due"].(time.Time)

		if !pym.Payed {
			pym.DatePayed.Valid = false
		} else {
			datePayed, ok := p.Args["date_payed"].(time.Time)
			if !ok {
				datePayed = time.Now()
			}

			pym.DatePayed.Time = datePayed
			pym.DatePayed.Valid = true
		}

		err = pym.Update()
		if err == nil {
			return pym, nil
		} else {
			return nil, err
		}
	},
}

var deletePayment = &graphql.Field{
	Type:        paymentType,
	Description: "Deletes payment",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		pym, err := payment.ForID(p.Args["id"].(int))
		if err != nil {
			return nil, err
		}

		err = pym.Delete()
		if err == nil {
			return pym, nil
		} else {
			return nil, err
		}
	},
}

var mutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"create_payment": createPayment,
			"update_payment": updatePayment,
			"delete_payment": deletePayment,
		},
	},
)
