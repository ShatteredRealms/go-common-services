// Code generated by MockGen. DO NOT EDIT.
// Source: /home/wil/dev/sro/go-common-service/pkg/pb/health_grpc.pb.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=/home/wil/dev/sro/go-common-service/pkg/pb/health_grpc.pb.go -destination=/home/wil/dev/sro/go-common-service/pkg/mocks/health_grpc.pb_mock.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	pb "github.com/ShatteredRealms/go-common-service/pkg/pb"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockHealthServiceClient is a mock of HealthServiceClient interface.
type MockHealthServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockHealthServiceClientMockRecorder
	isgomock struct{}
}

// MockHealthServiceClientMockRecorder is the mock recorder for MockHealthServiceClient.
type MockHealthServiceClientMockRecorder struct {
	mock *MockHealthServiceClient
}

// NewMockHealthServiceClient creates a new mock instance.
func NewMockHealthServiceClient(ctrl *gomock.Controller) *MockHealthServiceClient {
	mock := &MockHealthServiceClient{ctrl: ctrl}
	mock.recorder = &MockHealthServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHealthServiceClient) EXPECT() *MockHealthServiceClientMockRecorder {
	return m.recorder
}

// Health mocks base method.
func (m *MockHealthServiceClient) Health(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*pb.HealthMessage, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Health", varargs...)
	ret0, _ := ret[0].(*pb.HealthMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Health indicates an expected call of Health.
func (mr *MockHealthServiceClientMockRecorder) Health(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Health", reflect.TypeOf((*MockHealthServiceClient)(nil).Health), varargs...)
}

// MockHealthServiceServer is a mock of HealthServiceServer interface.
type MockHealthServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockHealthServiceServerMockRecorder
	isgomock struct{}
}

// MockHealthServiceServerMockRecorder is the mock recorder for MockHealthServiceServer.
type MockHealthServiceServerMockRecorder struct {
	mock *MockHealthServiceServer
}

// NewMockHealthServiceServer creates a new mock instance.
func NewMockHealthServiceServer(ctrl *gomock.Controller) *MockHealthServiceServer {
	mock := &MockHealthServiceServer{ctrl: ctrl}
	mock.recorder = &MockHealthServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHealthServiceServer) EXPECT() *MockHealthServiceServerMockRecorder {
	return m.recorder
}

// Health mocks base method.
func (m *MockHealthServiceServer) Health(arg0 context.Context, arg1 *emptypb.Empty) (*pb.HealthMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Health", arg0, arg1)
	ret0, _ := ret[0].(*pb.HealthMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Health indicates an expected call of Health.
func (mr *MockHealthServiceServerMockRecorder) Health(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Health", reflect.TypeOf((*MockHealthServiceServer)(nil).Health), arg0, arg1)
}

// mustEmbedUnimplementedHealthServiceServer mocks base method.
func (m *MockHealthServiceServer) mustEmbedUnimplementedHealthServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedHealthServiceServer")
}

// mustEmbedUnimplementedHealthServiceServer indicates an expected call of mustEmbedUnimplementedHealthServiceServer.
func (mr *MockHealthServiceServerMockRecorder) mustEmbedUnimplementedHealthServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedHealthServiceServer", reflect.TypeOf((*MockHealthServiceServer)(nil).mustEmbedUnimplementedHealthServiceServer))
}

// MockUnsafeHealthServiceServer is a mock of UnsafeHealthServiceServer interface.
type MockUnsafeHealthServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeHealthServiceServerMockRecorder
	isgomock struct{}
}

// MockUnsafeHealthServiceServerMockRecorder is the mock recorder for MockUnsafeHealthServiceServer.
type MockUnsafeHealthServiceServerMockRecorder struct {
	mock *MockUnsafeHealthServiceServer
}

// NewMockUnsafeHealthServiceServer creates a new mock instance.
func NewMockUnsafeHealthServiceServer(ctrl *gomock.Controller) *MockUnsafeHealthServiceServer {
	mock := &MockUnsafeHealthServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeHealthServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeHealthServiceServer) EXPECT() *MockUnsafeHealthServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedHealthServiceServer mocks base method.
func (m *MockUnsafeHealthServiceServer) mustEmbedUnimplementedHealthServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedHealthServiceServer")
}

// mustEmbedUnimplementedHealthServiceServer indicates an expected call of mustEmbedUnimplementedHealthServiceServer.
func (mr *MockUnsafeHealthServiceServerMockRecorder) mustEmbedUnimplementedHealthServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedHealthServiceServer", reflect.TypeOf((*MockUnsafeHealthServiceServer)(nil).mustEmbedUnimplementedHealthServiceServer))
}
