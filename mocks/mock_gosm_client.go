// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"

	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	gosm "github.com/thal0x/gosm"

	mock "github.com/stretchr/testify/mock"

	types "github.com/CosmWasm/wasmd/x/wasm/types"
)

// GosmClient is an autogenerated mock type for the GosmClient type
type GosmClient struct {
	mock.Mock
}

// ABCIQuery provides a mock function with given fields: ctx, path, req
func (_m *GosmClient) ABCIQuery(ctx context.Context, path string, req gosm.Marshaler) (*coretypes.ResultABCIQuery, error) {
	ret := _m.Called(ctx, path, req)

	if len(ret) == 0 {
		panic("no return value specified for ABCIQuery")
	}

	var r0 *coretypes.ResultABCIQuery
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, gosm.Marshaler) (*coretypes.ResultABCIQuery, error)); ok {
		return rf(ctx, path, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, gosm.Marshaler) *coretypes.ResultABCIQuery); ok {
		r0 = rf(ctx, path, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coretypes.ResultABCIQuery)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, gosm.Marshaler) error); ok {
		r1 = rf(ctx, path, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Block provides a mock function with given fields: ctx, height
func (_m *GosmClient) Block(ctx context.Context, height *int64) (*coretypes.ResultBlock, error) {
	ret := _m.Called(ctx, height)

	if len(ret) == 0 {
		panic("no return value specified for Block")
	}

	var r0 *coretypes.ResultBlock
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *int64) (*coretypes.ResultBlock, error)); ok {
		return rf(ctx, height)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *int64) *coretypes.ResultBlock); ok {
		r0 = rf(ctx, height)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coretypes.ResultBlock)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *int64) error); ok {
		r1 = rf(ctx, height)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlockResults provides a mock function with given fields: ctx, height
func (_m *GosmClient) BlockResults(ctx context.Context, height *int64) (*coretypes.ResultBlockResults, error) {
	ret := _m.Called(ctx, height)

	if len(ret) == 0 {
		panic("no return value specified for BlockResults")
	}

	var r0 *coretypes.ResultBlockResults
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *int64) (*coretypes.ResultBlockResults, error)); ok {
		return rf(ctx, height)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *int64) *coretypes.ResultBlockResults); ok {
		r0 = rf(ctx, height)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coretypes.ResultBlockResults)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *int64) error); ok {
		r1 = rf(ctx, height)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CometLegacyEncoding provides a mock function with given fields:
func (_m *GosmClient) CometLegacyEncoding() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CometLegacyEncoding")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// QuerySmartContractState provides a mock function with given fields: ctx, address, query
func (_m *GosmClient) QuerySmartContractState(ctx context.Context, address string, query interface{}) (*types.QuerySmartContractStateResponse, error) {
	ret := _m.Called(ctx, address, query)

	if len(ret) == 0 {
		panic("no return value specified for QuerySmartContractState")
	}

	var r0 *types.QuerySmartContractStateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) (*types.QuerySmartContractStateResponse, error)); ok {
		return rf(ctx, address, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) *types.QuerySmartContractStateResponse); ok {
		r0 = rf(ctx, address, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QuerySmartContractStateResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, interface{}) error); ok {
		r1 = rf(ctx, address, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewGosmClient creates a new instance of GosmClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGosmClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *GosmClient {
	mock := &GosmClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
