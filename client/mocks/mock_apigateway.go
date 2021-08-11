// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cq-provider-aws/client (interfaces: ApigatewayClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	apigateway "github.com/aws/aws-sdk-go-v2/service/apigateway"
	gomock "github.com/golang/mock/gomock"
)

// MockApigatewayClient is a mock of ApigatewayClient interface.
type MockApigatewayClient struct {
	ctrl     *gomock.Controller
	recorder *MockApigatewayClientMockRecorder
}

// MockApigatewayClientMockRecorder is the mock recorder for MockApigatewayClient.
type MockApigatewayClientMockRecorder struct {
	mock *MockApigatewayClient
}

// NewMockApigatewayClient creates a new mock instance.
func NewMockApigatewayClient(ctrl *gomock.Controller) *MockApigatewayClient {
	mock := &MockApigatewayClient{ctrl: ctrl}
	mock.recorder = &MockApigatewayClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApigatewayClient) EXPECT() *MockApigatewayClientMockRecorder {
	return m.recorder
}

// GetApiKeys mocks base method.
func (m *MockApigatewayClient) GetApiKeys(arg0 context.Context, arg1 *apigateway.GetApiKeysInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetApiKeysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetApiKeys", varargs...)
	ret0, _ := ret[0].(*apigateway.GetApiKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApiKeys indicates an expected call of GetApiKeys.
func (mr *MockApigatewayClientMockRecorder) GetApiKeys(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApiKeys", reflect.TypeOf((*MockApigatewayClient)(nil).GetApiKeys), varargs...)
}

// GetAuthorizers mocks base method.
func (m *MockApigatewayClient) GetAuthorizers(arg0 context.Context, arg1 *apigateway.GetAuthorizersInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetAuthorizersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAuthorizers", varargs...)
	ret0, _ := ret[0].(*apigateway.GetAuthorizersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthorizers indicates an expected call of GetAuthorizers.
func (mr *MockApigatewayClientMockRecorder) GetAuthorizers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthorizers", reflect.TypeOf((*MockApigatewayClient)(nil).GetAuthorizers), varargs...)
}

// GetBasePathMappings mocks base method.
func (m *MockApigatewayClient) GetBasePathMappings(arg0 context.Context, arg1 *apigateway.GetBasePathMappingsInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetBasePathMappingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBasePathMappings", varargs...)
	ret0, _ := ret[0].(*apigateway.GetBasePathMappingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBasePathMappings indicates an expected call of GetBasePathMappings.
func (mr *MockApigatewayClientMockRecorder) GetBasePathMappings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBasePathMappings", reflect.TypeOf((*MockApigatewayClient)(nil).GetBasePathMappings), varargs...)
}

// GetClientCertificates mocks base method.
func (m *MockApigatewayClient) GetClientCertificates(arg0 context.Context, arg1 *apigateway.GetClientCertificatesInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetClientCertificatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetClientCertificates", varargs...)
	ret0, _ := ret[0].(*apigateway.GetClientCertificatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClientCertificates indicates an expected call of GetClientCertificates.
func (mr *MockApigatewayClientMockRecorder) GetClientCertificates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClientCertificates", reflect.TypeOf((*MockApigatewayClient)(nil).GetClientCertificates), varargs...)
}

// GetDeployments mocks base method.
func (m *MockApigatewayClient) GetDeployments(arg0 context.Context, arg1 *apigateway.GetDeploymentsInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetDeploymentsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDeployments", varargs...)
	ret0, _ := ret[0].(*apigateway.GetDeploymentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeployments indicates an expected call of GetDeployments.
func (mr *MockApigatewayClientMockRecorder) GetDeployments(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeployments", reflect.TypeOf((*MockApigatewayClient)(nil).GetDeployments), varargs...)
}

// GetDocumentationParts mocks base method.
func (m *MockApigatewayClient) GetDocumentationParts(arg0 context.Context, arg1 *apigateway.GetDocumentationPartsInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetDocumentationPartsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDocumentationParts", varargs...)
	ret0, _ := ret[0].(*apigateway.GetDocumentationPartsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDocumentationParts indicates an expected call of GetDocumentationParts.
func (mr *MockApigatewayClientMockRecorder) GetDocumentationParts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDocumentationParts", reflect.TypeOf((*MockApigatewayClient)(nil).GetDocumentationParts), varargs...)
}

// GetDocumentationVersions mocks base method.
func (m *MockApigatewayClient) GetDocumentationVersions(arg0 context.Context, arg1 *apigateway.GetDocumentationVersionsInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetDocumentationVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDocumentationVersions", varargs...)
	ret0, _ := ret[0].(*apigateway.GetDocumentationVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDocumentationVersions indicates an expected call of GetDocumentationVersions.
func (mr *MockApigatewayClientMockRecorder) GetDocumentationVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDocumentationVersions", reflect.TypeOf((*MockApigatewayClient)(nil).GetDocumentationVersions), varargs...)
}

// GetDomainNames mocks base method.
func (m *MockApigatewayClient) GetDomainNames(arg0 context.Context, arg1 *apigateway.GetDomainNamesInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetDomainNamesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDomainNames", varargs...)
	ret0, _ := ret[0].(*apigateway.GetDomainNamesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDomainNames indicates an expected call of GetDomainNames.
func (mr *MockApigatewayClientMockRecorder) GetDomainNames(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDomainNames", reflect.TypeOf((*MockApigatewayClient)(nil).GetDomainNames), varargs...)
}

// GetGatewayResponses mocks base method.
func (m *MockApigatewayClient) GetGatewayResponses(arg0 context.Context, arg1 *apigateway.GetGatewayResponsesInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetGatewayResponsesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetGatewayResponses", varargs...)
	ret0, _ := ret[0].(*apigateway.GetGatewayResponsesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGatewayResponses indicates an expected call of GetGatewayResponses.
func (mr *MockApigatewayClientMockRecorder) GetGatewayResponses(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGatewayResponses", reflect.TypeOf((*MockApigatewayClient)(nil).GetGatewayResponses), varargs...)
}

// GetModelTemplate mocks base method.
func (m *MockApigatewayClient) GetModelTemplate(arg0 context.Context, arg1 *apigateway.GetModelTemplateInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetModelTemplateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetModelTemplate", varargs...)
	ret0, _ := ret[0].(*apigateway.GetModelTemplateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetModelTemplate indicates an expected call of GetModelTemplate.
func (mr *MockApigatewayClientMockRecorder) GetModelTemplate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModelTemplate", reflect.TypeOf((*MockApigatewayClient)(nil).GetModelTemplate), varargs...)
}

// GetModels mocks base method.
func (m *MockApigatewayClient) GetModels(arg0 context.Context, arg1 *apigateway.GetModelsInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetModelsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetModels", varargs...)
	ret0, _ := ret[0].(*apigateway.GetModelsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetModels indicates an expected call of GetModels.
func (mr *MockApigatewayClientMockRecorder) GetModels(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModels", reflect.TypeOf((*MockApigatewayClient)(nil).GetModels), varargs...)
}

// GetRequestValidators mocks base method.
func (m *MockApigatewayClient) GetRequestValidators(arg0 context.Context, arg1 *apigateway.GetRequestValidatorsInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetRequestValidatorsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRequestValidators", varargs...)
	ret0, _ := ret[0].(*apigateway.GetRequestValidatorsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRequestValidators indicates an expected call of GetRequestValidators.
func (mr *MockApigatewayClientMockRecorder) GetRequestValidators(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequestValidators", reflect.TypeOf((*MockApigatewayClient)(nil).GetRequestValidators), varargs...)
}

// GetResources mocks base method.
func (m *MockApigatewayClient) GetResources(arg0 context.Context, arg1 *apigateway.GetResourcesInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetResourcesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetResources", varargs...)
	ret0, _ := ret[0].(*apigateway.GetResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResources indicates an expected call of GetResources.
func (mr *MockApigatewayClientMockRecorder) GetResources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResources", reflect.TypeOf((*MockApigatewayClient)(nil).GetResources), varargs...)
}

// GetRestApis mocks base method.
func (m *MockApigatewayClient) GetRestApis(arg0 context.Context, arg1 *apigateway.GetRestApisInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetRestApisOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRestApis", varargs...)
	ret0, _ := ret[0].(*apigateway.GetRestApisOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRestApis indicates an expected call of GetRestApis.
func (mr *MockApigatewayClientMockRecorder) GetRestApis(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRestApis", reflect.TypeOf((*MockApigatewayClient)(nil).GetRestApis), varargs...)
}

// GetStages mocks base method.
func (m *MockApigatewayClient) GetStages(arg0 context.Context, arg1 *apigateway.GetStagesInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetStagesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStages", varargs...)
	ret0, _ := ret[0].(*apigateway.GetStagesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStages indicates an expected call of GetStages.
func (mr *MockApigatewayClientMockRecorder) GetStages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStages", reflect.TypeOf((*MockApigatewayClient)(nil).GetStages), varargs...)
}

// GetUsagePlanKeys mocks base method.
func (m *MockApigatewayClient) GetUsagePlanKeys(arg0 context.Context, arg1 *apigateway.GetUsagePlanKeysInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetUsagePlanKeysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUsagePlanKeys", varargs...)
	ret0, _ := ret[0].(*apigateway.GetUsagePlanKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsagePlanKeys indicates an expected call of GetUsagePlanKeys.
func (mr *MockApigatewayClientMockRecorder) GetUsagePlanKeys(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsagePlanKeys", reflect.TypeOf((*MockApigatewayClient)(nil).GetUsagePlanKeys), varargs...)
}

// GetUsagePlans mocks base method.
func (m *MockApigatewayClient) GetUsagePlans(arg0 context.Context, arg1 *apigateway.GetUsagePlansInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetUsagePlansOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUsagePlans", varargs...)
	ret0, _ := ret[0].(*apigateway.GetUsagePlansOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsagePlans indicates an expected call of GetUsagePlans.
func (mr *MockApigatewayClientMockRecorder) GetUsagePlans(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsagePlans", reflect.TypeOf((*MockApigatewayClient)(nil).GetUsagePlans), varargs...)
}

// GetVpcLinks mocks base method.
func (m *MockApigatewayClient) GetVpcLinks(arg0 context.Context, arg1 *apigateway.GetVpcLinksInput, arg2 ...func(*apigateway.Options)) (*apigateway.GetVpcLinksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetVpcLinks", varargs...)
	ret0, _ := ret[0].(*apigateway.GetVpcLinksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVpcLinks indicates an expected call of GetVpcLinks.
func (mr *MockApigatewayClientMockRecorder) GetVpcLinks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVpcLinks", reflect.TypeOf((*MockApigatewayClient)(nil).GetVpcLinks), varargs...)
}