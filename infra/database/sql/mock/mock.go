// Code generated by MockGen. DO NOT EDIT.
// Source: infra/database/sql/database.go

// Package mock is a generated GoMock package.
package mock

import (
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sql0 "github.com/masudur-rahman/pawsitively-purrfect/infra/database/sql"
)

// MockDatabase is a mock of Database interface.
type MockDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseMockRecorder
}

// MockDatabaseMockRecorder is the mock recorder for MockDatabase.
type MockDatabaseMockRecorder struct {
	mock *MockDatabase
}

// NewMockDatabase creates a new mock instance.
func NewMockDatabase(ctrl *gomock.Controller) *MockDatabase {
	mock := &MockDatabase{ctrl: ctrl}
	mock.recorder = &MockDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabase) EXPECT() *MockDatabaseMockRecorder {
	return m.recorder
}

// DeleteOne mocks base method.
func (m *MockDatabase) DeleteOne() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOne")
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOne indicates an expected call of DeleteOne.
func (mr *MockDatabaseMockRecorder) DeleteOne() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOne", reflect.TypeOf((*MockDatabase)(nil).DeleteOne))
}

// Exec mocks base method.
func (m *MockDatabase) Exec(query string, args ...interface{}) (sql.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockDatabaseMockRecorder) Exec(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockDatabase)(nil).Exec), varargs...)
}

// FindMany mocks base method.
func (m *MockDatabase) FindMany(documents interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindMany", documents)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindMany indicates an expected call of FindMany.
func (mr *MockDatabaseMockRecorder) FindMany(documents interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindMany", reflect.TypeOf((*MockDatabase)(nil).FindMany), documents)
}

// FindOne mocks base method.
func (m *MockDatabase) FindOne(document interface{}) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", document)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne.
func (mr *MockDatabaseMockRecorder) FindOne(document interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockDatabase)(nil).FindOne), document)
}

// ID mocks base method.
func (m *MockDatabase) ID(id string) sql0.Database {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID", id)
	ret0, _ := ret[0].(sql0.Database)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockDatabaseMockRecorder) ID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockDatabase)(nil).ID), id)
}

// InsertMany mocks base method.
func (m *MockDatabase) InsertMany(documents []interface{}) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertMany", documents)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertMany indicates an expected call of InsertMany.
func (mr *MockDatabaseMockRecorder) InsertMany(documents interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertMany", reflect.TypeOf((*MockDatabase)(nil).InsertMany), documents)
}

// InsertOne mocks base method.
func (m *MockDatabase) InsertOne(document interface{}) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertOne", document)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertOne indicates an expected call of InsertOne.
func (mr *MockDatabaseMockRecorder) InsertOne(document interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOne", reflect.TypeOf((*MockDatabase)(nil).InsertOne), document)
}

// Query mocks base method.
func (m *MockDatabase) Query(query string, args ...interface{}) (*sql.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(*sql.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockDatabaseMockRecorder) Query(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockDatabase)(nil).Query), varargs...)
}

// SetFilter mocks base method.
func (m *MockDatabase) SetFilter(filter string, args ...interface{}) sql0.Database {
	m.ctrl.T.Helper()
	varargs := []interface{}{filter}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetFilter", varargs...)
	ret0, _ := ret[0].(sql0.Database)
	return ret0
}

// SetFilter indicates an expected call of SetFilter.
func (mr *MockDatabaseMockRecorder) SetFilter(filter interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{filter}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFilter", reflect.TypeOf((*MockDatabase)(nil).SetFilter), varargs...)
}

// UpdateOne mocks base method.
func (m *MockDatabase) UpdateOne(document interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOne", document)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOne indicates an expected call of UpdateOne.
func (mr *MockDatabaseMockRecorder) UpdateOne(document interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOne", reflect.TypeOf((*MockDatabase)(nil).UpdateOne), document)
}