package arangodb

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"

	"github.com/masudur-rahman/pawsitively-purrfect/configs"
	"github.com/masudur-rahman/pawsitively-purrfect/infra/database/nosql"
	"github.com/masudur-rahman/pawsitively-purrfect/pkg"

	"github.com/masudur-rahman/go-oneliners"

	arango "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

type ArangoDB struct {
	ctx            context.Context
	db             arango.Database
	id             string
	collectionName string
}

func NewArangoDB(ctx context.Context, db arango.Database) ArangoDB {
	return ArangoDB{
		db:  db,
		ctx: ctx,
	}
}

func InitializeArangoDB(ctx context.Context) (arango.Database, error) {
	cfg := configs.PurrfectConfig.Database.ArangoDB
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{fmt.Sprintf("http://%s:%s", cfg.Host, cfg.Port)},
		TLSConfig: &tls.Config{ /*...*/ },
	})
	if err != nil {
		return nil, err
	}

	c, err := arango.NewClient(arango.ClientConfig{
		Connection:     conn,
		Authentication: arango.BasicAuthentication(cfg.User, cfg.Password),
	})
	if err != nil {
		return nil, err
	}

	db, err := c.Database(ctx, cfg.Name)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (a ArangoDB) Collection(collection string) nosql.Database {
	a.collectionName = collection
	return a
}

func (a ArangoDB) ID(id string) nosql.Database {
	a.id = id
	return a
}

func (a ArangoDB) FindOne(document interface{}, filter ...interface{}) (bool, error) {
	if a.id == "" && filter == nil {
		return false, errors.New("must provide id and/or filter")
	}

	collection, err := getDBCollection(a.ctx, a.db, a.collectionName)
	if err != nil {
		return false, err
	}

	if filter == nil {
		meta, err := collection.ReadDocument(a.ctx, a.id, document)
		return meta.ID != "", err
	}

	query := generateArangoQuery(a.collectionName, filter[0], false)
	fmt.Println("Find One => ", query.queryString, query.bindVars)
	results, err := executeArangoQuery(a.ctx, a.db, query, 1)
	if err != nil {
		return false, err
	}

	oneliners.PrettyJson(results, "Data")
	if len(results) != 1 {
		return false, nil
	}

	//reflect.ValueOf(documents).Elem().Set(reflect.ValueOf(results))
	if err = pkg.ParseInto(results[0], document); err != nil {
		return false, err
	}
	return true, nil
}

func (a ArangoDB) FindMany(documents interface{}, filter interface{}) error {
	_, err := getDBCollection(a.ctx, a.db, a.collectionName)
	if err != nil {
		return err
	}

	query := generateArangoQuery(a.collectionName, filter, false)
	results, err := executeArangoQuery(a.ctx, a.db, query, -1)
	if err != nil {
		return err
	}

	return pkg.ParseInto(results[0], documents)
}

func (a ArangoDB) InsertOne(document interface{}) (id string, err error) {
	collection, err := getDBCollection(a.ctx, a.db, a.collectionName)
	if err != nil {
		return "", err
	}

	meta, err := collection.CreateDocument(a.ctx, document)
	if err != nil {
		return "", err
	}

	return meta.Key, nil
}

func (a ArangoDB) InsertMany(documents []interface{}) ([]string, error) {
	collection, err := getDBCollection(a.ctx, a.db, a.collectionName)
	if err != nil {
		return nil, err
	}

	metas, _, err := collection.CreateDocuments(a.ctx, documents)
	if err != nil {
		return nil, err
	}

	// Extract IDs of inserted documents
	ids := make([]string, len(metas))
	for i, result := range metas {
		ids[i] = string(result.ID)
	}

	return ids, nil
}

func (a ArangoDB) UpdateOne(document interface{}) error {
	if a.id == "" {
		return errors.New("id must be provided")
	}

	collection, err := getDBCollection(a.ctx, a.db, a.collectionName)
	if err != nil {
		return err
	}

	_, err = collection.UpdateDocument(a.ctx, a.id, document)
	return err
}

func (a ArangoDB) DeleteOne(filter ...interface{}) error {
	if a.id == "" && filter == nil {
		return errors.New("must provide id and/or filter")
	}

	collection, err := getDBCollection(a.ctx, a.db, a.collectionName)
	if err != nil {
		return err
	}

	if filter == nil {
		_, err = collection.RemoveDocument(a.ctx, a.id)
		return err
	}

	query := generateArangoQuery(a.collectionName, filter[0], true)
	_, err = executeArangoQuery(a.ctx, a.db, query, 1)
	if err != nil {
		return err
	}

	return nil
}

func (a ArangoDB) Query(query string, bindParams map[string]interface{}) (interface{}, error) {
	_, err := getDBCollection(a.ctx, a.db, a.collectionName)
	if err != nil {
		return nil, err
	}

	return executeArangoQuery(a.ctx, a.db, &Query{queryString: query, bindVars: bindParams}, -1)
}
