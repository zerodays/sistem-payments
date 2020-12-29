package graphql

import (
	"encoding/json"
	"github.com/graphql-go/graphql"
	"github.com/zerodays/sistem-payments/internal/handle/errors"
	"github.com/zerodays/sistem-payments/internal/logger"
	"io/ioutil"
	"net/http"
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	},
)

func Handle(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Log.Warn().Err(err).Send()
		errors.Response(w, errors.UnknownError)
		return
	}

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: string(body),
	})

	if len(result.Errors) > 0 {
		for _, err := range result.Errors {
			logger.Log.Warn().Err(err).Send()
		}
	}

	_ = json.NewEncoder(w).Encode(result)
}
