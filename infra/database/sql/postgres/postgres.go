package postgres

import (
	"context"
	"database/sql"

	isql "github.com/masudur-rahman/pawsitively-purrfect/infra/database/sql"
	"github.com/masudur-rahman/pawsitively-purrfect/infra/database/sql/postgres/pb"
)

type Database struct {
	ctx    context.Context
	table  string
	id     string
	client pb.PostgresClient
}

func NewDatabase(client pb.PostgresClient) Database {
	return Database{client: client}
}

func (d Database) Table(name string) isql.Database {
	d.table = name
	return d
}

func (d Database) ID(id string) isql.Database {
	d.id = id
	return d
}

func (d Database) SetFilter(filter string, args ...interface{}) isql.Database {
	//TODO implement me
	panic("implement me")
}

func (d Database) FindOne(document interface{}) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (d Database) FindMany(documents interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (d Database) InsertOne(document interface{}) (id string, err error) {
	//TODO implement me
	panic("implement me")
}

func (d Database) InsertMany(documents []interface{}) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (d Database) UpdateOne(document interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (d Database) DeleteOne() error {
	//TODO implement me
	panic("implement me")
}

func (d Database) Query(query string, args ...interface{}) (*sql.Rows, error) {
	//TODO implement me
	panic("implement me")
}

func (d Database) Exec(query string, args ...interface{}) (sql.Result, error) {
	//TODO implement me
	panic("implement me")
}
