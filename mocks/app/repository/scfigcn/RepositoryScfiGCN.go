// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	mongo "project-name/app/models/mongo"

	mock "github.com/stretchr/testify/mock"
)

// RepositoryScfiGCN is an autogenerated mock type for the RepositoryScfiGCN type
type RepositoryScfiGCN struct {
	mock.Mock
}

// ErrorReadMongo provides a mock function with given fields: in
func (_m *RepositoryScfiGCN) ErrorReadMongo(in string) (mongo.ScifGCN, error) {
	ret := _m.Called(in)

	var r0 mongo.ScifGCN
	if rf, ok := ret.Get(0).(func(string) mongo.ScifGCN); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Get(0).(mongo.ScifGCN)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadMongo provides a mock function with given fields: in
func (_m *RepositoryScfiGCN) ReadMongo(in string) (mongo.ScifGCN, error) {
	ret := _m.Called(in)

	var r0 mongo.ScifGCN
	if rf, ok := ret.Get(0).(func(string) mongo.ScifGCN); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Get(0).(mongo.ScifGCN)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
