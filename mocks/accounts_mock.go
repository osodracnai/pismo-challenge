// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/server/accounts/accounts.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	database "github.com/osodracnai/pismo-challenge/pkg/database"
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

// GetAccountByID mocks base method.
func (m *MockDatabase) GetAccountByID(ctx context.Context, id string) (*database.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountByID", ctx, id)
	ret0, _ := ret[0].(*database.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountByID indicates an expected call of GetAccountByID.
func (mr *MockDatabaseMockRecorder) GetAccountByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountByID", reflect.TypeOf((*MockDatabase)(nil).GetAccountByID), ctx, id)
}

// InsertAccount mocks base method.
func (m *MockDatabase) InsertAccount(ctx context.Context, account database.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertAccount", ctx, account)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertAccount indicates an expected call of InsertAccount.
func (mr *MockDatabaseMockRecorder) InsertAccount(ctx, account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertAccount", reflect.TypeOf((*MockDatabase)(nil).InsertAccount), ctx, account)
}
