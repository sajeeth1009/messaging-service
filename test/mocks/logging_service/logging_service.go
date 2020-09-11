// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/influenzanet/logging-service/pkg/api (interfaces: LoggingServiceApiClient)

// Package mock_api is a generated GoMock package.
package mock_api

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	api_types "github.com/influenzanet/go-utils/pkg/api_types"
	api "github.com/influenzanet/logging-service/pkg/api"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

// MockLoggingServiceApiClient is a mock of LoggingServiceApiClient interface
type MockLoggingServiceApiClient struct {
	ctrl     *gomock.Controller
	recorder *MockLoggingServiceApiClientMockRecorder
}

// MockLoggingServiceApiClientMockRecorder is the mock recorder for MockLoggingServiceApiClient
type MockLoggingServiceApiClientMockRecorder struct {
	mock *MockLoggingServiceApiClient
}

// NewMockLoggingServiceApiClient creates a new mock instance
func NewMockLoggingServiceApiClient(ctrl *gomock.Controller) *MockLoggingServiceApiClient {
	mock := &MockLoggingServiceApiClient{ctrl: ctrl}
	mock.recorder = &MockLoggingServiceApiClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLoggingServiceApiClient) EXPECT() *MockLoggingServiceApiClientMockRecorder {
	return m.recorder
}

// GetLogs mocks base method
func (m *MockLoggingServiceApiClient) GetLogs(arg0 context.Context, arg1 *api.LogQuery, arg2 ...grpc.CallOption) (api.LoggingServiceApi_GetLogsClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLogs", varargs...)
	ret0, _ := ret[0].(api.LoggingServiceApi_GetLogsClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogs indicates an expected call of GetLogs
func (mr *MockLoggingServiceApiClientMockRecorder) GetLogs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogs", reflect.TypeOf((*MockLoggingServiceApiClient)(nil).GetLogs), varargs...)
}

// SaveLogEvent mocks base method
func (m *MockLoggingServiceApiClient) SaveLogEvent(arg0 context.Context, arg1 *api.NewLogEvent, arg2 ...grpc.CallOption) (*api_types.ServiceStatus, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SaveLogEvent", varargs...)
	ret0, _ := ret[0].(*api_types.ServiceStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveLogEvent indicates an expected call of SaveLogEvent
func (mr *MockLoggingServiceApiClientMockRecorder) SaveLogEvent(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveLogEvent", reflect.TypeOf((*MockLoggingServiceApiClient)(nil).SaveLogEvent), varargs...)
}

// Status mocks base method
func (m *MockLoggingServiceApiClient) Status(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*api_types.ServiceStatus, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Status", varargs...)
	ret0, _ := ret[0].(*api_types.ServiceStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status
func (mr *MockLoggingServiceApiClientMockRecorder) Status(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockLoggingServiceApiClient)(nil).Status), varargs...)
}