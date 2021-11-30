// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cq-provider-aws/client (interfaces: CodebuildClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	codebuild "github.com/aws/aws-sdk-go-v2/service/codebuild"
	gomock "github.com/golang/mock/gomock"
)

// MockCodebuildClient is a mock of CodebuildClient interface.
type MockCodebuildClient struct {
	ctrl     *gomock.Controller
	recorder *MockCodebuildClientMockRecorder
}

// MockCodebuildClientMockRecorder is the mock recorder for MockCodebuildClient.
type MockCodebuildClientMockRecorder struct {
	mock *MockCodebuildClient
}

// NewMockCodebuildClient creates a new mock instance.
func NewMockCodebuildClient(ctrl *gomock.Controller) *MockCodebuildClient {
	mock := &MockCodebuildClient{ctrl: ctrl}
	mock.recorder = &MockCodebuildClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCodebuildClient) EXPECT() *MockCodebuildClientMockRecorder {
	return m.recorder
}

// BatchGetProjects mocks base method.
func (m *MockCodebuildClient) BatchGetProjects(arg0 context.Context, arg1 *codebuild.BatchGetProjectsInput, arg2 ...func(*codebuild.Options)) (*codebuild.BatchGetProjectsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchGetProjects", varargs...)
	ret0, _ := ret[0].(*codebuild.BatchGetProjectsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchGetProjects indicates an expected call of BatchGetProjects.
func (mr *MockCodebuildClientMockRecorder) BatchGetProjects(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGetProjects", reflect.TypeOf((*MockCodebuildClient)(nil).BatchGetProjects), varargs...)
}

// ListProjects mocks base method.
func (m *MockCodebuildClient) ListProjects(arg0 context.Context, arg1 *codebuild.ListProjectsInput, arg2 ...func(*codebuild.Options)) (*codebuild.ListProjectsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListProjects", varargs...)
	ret0, _ := ret[0].(*codebuild.ListProjectsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProjects indicates an expected call of ListProjects.
func (mr *MockCodebuildClientMockRecorder) ListProjects(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjects", reflect.TypeOf((*MockCodebuildClient)(nil).ListProjects), varargs...)
}