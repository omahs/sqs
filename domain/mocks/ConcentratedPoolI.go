// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	types "github.com/osmosis-labs/osmosis/v19/x/poolmanager/types"
	mock "github.com/stretchr/testify/mock"
)

// ConcentratedPoolI is an autogenerated mock type for the ConcentratedPoolI type
type ConcentratedPoolI struct {
	mock.Mock
}

type ConcentratedPoolI_Expecter struct {
	mock *mock.Mock
}

func (_m *ConcentratedPoolI) EXPECT() *ConcentratedPoolI_Expecter {
	return &ConcentratedPoolI_Expecter{mock: &_m.Mock}
}

// GetId provides a mock function with given fields:
func (_m *ConcentratedPoolI) GetId() uint64 {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// ConcentratedPoolI_GetId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetId'
type ConcentratedPoolI_GetId_Call struct {
	*mock.Call
}

// GetId is a helper method to define mock.On call
func (_e *ConcentratedPoolI_Expecter) GetId() *ConcentratedPoolI_GetId_Call {
	return &ConcentratedPoolI_GetId_Call{Call: _e.mock.On("GetId")}
}

func (_c *ConcentratedPoolI_GetId_Call) Run(run func()) *ConcentratedPoolI_GetId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ConcentratedPoolI_GetId_Call) Return(_a0 uint64) *ConcentratedPoolI_GetId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConcentratedPoolI_GetId_Call) RunAndReturn(run func() uint64) *ConcentratedPoolI_GetId_Call {
	_c.Call.Return(run)
	return _c
}

// GetType provides a mock function with given fields:
func (_m *ConcentratedPoolI) GetType() types.PoolType {
	ret := _m.Called()

	var r0 types.PoolType
	if rf, ok := ret.Get(0).(func() types.PoolType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.PoolType)
	}

	return r0
}

// ConcentratedPoolI_GetType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetType'
type ConcentratedPoolI_GetType_Call struct {
	*mock.Call
}

// GetType is a helper method to define mock.On call
func (_e *ConcentratedPoolI_Expecter) GetType() *ConcentratedPoolI_GetType_Call {
	return &ConcentratedPoolI_GetType_Call{Call: _e.mock.On("GetType")}
}

func (_c *ConcentratedPoolI_GetType_Call) Run(run func()) *ConcentratedPoolI_GetType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ConcentratedPoolI_GetType_Call) Return(_a0 types.PoolType) *ConcentratedPoolI_GetType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConcentratedPoolI_GetType_Call) RunAndReturn(run func() types.PoolType) *ConcentratedPoolI_GetType_Call {
	_c.Call.Return(run)
	return _c
}

// NewConcentratedPoolI creates a new instance of ConcentratedPoolI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConcentratedPoolI(t interface {
	mock.TestingT
	Cleanup(func())
}) *ConcentratedPoolI {
	mock := &ConcentratedPoolI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
