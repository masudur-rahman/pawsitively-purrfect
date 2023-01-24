package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/masudur-rahman/pawsitively-purrfect/api/graphql/resolvers"
	"github.com/masudur-rahman/pawsitively-purrfect/api/graphql/schema"

	"github.com/flamego/flamego"
	"github.com/graphql-go/graphql"
)

type RequestOptions struct {
	Query         string                 `json:"query" url:"query" schema:"query"`
	Variables     map[string]interface{} `json:"variables" url:"variables" schema:"variables"`
	OperationName string                 `json:"operationName" url:"operationName" schema:"operationName"`
}

func ServeGraphQL(ctx flamego.Context, opts RequestOptions, resolver *resolvers.Resolver) {
	schemas, err := schema.PurrfectSchema(resolver)
	if err != nil {
		ctx.ResponseWriter().WriteHeader(http.StatusInternalServerError)
		ctx.ResponseWriter().Write([]byte(err.Error()))
		return
	}

	result := graphql.Do(graphql.Params{
		Schema:         schemas,
		RequestString:  opts.Query,
		VariableValues: opts.Variables,
		OperationName:  opts.OperationName,
		Context:        ctx.Request().Context(),
	})

	if result.HasErrors() {
		var errors []string
		for _, e := range result.Errors {
			errors = append(errors, e.Error())
		}
		ctx.ResponseWriter().WriteHeader(http.StatusBadRequest)
		ctx.ResponseWriter().Write([]byte(result.Errors[0].Error()))
		return
	}

	ctx.ResponseWriter().Header().Set("Content-Type", "application/json")
	ctx.ResponseWriter().WriteHeader(http.StatusOK)
	json.NewEncoder(ctx.ResponseWriter()).Encode(result.Data)
	return
}
