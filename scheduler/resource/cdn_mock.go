// Code generated by MockGen. DO NOT EDIT.
// Source: cdn.go

// Package resource is a generated GoMock package.
package resource

import (
	context "context"
	reflect "reflect"

	dfnet "d7y.io/dragonfly/v2/internal/dfnet"
	base "d7y.io/dragonfly/v2/pkg/rpc/base"
	cdnsystem "d7y.io/dragonfly/v2/pkg/rpc/cdnsystem"
	client "d7y.io/dragonfly/v2/pkg/rpc/cdnsystem/client"
	scheduler "d7y.io/dragonfly/v2/pkg/rpc/scheduler"
	config "d7y.io/dragonfly/v2/scheduler/config"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockCDN is a mock of CDN interface.
type MockCDN struct {
	ctrl     *gomock.Controller
	recorder *MockCDNMockRecorder
}

// MockCDNMockRecorder is the mock recorder for MockCDN.
type MockCDNMockRecorder struct {
	mock *MockCDN
}

// NewMockCDN creates a new mock instance.
func NewMockCDN(ctrl *gomock.Controller) *MockCDN {
	mock := &MockCDN{ctrl: ctrl}
	mock.recorder = &MockCDNMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCDN) EXPECT() *MockCDNMockRecorder {
	return m.recorder
}

// Client mocks base method.
func (m *MockCDN) Client() CDNClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Client")
	ret0, _ := ret[0].(CDNClient)
	return ret0
}

// Client indicates an expected call of Client.
func (mr *MockCDNMockRecorder) Client() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockCDN)(nil).Client))
}

// TriggerTask mocks base method.
func (m *MockCDN) TriggerTask(arg0 context.Context, arg1 *Task) (*Peer, *scheduler.PeerResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TriggerTask", arg0, arg1)
	ret0, _ := ret[0].(*Peer)
	ret1, _ := ret[1].(*scheduler.PeerResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// TriggerTask indicates an expected call of TriggerTask.
func (mr *MockCDNMockRecorder) TriggerTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TriggerTask", reflect.TypeOf((*MockCDN)(nil).TriggerTask), arg0, arg1)
}

// MockCDNClient is a mock of CDNClient interface.
type MockCDNClient struct {
	ctrl     *gomock.Controller
	recorder *MockCDNClientMockRecorder
}

// MockCDNClientMockRecorder is the mock recorder for MockCDNClient.
type MockCDNClientMockRecorder struct {
	mock *MockCDNClient
}

// NewMockCDNClient creates a new mock instance.
func NewMockCDNClient(ctrl *gomock.Controller) *MockCDNClient {
	mock := &MockCDNClient{ctrl: ctrl}
	mock.recorder = &MockCDNClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCDNClient) EXPECT() *MockCDNClientMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockCDNClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockCDNClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockCDNClient)(nil).Close))
}

// GetPieceTasks mocks base method.
func (m *MockCDNClient) GetPieceTasks(ctx context.Context, addr dfnet.NetAddr, req *base.PieceTaskRequest, opts ...grpc.CallOption) (*base.PiecePacket, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, addr, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPieceTasks", varargs...)
	ret0, _ := ret[0].(*base.PiecePacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPieceTasks indicates an expected call of GetPieceTasks.
func (mr *MockCDNClientMockRecorder) GetPieceTasks(ctx, addr, req interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, addr, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPieceTasks", reflect.TypeOf((*MockCDNClient)(nil).GetPieceTasks), varargs...)
}

// ObtainSeeds mocks base method.
func (m *MockCDNClient) ObtainSeeds(ctx context.Context, sr *cdnsystem.SeedRequest, opts ...grpc.CallOption) (*client.PieceSeedStream, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, sr}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ObtainSeeds", varargs...)
	ret0, _ := ret[0].(*client.PieceSeedStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ObtainSeeds indicates an expected call of ObtainSeeds.
func (mr *MockCDNClientMockRecorder) ObtainSeeds(ctx, sr interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, sr}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObtainSeeds", reflect.TypeOf((*MockCDNClient)(nil).ObtainSeeds), varargs...)
}

// OnNotify mocks base method.
func (m *MockCDNClient) OnNotify(arg0 *config.DynconfigData) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnNotify", arg0)
}

// OnNotify indicates an expected call of OnNotify.
func (mr *MockCDNClientMockRecorder) OnNotify(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnNotify", reflect.TypeOf((*MockCDNClient)(nil).OnNotify), arg0)
}

// SyncPieceTasks mocks base method.
func (m *MockCDNClient) SyncPieceTasks(ctx context.Context, addr dfnet.NetAddr, ptr *base.PieceTaskRequest, opts ...grpc.CallOption) (cdnsystem.Seeder_SyncPieceTasksClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, addr, ptr}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SyncPieceTasks", varargs...)
	ret0, _ := ret[0].(cdnsystem.Seeder_SyncPieceTasksClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncPieceTasks indicates an expected call of SyncPieceTasks.
func (mr *MockCDNClientMockRecorder) SyncPieceTasks(ctx, addr, ptr interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, addr, ptr}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncPieceTasks", reflect.TypeOf((*MockCDNClient)(nil).SyncPieceTasks), varargs...)
}

// UpdateState mocks base method.
func (m *MockCDNClient) UpdateState(addrs []dfnet.NetAddr) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateState", addrs)
}

// UpdateState indicates an expected call of UpdateState.
func (mr *MockCDNClientMockRecorder) UpdateState(addrs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateState", reflect.TypeOf((*MockCDNClient)(nil).UpdateState), addrs)
}
