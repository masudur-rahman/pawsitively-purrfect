package handlers

import (
	"net/http"

	"github.com/masudur-rahman/pawsitively-purrfect/api/graphql/resolvers"
	"github.com/masudur-rahman/pawsitively-purrfect/api/graphql/schema"
	"github.com/masudur-rahman/pawsitively-purrfect/api/http/middlewares"
	"github.com/masudur-rahman/pawsitively-purrfect/infra/logr"
	"github.com/masudur-rahman/pawsitively-purrfect/models"
	"github.com/masudur-rahman/pawsitively-purrfect/services/all"

	"github.com/flamego/session"
	"github.com/graphql-go/graphql"
)

type RequestOptions struct {
	Query         string                 `json:"query" url:"query" schema:"query"`
	Variables     map[string]interface{} `json:"variables" url:"variables" schema:"variables"`
	OperationName string                 `json:"operationName" url:"operationName" schema:"operationName"`
}

func ServeGraphQL(ctx *middlewares.PurrfectContext, sess session.Session, opts RequestOptions, svc *all.Services) {
	resolver := resolvers.NewResolver(ctx, svc)
	schemas, err := schema.PurrfectSchema(resolver)
	if err != nil {
		ServeJson(ctx.ResponseWriter(), http.StatusInternalServerError, gqlError(err))
		logr.DefaultLogger.Errorw("Cannot fetch purrfect schema", "error", err)
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
		var status int
		var errMsg string
		for idx, e := range result.Errors {
			status, errMsg = models.ParseStatusError(e.OriginalError())
			result.Errors[idx].Message = errMsg
			logr.DefaultLogger.Errorw("Serve GraphQL", "error", e.OriginalError())
		}

		ServeJson(ctx.ResponseWriter(), status, result)
		return
	}

	if opts.IsLoginMutation() {
		HandlePostLoginWithGQLResult(ctx, sess, result)
		return
	}

	ServeJson(ctx.ResponseWriter(), http.StatusOK, result)
	return
}
