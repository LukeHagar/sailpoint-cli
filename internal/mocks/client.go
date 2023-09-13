// Code generated by MockGen. DO NOT EDIT.
// Source: internal/client/client.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	io "io"
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockClient) Delete(ctx context.Context, url string, params map[string]string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, url, params)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockClientMockRecorder) Delete(ctx, url, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClient)(nil).Delete), ctx, url, params)
}

// Get mocks base method.
func (m *MockClient) Get(ctx context.Context, url string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, url)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockClientMockRecorder) Get(ctx, url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClient)(nil).Get), ctx, url)
}

// Patch mocks base method.
func (m *MockClient) Patch(ctx context.Context, url string, body io.Reader) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Patch", ctx, url, body)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockClientMockRecorder) Patch(ctx, url, body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockClient)(nil).Patch), ctx, url, body)
}

// Post mocks base method.
func (m *MockClient) Post(ctx context.Context, url, contentType string, body io.Reader) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Post", ctx, url, contentType, body)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Post indicates an expected call of Post.
func (mr *MockClientMockRecorder) Post(ctx, url, contentType, body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockClient)(nil).Post), ctx, url, contentType, body)
}

// Put mocks base method.
func (m *MockClient) Put(ctx context.Context, url, contentType string, body io.Reader) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", ctx, url, contentType, body)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Put indicates an expected call of Put.
func (mr *MockClientMockRecorder) Put(ctx, url, contentType, body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockClient)(nil).Put), ctx, url, contentType, body)
}
