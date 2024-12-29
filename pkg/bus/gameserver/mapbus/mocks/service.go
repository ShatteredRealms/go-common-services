// Code generated by MockGen. DO NOT EDIT.
// Source: /home/wil/sro/go-common-service/pkg/bus/gameserver/mapbus/service.go
//
// Generated by this command:
//
//	mockgen -source=/home/wil/sro/go-common-service/pkg/bus/gameserver/mapbus/service.go -destination=/home/wil/sro/go-common-service/pkg/bus/gameserver/mapbus/mocks/service.go
//

// Package mock_mapbus is a generated GoMock package.
package mock_mapbus

import (
	context "context"
	reflect "reflect"

	bus "github.com/ShatteredRealms/go-common-service/pkg/bus"
	mapbus "github.com/ShatteredRealms/go-common-service/pkg/bus/gameserver/mapbus"
	gomock "go.uber.org/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
	isgomock struct{}
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// GetMapById mocks base method.
func (m *MockService) GetMapById(ctx context.Context, mapId string) (*mapbus.Map, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMapById", ctx, mapId)
	ret0, _ := ret[0].(*mapbus.Map)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMapById indicates an expected call of GetMapById.
func (mr *MockServiceMockRecorder) GetMapById(ctx, mapId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMapById", reflect.TypeOf((*MockService)(nil).GetMapById), ctx, mapId)
}

// GetMaps mocks base method.
func (m *MockService) GetMaps(ctx context.Context) (*mapbus.Maps, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMaps", ctx)
	ret0, _ := ret[0].(*mapbus.Maps)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMaps indicates an expected call of GetMaps.
func (mr *MockServiceMockRecorder) GetMaps(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaps", reflect.TypeOf((*MockService)(nil).GetMaps), ctx)
}

// GetResetter mocks base method.
func (m *MockService) GetResetter() bus.Resettable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResetter")
	ret0, _ := ret[0].(bus.Resettable)
	return ret0
}

// GetResetter indicates an expected call of GetResetter.
func (mr *MockServiceMockRecorder) GetResetter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResetter", reflect.TypeOf((*MockService)(nil).GetResetter))
}

// IsProcessing mocks base method.
func (m *MockService) IsProcessing() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsProcessing")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsProcessing indicates an expected call of IsProcessing.
func (mr *MockServiceMockRecorder) IsProcessing() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsProcessing", reflect.TypeOf((*MockService)(nil).IsProcessing))
}

// StartProcessing mocks base method.
func (m *MockService) StartProcessing(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartProcessing", ctx)
}

// StartProcessing indicates an expected call of StartProcessing.
func (mr *MockServiceMockRecorder) StartProcessing(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartProcessing", reflect.TypeOf((*MockService)(nil).StartProcessing), ctx)
}

// StopProcessing mocks base method.
func (m *MockService) StopProcessing() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StopProcessing")
}

// StopProcessing indicates an expected call of StopProcessing.
func (mr *MockServiceMockRecorder) StopProcessing() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopProcessing", reflect.TypeOf((*MockService)(nil).StopProcessing))
}
