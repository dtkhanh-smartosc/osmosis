// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/osmosis-labs/osmosis/v15/x/cosmwasmpool/types (interfaces: ContractKeeper)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	types "github.com/cosmos/cosmos-sdk/types"
	gomock "github.com/golang/mock/gomock"
)

// MockContractKeeper is a mock of ContractKeeper interface.
type MockContractKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockContractKeeperMockRecorder
}

// MockContractKeeperMockRecorder is the mock recorder for MockContractKeeper.
type MockContractKeeperMockRecorder struct {
	mock *MockContractKeeper
}

// NewMockContractKeeper creates a new mock instance.
func NewMockContractKeeper(ctrl *gomock.Controller) *MockContractKeeper {
	mock := &MockContractKeeper{ctrl: ctrl}
	mock.recorder = &MockContractKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContractKeeper) EXPECT() *MockContractKeeperMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockContractKeeper) Execute(arg0 types.Context, arg1, arg2 types.AccAddress, arg3 []byte, arg4 types.Coins) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockContractKeeperMockRecorder) Execute(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockContractKeeper)(nil).Execute), arg0, arg1, arg2, arg3, arg4)
}

// Instantiate mocks base method.
func (m *MockContractKeeper) Instantiate(arg0 types.Context, arg1 uint64, arg2, arg3 types.AccAddress, arg4 []byte, arg5 string, arg6 types.Coins) (types.AccAddress, []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Instantiate", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(types.AccAddress)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Instantiate indicates an expected call of Instantiate.
func (mr *MockContractKeeperMockRecorder) Instantiate(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Instantiate", reflect.TypeOf((*MockContractKeeper)(nil).Instantiate), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// Sudo mocks base method.
func (m *MockContractKeeper) Sudo(arg0 types.Context, arg1 types.AccAddress, arg2 []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sudo", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sudo indicates an expected call of Sudo.
func (mr *MockContractKeeperMockRecorder) Sudo(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sudo", reflect.TypeOf((*MockContractKeeper)(nil).Sudo), arg0, arg1, arg2)
}
