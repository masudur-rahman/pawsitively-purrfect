package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/masudur-rahman/pawsitively-purrfect/infra/logr"

	"github.com/flamego/flamego"
	"github.com/graphql-go/graphql"
)

type RequestOptions struct {
	Query         string                 `json:"query" url:"query" schema:"query"`
	Variables     map[string]interface{} `json:"variables" url:"variables" schema:"variables"`
	OperationName string                 `json:"operationName" url:"operationName" schema:"operationName"`
}

func ServeGraphQL(ctx flamego.Context, opts RequestOptions, schemas graphql.Schema) {
	result := graphql.Do(graphql.Params{
		Schema:         schemas,
		RequestString:  opts.Query,
		VariableValues: opts.Variables,
		OperationName:  opts.OperationName,
		Context:        ctx.Request().Context(),
	})

	if result.HasErrors() {
		for _, e := range result.Errors {
			if e.OriginalError() != nil {
				logr.DefaultLogger.Errorf(e.OriginalError().Error())
			}
		}

		ServeJson(ctx.ResponseWriter(), http.StatusBadRequest, result)
		return
	}

	ServeJson(ctx.ResponseWriter(), http.StatusOK, result)
	return
}

func ServeJson(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println(err.Error())
	}
}
