package sql

import "database/sql"

type Database interface {
	ID(id string) Database

	SetFilter(filter string, args ...interface{}) Database

	FindOne(document interface{}) (bool, error)
	FindMany(documents interface{}) error

	InsertOne(document interface{}) (id string, err error)
	InsertMany(documents []interface{}) ([]string, error)

	UpdateOne(document interface{}) error

	DeleteOne() error

	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}
