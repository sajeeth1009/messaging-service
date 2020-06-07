// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/influenzanet/study-service/pkg/api (interfaces: StudyServiceApiClient)

// Package mock_api is a generated GoMock package.
package mock_api

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	api "github.com/influenzanet/study-service/pkg/api"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

// MockStudyServiceApiClient is a mock of StudyServiceApiClient interface
type MockStudyServiceApiClient struct {
	ctrl     *gomock.Controller
	recorder *MockStudyServiceApiClientMockRecorder
}

// MockStudyServiceApiClientMockRecorder is the mock recorder for MockStudyServiceApiClient
type MockStudyServiceApiClientMockRecorder struct {
	mock *MockStudyServiceApiClient
}

// NewMockStudyServiceApiClient creates a new mock instance
func NewMockStudyServiceApiClient(ctrl *gomock.Controller) *MockStudyServiceApiClient {
	mock := &MockStudyServiceApiClient{ctrl: ctrl}
	mock.recorder = &MockStudyServiceApiClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStudyServiceApiClient) EXPECT() *MockStudyServiceApiClientMockRecorder {
	return m.recorder
}

// CreateNewStudy mocks base method
func (m *MockStudyServiceApiClient) CreateNewStudy(arg0 context.Context, arg1 *api.NewStudyRequest, arg2 ...grpc.CallOption) (*api.Study, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateNewStudy", varargs...)
	ret0, _ := ret[0].(*api.Study)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNewStudy indicates an expected call of CreateNewStudy
func (mr *MockStudyServiceApiClientMockRecorder) CreateNewStudy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewStudy", reflect.TypeOf((*MockStudyServiceApiClient)(nil).CreateNewStudy), varargs...)
}

// EnterStudy mocks base method
func (m *MockStudyServiceApiClient) EnterStudy(arg0 context.Context, arg1 *api.EnterStudyRequest, arg2 ...grpc.CallOption) (*api.AssignedSurveys, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EnterStudy", varargs...)
	ret0, _ := ret[0].(*api.AssignedSurveys)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnterStudy indicates an expected call of EnterStudy
func (mr *MockStudyServiceApiClientMockRecorder) EnterStudy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnterStudy", reflect.TypeOf((*MockStudyServiceApiClient)(nil).EnterStudy), varargs...)
}

// GetActiveStudies mocks base method
func (m *MockStudyServiceApiClient) GetActiveStudies(arg0 context.Context, arg1 *api.TokenInfos, arg2 ...grpc.CallOption) (*api.Studies, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetActiveStudies", varargs...)
	ret0, _ := ret[0].(*api.Studies)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActiveStudies indicates an expected call of GetActiveStudies
func (mr *MockStudyServiceApiClientMockRecorder) GetActiveStudies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActiveStudies", reflect.TypeOf((*MockStudyServiceApiClient)(nil).GetActiveStudies), varargs...)
}

// GetAllStudies mocks base method
func (m *MockStudyServiceApiClient) GetAllStudies(arg0 context.Context, arg1 *api.TokenInfos, arg2 ...grpc.CallOption) (*api.Studies, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAllStudies", varargs...)
	ret0, _ := ret[0].(*api.Studies)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllStudies indicates an expected call of GetAllStudies
func (mr *MockStudyServiceApiClientMockRecorder) GetAllStudies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllStudies", reflect.TypeOf((*MockStudyServiceApiClient)(nil).GetAllStudies), varargs...)
}

// GetAssignedSurvey mocks base method
func (m *MockStudyServiceApiClient) GetAssignedSurvey(arg0 context.Context, arg1 *api.SurveyReferenceRequest, arg2 ...grpc.CallOption) (*api.SurveyAndContext, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAssignedSurvey", varargs...)
	ret0, _ := ret[0].(*api.SurveyAndContext)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAssignedSurvey indicates an expected call of GetAssignedSurvey
func (mr *MockStudyServiceApiClientMockRecorder) GetAssignedSurvey(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAssignedSurvey", reflect.TypeOf((*MockStudyServiceApiClient)(nil).GetAssignedSurvey), varargs...)
}

// GetAssignedSurveys mocks base method
func (m *MockStudyServiceApiClient) GetAssignedSurveys(arg0 context.Context, arg1 *api.TokenInfos, arg2 ...grpc.CallOption) (*api.AssignedSurveys, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAssignedSurveys", varargs...)
	ret0, _ := ret[0].(*api.AssignedSurveys)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAssignedSurveys indicates an expected call of GetAssignedSurveys
func (mr *MockStudyServiceApiClientMockRecorder) GetAssignedSurveys(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAssignedSurveys", reflect.TypeOf((*MockStudyServiceApiClient)(nil).GetAssignedSurveys), varargs...)
}

// GetStudiesForUser mocks base method
func (m *MockStudyServiceApiClient) GetStudiesForUser(arg0 context.Context, arg1 *api.GetStudiesForUserReq, arg2 ...grpc.CallOption) (*api.Studies, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStudiesForUser", varargs...)
	ret0, _ := ret[0].(*api.Studies)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStudiesForUser indicates an expected call of GetStudiesForUser
func (mr *MockStudyServiceApiClientMockRecorder) GetStudiesForUser(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudiesForUser", reflect.TypeOf((*MockStudyServiceApiClient)(nil).GetStudiesForUser), varargs...)
}

// GetStudy mocks base method
func (m *MockStudyServiceApiClient) GetStudy(arg0 context.Context, arg1 *api.StudyReferenceReq, arg2 ...grpc.CallOption) (*api.Study, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStudy", varargs...)
	ret0, _ := ret[0].(*api.Study)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStudy indicates an expected call of GetStudy
func (mr *MockStudyServiceApiClientMockRecorder) GetStudy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudy", reflect.TypeOf((*MockStudyServiceApiClient)(nil).GetStudy), varargs...)
}

// GetStudyResponseStatistics mocks base method
func (m *MockStudyServiceApiClient) GetStudyResponseStatistics(arg0 context.Context, arg1 *api.SurveyResponseQuery, arg2 ...grpc.CallOption) (*api.StudyResponseStatistics, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStudyResponseStatistics", varargs...)
	ret0, _ := ret[0].(*api.StudyResponseStatistics)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStudyResponseStatistics indicates an expected call of GetStudyResponseStatistics
func (mr *MockStudyServiceApiClientMockRecorder) GetStudyResponseStatistics(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudyResponseStatistics", reflect.TypeOf((*MockStudyServiceApiClient)(nil).GetStudyResponseStatistics), varargs...)
}

// GetStudySurveyInfos mocks base method
func (m *MockStudyServiceApiClient) GetStudySurveyInfos(arg0 context.Context, arg1 *api.StudyReferenceReq, arg2 ...grpc.CallOption) (*api.SurveyInfoResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStudySurveyInfos", varargs...)
	ret0, _ := ret[0].(*api.SurveyInfoResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStudySurveyInfos indicates an expected call of GetStudySurveyInfos
func (mr *MockStudyServiceApiClientMockRecorder) GetStudySurveyInfos(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudySurveyInfos", reflect.TypeOf((*MockStudyServiceApiClient)(nil).GetStudySurveyInfos), varargs...)
}

// GetSurveyDefForStudy mocks base method
func (m *MockStudyServiceApiClient) GetSurveyDefForStudy(arg0 context.Context, arg1 *api.SurveyReferenceRequest, arg2 ...grpc.CallOption) (*api.Survey, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSurveyDefForStudy", varargs...)
	ret0, _ := ret[0].(*api.Survey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSurveyDefForStudy indicates an expected call of GetSurveyDefForStudy
func (mr *MockStudyServiceApiClientMockRecorder) GetSurveyDefForStudy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSurveyDefForStudy", reflect.TypeOf((*MockStudyServiceApiClient)(nil).GetSurveyDefForStudy), varargs...)
}

// HasParticipantStateWithCondition mocks base method
func (m *MockStudyServiceApiClient) HasParticipantStateWithCondition(arg0 context.Context, arg1 *api.ProfilesWithConditionReq, arg2 ...grpc.CallOption) (*api.ServiceStatus, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HasParticipantStateWithCondition", varargs...)
	ret0, _ := ret[0].(*api.ServiceStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasParticipantStateWithCondition indicates an expected call of HasParticipantStateWithCondition
func (mr *MockStudyServiceApiClientMockRecorder) HasParticipantStateWithCondition(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasParticipantStateWithCondition", reflect.TypeOf((*MockStudyServiceApiClient)(nil).HasParticipantStateWithCondition), varargs...)
}

// LeaveStudy mocks base method
func (m *MockStudyServiceApiClient) LeaveStudy(arg0 context.Context, arg1 *api.LeaveStudyMsg, arg2 ...grpc.CallOption) (*api.AssignedSurveys, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LeaveStudy", varargs...)
	ret0, _ := ret[0].(*api.AssignedSurveys)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LeaveStudy indicates an expected call of LeaveStudy
func (mr *MockStudyServiceApiClientMockRecorder) LeaveStudy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LeaveStudy", reflect.TypeOf((*MockStudyServiceApiClient)(nil).LeaveStudy), varargs...)
}

// PostponeSurvey mocks base method
func (m *MockStudyServiceApiClient) PostponeSurvey(arg0 context.Context, arg1 *api.PostponeSurveyRequest, arg2 ...grpc.CallOption) (*api.AssignedSurveys, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PostponeSurvey", varargs...)
	ret0, _ := ret[0].(*api.AssignedSurveys)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostponeSurvey indicates an expected call of PostponeSurvey
func (mr *MockStudyServiceApiClientMockRecorder) PostponeSurvey(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostponeSurvey", reflect.TypeOf((*MockStudyServiceApiClient)(nil).PostponeSurvey), varargs...)
}

// RemoveSurveyFromStudy mocks base method
func (m *MockStudyServiceApiClient) RemoveSurveyFromStudy(arg0 context.Context, arg1 *api.SurveyReferenceRequest, arg2 ...grpc.CallOption) (*api.ServiceStatus, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveSurveyFromStudy", varargs...)
	ret0, _ := ret[0].(*api.ServiceStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveSurveyFromStudy indicates an expected call of RemoveSurveyFromStudy
func (mr *MockStudyServiceApiClientMockRecorder) RemoveSurveyFromStudy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSurveyFromStudy", reflect.TypeOf((*MockStudyServiceApiClient)(nil).RemoveSurveyFromStudy), varargs...)
}

// SaveSurveyToStudy mocks base method
func (m *MockStudyServiceApiClient) SaveSurveyToStudy(arg0 context.Context, arg1 *api.AddSurveyReq, arg2 ...grpc.CallOption) (*api.Survey, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SaveSurveyToStudy", varargs...)
	ret0, _ := ret[0].(*api.Survey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveSurveyToStudy indicates an expected call of SaveSurveyToStudy
func (mr *MockStudyServiceApiClientMockRecorder) SaveSurveyToStudy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveSurveyToStudy", reflect.TypeOf((*MockStudyServiceApiClient)(nil).SaveSurveyToStudy), varargs...)
}

// Status mocks base method
func (m *MockStudyServiceApiClient) Status(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*api.ServiceStatus, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Status", varargs...)
	ret0, _ := ret[0].(*api.ServiceStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status
func (mr *MockStudyServiceApiClientMockRecorder) Status(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockStudyServiceApiClient)(nil).Status), varargs...)
}

// StreamStudyResponses mocks base method
func (m *MockStudyServiceApiClient) StreamStudyResponses(arg0 context.Context, arg1 *api.SurveyResponseQuery, arg2 ...grpc.CallOption) (api.StudyServiceApi_StreamStudyResponsesClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StreamStudyResponses", varargs...)
	ret0, _ := ret[0].(api.StudyServiceApi_StreamStudyResponsesClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StreamStudyResponses indicates an expected call of StreamStudyResponses
func (mr *MockStudyServiceApiClientMockRecorder) StreamStudyResponses(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamStudyResponses", reflect.TypeOf((*MockStudyServiceApiClient)(nil).StreamStudyResponses), varargs...)
}

// SubmitResponse mocks base method
func (m *MockStudyServiceApiClient) SubmitResponse(arg0 context.Context, arg1 *api.SubmitResponseReq, arg2 ...grpc.CallOption) (*api.AssignedSurveys, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubmitResponse", varargs...)
	ret0, _ := ret[0].(*api.AssignedSurveys)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitResponse indicates an expected call of SubmitResponse
func (mr *MockStudyServiceApiClientMockRecorder) SubmitResponse(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitResponse", reflect.TypeOf((*MockStudyServiceApiClient)(nil).SubmitResponse), varargs...)
}

// SubmitStatusReport mocks base method
func (m *MockStudyServiceApiClient) SubmitStatusReport(arg0 context.Context, arg1 *api.StatusReportRequest, arg2 ...grpc.CallOption) (*api.AssignedSurveys, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubmitStatusReport", varargs...)
	ret0, _ := ret[0].(*api.AssignedSurveys)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitStatusReport indicates an expected call of SubmitStatusReport
func (mr *MockStudyServiceApiClientMockRecorder) SubmitStatusReport(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitStatusReport", reflect.TypeOf((*MockStudyServiceApiClient)(nil).SubmitStatusReport), varargs...)
}
