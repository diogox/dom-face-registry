// Code generated by MockGen. DO NOT EDIT.
// Source: ./server.go

// Package grpc is a generated GoMock package.
package grpc

import (
	context "context"
	pb "github.com/diogox/dom-face-registry/internal/pb"
	person "github.com/diogox/dom-face-registry/internal/person"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	reflect "reflect"
)

// MockRegistryService is a mock of RegistryService interface
type MockRegistryService struct {
	ctrl     *gomock.Controller
	recorder *MockRegistryServiceMockRecorder
}

// MockRegistryServiceMockRecorder is the mock recorder for MockRegistryService
type MockRegistryServiceMockRecorder struct {
	mock *MockRegistryService
}

// NewMockRegistryService creates a new mock instance
func NewMockRegistryService(ctrl *gomock.Controller) *MockRegistryService {
	mock := &MockRegistryService{ctrl: ctrl}
	mock.recorder = &MockRegistryServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRegistryService) EXPECT() *MockRegistryServiceMockRecorder {
	return m.recorder
}

// GetPeople mocks base method
func (m *MockRegistryService) GetPeople(ctx context.Context) ([]person.Person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPeople", ctx)
	ret0, _ := ret[0].([]person.Person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPeople indicates an expected call of GetPeople
func (mr *MockRegistryServiceMockRecorder) GetPeople(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPeople", reflect.TypeOf((*MockRegistryService)(nil).GetPeople), ctx)
}

// AddPerson mocks base method
func (m *MockRegistryService) AddPerson(ctx context.Context, personInfo person.Person) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPerson", ctx, personInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPerson indicates an expected call of AddPerson
func (mr *MockRegistryServiceMockRecorder) AddPerson(ctx, personInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPerson", reflect.TypeOf((*MockRegistryService)(nil).AddPerson), ctx, personInfo)
}

// RemovePerson mocks base method
func (m *MockRegistryService) RemovePerson(ctx context.Context, personID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemovePerson", ctx, personID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemovePerson indicates an expected call of RemovePerson
func (mr *MockRegistryServiceMockRecorder) RemovePerson(ctx, personID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePerson", reflect.TypeOf((*MockRegistryService)(nil).RemovePerson), ctx, personID)
}

// RecognizeFace mocks base method
func (m *MockRegistryService) RecognizeFace(ctx context.Context, imgData []byte) (person.Person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecognizeFace", ctx, imgData)
	ret0, _ := ret[0].(person.Person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RecognizeFace indicates an expected call of RecognizeFace
func (mr *MockRegistryServiceMockRecorder) RecognizeFace(ctx, imgData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecognizeFace", reflect.TypeOf((*MockRegistryService)(nil).RecognizeFace), ctx, imgData)
}

// AddFace mocks base method
func (m *MockRegistryService) AddFace(ctx context.Context, imgData []byte, personID uuid.UUID) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFace", ctx, imgData, personID)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddFace indicates an expected call of AddFace
func (mr *MockRegistryServiceMockRecorder) AddFace(ctx, imgData, personID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFace", reflect.TypeOf((*MockRegistryService)(nil).AddFace), ctx, imgData, personID)
}

// RemoveFace mocks base method
func (m *MockRegistryService) RemoveFace(ctx context.Context, faceID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFace", ctx, faceID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFace indicates an expected call of RemoveFace
func (mr *MockRegistryServiceMockRecorder) RemoveFace(ctx, faceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFace", reflect.TypeOf((*MockRegistryService)(nil).RemoveFace), ctx, faceID)
}

// MockPeopleConverter is a mock of PeopleConverter interface
type MockPeopleConverter struct {
	ctrl     *gomock.Controller
	recorder *MockPeopleConverterMockRecorder
}

// MockPeopleConverterMockRecorder is the mock recorder for MockPeopleConverter
type MockPeopleConverterMockRecorder struct {
	mock *MockPeopleConverter
}

// NewMockPeopleConverter creates a new mock instance
func NewMockPeopleConverter(ctrl *gomock.Controller) *MockPeopleConverter {
	mock := &MockPeopleConverter{ctrl: ctrl}
	mock.recorder = &MockPeopleConverterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPeopleConverter) EXPECT() *MockPeopleConverterMockRecorder {
	return m.recorder
}

// PersonAsResponse mocks base method
func (m *MockPeopleConverter) PersonAsResponse(resPerson person.Person) *pb.Person {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PersonAsResponse", resPerson)
	ret0, _ := ret[0].(*pb.Person)
	return ret0
}

// PersonAsResponse indicates an expected call of PersonAsResponse
func (mr *MockPeopleConverterMockRecorder) PersonAsResponse(resPerson interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PersonAsResponse", reflect.TypeOf((*MockPeopleConverter)(nil).PersonAsResponse), resPerson)
}

// PersonAsRequest mocks base method
func (m *MockPeopleConverter) PersonAsRequest(reqPerson pb.Person) person.Person {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PersonAsRequest", reqPerson)
	ret0, _ := ret[0].(person.Person)
	return ret0
}

// PersonAsRequest indicates an expected call of PersonAsRequest
func (mr *MockPeopleConverterMockRecorder) PersonAsRequest(reqPerson interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PersonAsRequest", reflect.TypeOf((*MockPeopleConverter)(nil).PersonAsRequest), reqPerson)
}

// PeopleAsResponse mocks base method
func (m *MockPeopleConverter) PeopleAsResponse(reqPeople []person.Person) []*pb.Person {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeopleAsResponse", reqPeople)
	ret0, _ := ret[0].([]*pb.Person)
	return ret0
}

// PeopleAsResponse indicates an expected call of PeopleAsResponse
func (mr *MockPeopleConverterMockRecorder) PeopleAsResponse(reqPeople interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeopleAsResponse", reflect.TypeOf((*MockPeopleConverter)(nil).PeopleAsResponse), reqPeople)
}

// PeopleAsRequest mocks base method
func (m *MockPeopleConverter) PeopleAsRequest(resPeople []pb.Person) []person.Person {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeopleAsRequest", resPeople)
	ret0, _ := ret[0].([]person.Person)
	return ret0
}

// PeopleAsRequest indicates an expected call of PeopleAsRequest
func (mr *MockPeopleConverterMockRecorder) PeopleAsRequest(resPeople interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeopleAsRequest", reflect.TypeOf((*MockPeopleConverter)(nil).PeopleAsRequest), resPeople)
}