/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: ../client.go

// Package mock_publicips is a generated GoMock package.
package mock_publicips

import (
	context "context"
	reflect "reflect"

	network "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2021-02-01/network"
	azure "github.com/Azure/go-autorest/autorest/azure"
	gomock "github.com/golang/mock/gomock"
	azure0 "sigs.k8s.io/cluster-api-provider-azure/azure"
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

// CreateOrUpdateAsync mocks base method.
func (m *MockClient) CreateOrUpdateAsync(arg0 context.Context, arg1 azure0.ResourceSpecGetter) (azure.FutureAPI, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateAsync", arg0, arg1)
	ret0, _ := ret[0].(azure.FutureAPI)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdateAsync indicates an expected call of CreateOrUpdateAsync.
func (mr *MockClientMockRecorder) CreateOrUpdateAsync(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateAsync", reflect.TypeOf((*MockClient)(nil).CreateOrUpdateAsync), arg0, arg1)
}

// Delete mocks base method.
func (m *MockClient) Delete(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockClientMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClient)(nil).Delete), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockClient) Get(arg0 context.Context, arg1, arg2 string) (network.PublicIPAddress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(network.PublicIPAddress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockClientMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClient)(nil).Get), arg0, arg1, arg2)
}

// IsDone mocks base method.
func (m *MockClient) IsDone(arg0 context.Context, arg1 azure.FutureAPI) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDone", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsDone indicates an expected call of IsDone.
func (mr *MockClientMockRecorder) IsDone(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDone", reflect.TypeOf((*MockClient)(nil).IsDone), arg0, arg1)
}
