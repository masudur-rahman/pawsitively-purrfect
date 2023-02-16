package _map

import (
	"errors"
	"reflect"

	"github.com/masudur-rahman/pawsitively-purrfect/infra/database/mock"

	"github.com/rs/xid"
)

type MockDB struct {
	db     map[string]map[string]interface{}
	entity string
	id     string
}

func NewMockDB() *MockDB {
	return &MockDB{
		db: map[string]map[string]interface{}{},
	}
}

func (m *MockDB) Entity(name string) mock.Database {
	m.entity = name
	return m
}

func (m *MockDB) ID(id string) mock.Database {
	m.id = id
	return m
}

func (m *MockDB) FindOne(document interface{}, filter ...interface{}) (bool, error) {
	if m.entity == "" {
		return false, errors.New("must set entity")
	}
	if m.id == "" && filter == nil {
		return false, errors.New("must provide id and/or filter")
	}
	if filter == nil {
		doc, ok := m.db[m.entity][m.id]
		if !ok {
			return ok, nil
		}
		reflect.ValueOf(document).Elem().Set(reflect.ValueOf(doc))
	}

	return false, nil
}

func (m *MockDB) FindMany(documents interface{}, filter interface{}) error {
	if m.entity == "" {
		return errors.New("must set entity")
	}

	reflect.ValueOf(documents).Elem().Set(reflect.ValueOf(documents))
	return nil
}

func (m *MockDB) InsertOne(document interface{}) (id string, err error) {
	id = xid.New().String()
	reflect.ValueOf(document).Elem().FieldByName("ID").SetString(id)
	m.db[m.entity][id] = document
	return
}

func (m *MockDB) InsertMany(documents []interface{}) ([]string, error) {
	var ids []string
	for document := range documents {
		id := xid.New().String()
		reflect.ValueOf(document).Elem().FieldByName("ID").SetString(id)
		m.db[m.entity][id] = document
		ids = append(ids, id)
	}
	return ids, nil
}

func (m *MockDB) UpdateOne(document interface{}) error {
	if m.entity == "" {
		return errors.New("must set entity")
	}
	if m.id == "" {
		return errors.New("must provide id")
	}
	m.db[m.entity][m.id] = document
	return nil
}

func (m *MockDB) DeleteOne(filter ...interface{}) error {
	delete(m.db[m.entity], m.id)
	return nil
}
