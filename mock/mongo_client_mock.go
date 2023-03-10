// Code generated by MockGen. DO NOT EDIT.
// Source: mongodb/mongodb.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mongo "go.mongodb.org/mongo-driver/mongo"
	options "go.mongodb.org/mongo-driver/mongo/options"
)

// MockmongoClient is a mock of mongoClient interface.
type MockmongoClient struct {
	ctrl     *gomock.Controller
	recorder *MockmongoClientMockRecorder
}

// MockmongoClientMockRecorder is the mock recorder for MockmongoClient.
type MockmongoClientMockRecorder struct {
	mock *MockmongoClient
}

// NewMockmongoClient creates a new mock instance.
func NewMockmongoClient(ctrl *gomock.Controller) *MockmongoClient {
	mock := &MockmongoClient{ctrl: ctrl}
	mock.recorder = &MockmongoClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockmongoClient) EXPECT() *MockmongoClientMockRecorder {
	return m.recorder
}

// DeleteOne mocks base method.
func (m *MockmongoClient) DeleteOne(arg0 context.Context, arg1 interface{}, arg2 ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteOne", varargs...)
	ret0, _ := ret[0].(*mongo.DeleteResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteOne indicates an expected call of DeleteOne.
func (mr *MockmongoClientMockRecorder) DeleteOne(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOne", reflect.TypeOf((*MockmongoClient)(nil).DeleteOne), varargs...)
}

// Find mocks base method.
func (m *MockmongoClient) Find(arg0 context.Context, arg1 interface{}, arg2 ...*options.FindOptions) (*mongo.Cursor, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Find", varargs...)
	ret0, _ := ret[0].(*mongo.Cursor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockmongoClientMockRecorder) Find(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockmongoClient)(nil).Find), varargs...)
}

// FindOne mocks base method.
func (m *MockmongoClient) FindOne(arg0 context.Context, arg1 interface{}, arg2 ...*options.FindOneOptions) *mongo.SingleResult {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindOne", varargs...)
	ret0, _ := ret[0].(*mongo.SingleResult)
	return ret0
}

// FindOne indicates an expected call of FindOne.
func (mr *MockmongoClientMockRecorder) FindOne(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockmongoClient)(nil).FindOne), varargs...)
}

// InsertOne mocks base method.
func (m *MockmongoClient) InsertOne(arg0 context.Context, arg1 interface{}, arg2 ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InsertOne", varargs...)
	ret0, _ := ret[0].(*mongo.InsertOneResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertOne indicates an expected call of InsertOne.
func (mr *MockmongoClientMockRecorder) InsertOne(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOne", reflect.TypeOf((*MockmongoClient)(nil).InsertOne), varargs...)
}

// UpdateOne mocks base method.
func (m *MockmongoClient) UpdateOne(arg0 context.Context, arg1, arg2 interface{}, arg3 ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateOne", varargs...)
	ret0, _ := ret[0].(*mongo.UpdateResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOne indicates an expected call of UpdateOne.
func (mr *MockmongoClientMockRecorder) UpdateOne(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOne", reflect.TypeOf((*MockmongoClient)(nil).UpdateOne), varargs...)
}
