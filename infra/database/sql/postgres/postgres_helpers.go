package postgres

import (
	"net/http"

	"github.com/masudur-rahman/pawsitively-purrfect/infra/database/sql/postgres/pb"
	"github.com/masudur-rahman/pawsitively-purrfect/models"
)

func InitializePostgresClient() (pb.PostgresClient, error) {
	return nil, nil
}

func checkIDNonEmpty(id string) error {
	if id == "" {
		return models.StatusError{
			Status:  http.StatusBadRequest,
			Message: "must provide document id",
		}
	}
	return nil
}

func checkIdOrFilterNonEmpty(id string, filter interface{}) error {
	if id == "" && filter == nil {
		return models.StatusError{
			Status:  http.StatusBadRequest,
			Message: "must provide id and/or filter",
		}
	}
	return nil
}
