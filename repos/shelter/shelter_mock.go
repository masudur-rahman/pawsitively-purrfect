// Code generated by MockGen. DO NOT EDIT.
// Source: repos/shelter.go

// Package shelter is a generated GoMock package.
package shelter

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/masudur-rahman/pawsitively-purrfect/models"
)

// MockShelterRepository is a mock of ShelterRepository interface.
type MockShelterRepository struct {
	ctrl     *gomock.Controller
	recorder *MockShelterRepositoryMockRecorder
}

// MockShelterRepositoryMockRecorder is the mock recorder for MockShelterRepository.
type MockShelterRepositoryMockRecorder struct {
	mock *MockShelterRepository
}

// NewMockShelterRepository creates a new mock instance.
func NewMockShelterRepository(ctrl *gomock.Controller) *MockShelterRepository {
	mock := &MockShelterRepository{ctrl: ctrl}
	mock.recorder = &MockShelterRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockShelterRepository) EXPECT() *MockShelterRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockShelterRepository) Delete(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockShelterRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockShelterRepository)(nil).Delete), id)
}

// FindByID mocks base method.
func (m *MockShelterRepository) FindByID(id string) (*models.Shelter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", id)
	ret0, _ := ret[0].(*models.Shelter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockShelterRepositoryMockRecorder) FindByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockShelterRepository)(nil).FindByID), id)
}

// FindByLocation mocks base method.
func (m *MockShelterRepository) FindByLocation(location string) ([]*models.Shelter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByLocation", location)
	ret0, _ := ret[0].([]*models.Shelter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByLocation indicates an expected call of FindByLocation.
func (mr *MockShelterRepositoryMockRecorder) FindByLocation(location interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByLocation", reflect.TypeOf((*MockShelterRepository)(nil).FindByLocation), location)
}

// FindByName mocks base method.
func (m *MockShelterRepository) FindByName(name string) (*models.Shelter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByName", name)
	ret0, _ := ret[0].(*models.Shelter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByName indicates an expected call of FindByName.
func (mr *MockShelterRepositoryMockRecorder) FindByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByName", reflect.TypeOf((*MockShelterRepository)(nil).FindByName), name)
}

// FindByOwnerID mocks base method.
func (m *MockShelterRepository) FindByOwnerID(id string) ([]*models.Shelter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByOwnerID", id)
	ret0, _ := ret[0].([]*models.Shelter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByOwnerID indicates an expected call of FindByOwnerID.
func (mr *MockShelterRepositoryMockRecorder) FindByOwnerID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByOwnerID", reflect.TypeOf((*MockShelterRepository)(nil).FindByOwnerID), id)
}

// FindShelters mocks base method.
func (m *MockShelterRepository) FindShelters(filter models.Shelter) ([]*models.Shelter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindShelters", filter)
	ret0, _ := ret[0].([]*models.Shelter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindShelters indicates an expected call of FindShelters.
func (mr *MockShelterRepositoryMockRecorder) FindShelters(filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindShelters", reflect.TypeOf((*MockShelterRepository)(nil).FindShelters), filter)
}

// Save mocks base method.
func (m *MockShelterRepository) Save(shelter *models.Shelter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", shelter)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockShelterRepositoryMockRecorder) Save(shelter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockShelterRepository)(nil).Save), shelter)
}

// Update mocks base method.
func (m *MockShelterRepository) Update(shelter *models.Shelter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", shelter)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockShelterRepositoryMockRecorder) Update(shelter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockShelterRepository)(nil).Update), shelter)
}
