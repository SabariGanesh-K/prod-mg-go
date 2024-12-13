// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/SabariGanesh-K/prod-mgm-go/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	db "github.com/SabariGanesh-K/prod-mgm-go/db/sqlc"
	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// AddCompressedProductImageUrlsByID mocks base method.
func (m *MockStore) AddCompressedProductImageUrlsByID(arg0 context.Context, arg1 db.AddCompressedProductImageUrlsByIDParams) (db.Products, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCompressedProductImageUrlsByID", arg0, arg1)
	ret0, _ := ret[0].(db.Products)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddCompressedProductImageUrlsByID indicates an expected call of AddCompressedProductImageUrlsByID.
func (mr *MockStoreMockRecorder) AddCompressedProductImageUrlsByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCompressedProductImageUrlsByID", reflect.TypeOf((*MockStore)(nil).AddCompressedProductImageUrlsByID), arg0, arg1)
}

// CreateProduct mocks base method.
func (m *MockStore) CreateProduct(arg0 context.Context, arg1 db.CreateProductParams) (db.Products, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", arg0, arg1)
	ret0, _ := ret[0].(db.Products)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockStoreMockRecorder) CreateProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockStore)(nil).CreateProduct), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// GetProductByProductID mocks base method.
func (m *MockStore) GetProductByProductID(arg0 context.Context, arg1 string) (db.Products, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductByProductID", arg0, arg1)
	ret0, _ := ret[0].(db.Products)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductByProductID indicates an expected call of GetProductByProductID.
func (mr *MockStoreMockRecorder) GetProductByProductID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductByProductID", reflect.TypeOf((*MockStore)(nil).GetProductByProductID), arg0, arg1)
}

// GetProductsByUserID mocks base method.
func (m *MockStore) GetProductsByUserID(arg0 context.Context, arg1 db.GetProductsByUserIDParams) ([]db.Products, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductsByUserID", arg0, arg1)
	ret0, _ := ret[0].([]db.Products)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductsByUserID indicates an expected call of GetProductsByUserID.
func (mr *MockStoreMockRecorder) GetProductsByUserID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductsByUserID", reflect.TypeOf((*MockStore)(nil).GetProductsByUserID), arg0, arg1)
}

// GetUserByID mocks base method.
func (m *MockStore) GetUserByID(arg0 context.Context, arg1 string) (db.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", arg0, arg1)
	ret0, _ := ret[0].(db.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockStoreMockRecorder) GetUserByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockStore)(nil).GetUserByID), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) (db.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(db.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0, arg1)
}
