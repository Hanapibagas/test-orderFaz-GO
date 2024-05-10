// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	logistic "orderfaz-1/features/logistic"

	mock "github.com/stretchr/testify/mock"
)

// LogisticData is an autogenerated mock type for the LogisticDataInterface type
type LogisticData struct {
	mock.Mock
}

// CreateLogistic provides a mock function with given fields: input
func (_m *LogisticData) CreateLogistic(input logistic.CoreLogistic) error {
	ret := _m.Called(input)

	if len(ret) == 0 {
		panic("no return value specified for CreateLogistic")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(logistic.CoreLogistic) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SerachByQuery provides a mock function with given fields: origin_name, dastination_name
func (_m *LogisticData) SerachByQuery(origin_name string, dastination_name string) ([]logistic.CoreLogistic, error) {
	ret := _m.Called(origin_name, dastination_name)

	if len(ret) == 0 {
		panic("no return value specified for SerachByQuery")
	}

	var r0 []logistic.CoreLogistic
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) ([]logistic.CoreLogistic, error)); ok {
		return rf(origin_name, dastination_name)
	}
	if rf, ok := ret.Get(0).(func(string, string) []logistic.CoreLogistic); ok {
		r0 = rf(origin_name, dastination_name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]logistic.CoreLogistic)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(origin_name, dastination_name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewLogisticData creates a new instance of LogisticData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLogisticData(t interface {
	mock.TestingT
	Cleanup(func())
}) *LogisticData {
	mock := &LogisticData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
