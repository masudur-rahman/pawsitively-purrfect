package arangodb

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"

	arango "github.com/arangodb/go-driver"
)

type Query struct {
	queryString string
	bindVars    map[string]interface{}
}

func generateArangoQuery(collection string, filter interface{}, removeQuery bool) *Query {
	queryString := "FOR doc IN " + collection + " FILTER "
	bindVars := map[string]interface{}{}

	var filters []string

	val := reflect.ValueOf(filter)
	for idx := 0; idx < val.NumField(); idx++ {
		field := val.Type().Field(idx)
		if val.Field(idx).IsZero() {
			continue
		}

		fieldName := strings.ToLower(field.Name)
		filters = append(filters, fmt.Sprintf("doc.%s == @%s", fieldName, fieldName))
		bindVars[fieldName] = val.Field(idx).Interface()
	}

	if len(filters) > 0 {
		queryString += " FILTER "
		queryString += strings.Join(filters, " AND ")
	}

	if removeQuery {
		queryString += " REMOVE doc IN " + collection
	} else {
		queryString += " RETURN doc"
	}

	return &Query{
		queryString: queryString,
		bindVars:    bindVars,
	}
}

func executeArangoQuery(ctx context.Context, db arango.Database, query *Query, lim int64) ([]interface{}, error) {
	cursor, err := db.Query(ctx, query.queryString, query.bindVars)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := cursor.Close(); err != nil {
			log.Printf("Error closing cursor: %v", err)
		}
	}()

	var results []interface{}
	for {
		var doc interface{}
		_, err = cursor.ReadDocument(ctx, &doc)
		if arango.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return nil, err
		}

		results = append(results, doc)
		if lim > 0 && int64(len(results)) >= lim {
			break
		}
	}

	return results, nil
}
