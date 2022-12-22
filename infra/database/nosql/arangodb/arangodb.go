package arangodb

import (
	"context"
	"github.com/masudur-rahman/pawsitively-purrfect/infra/database/nosql"
)
import (
	arango "github.com/arangodb/go-driver"
)

type ArangoDB struct {
	db         arango.Database
	ctx        context.Context
	id         string
	collection string
}

func NewArangoDB(db arango.Database, ctx context.Context) *ArangoDB {
	return &ArangoDB{
		db:  db,
		ctx: ctx,
	}
}

func (a *ArangoDB) ID(id string) nosql.Database {
	a.id = id
	return a
}

func (a *ArangoDB) Collection(collection string) nosql.Database {
	a.collection = collection
	return a
}

func (a *ArangoDB) FindOne(document interface{}, filter interface{}) (bool, error) {
	return true, nil
}

func (a *ArangoDB) FindMany(documents interface{}, filter interface{}) error {

	return nil
}

func (a *ArangoDB) InsertOne(document interface{}) (id string, err error) {
	return "", nil
}

func (a *ArangoDB) InsertMany(documents []interface{}) ([]string, error) {
	return nil, nil
}

func (a *ArangoDB) UpdateOne(document interface{}) error {
	//coll := a.db.Collection(a.collection)

	return nil
}

func (a *ArangoDB) DeleteOne() error {
	//TODO implement me
	panic("implement me")
}

func (a *ArangoDB) Query(query string, bindParams map[string]interface{}) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}
