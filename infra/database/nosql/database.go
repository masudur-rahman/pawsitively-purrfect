package nosql

type Database interface {
	ID(id string) Database
	Collection(name string) Database

	FindOne(document interface{}, filter interface{}) (bool, error)
	FindMany(documents interface{}, filter interface{}) error

	InsertOne(document interface{}) (id string, err error)
	InsertMany(documents []interface{}) ([]string, error)

	UpdateOne(document interface{}) error

	DeleteOne() error

	Query(query string, bindParams map[string]interface{}) (interface{}, error)
}
