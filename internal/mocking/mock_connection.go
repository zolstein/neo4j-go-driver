/*
 * Copyright (c) 2002-2018 "Neo4j,"
 * Neo4j Sweden AB [http://seabolt.com]
 *
 * This file is part of seabolt.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by MockGen. DO NOT EDIT.
// Source: neo4j-go-connector/neo4j (interfaces: Connection)

// Package connector-mocks is a generated GoMock package.
package mocking

import (
	reflect "reflect"

	"github.com/neo4j-drivers/neo4j-go-connector"

	gomock "github.com/golang/mock/gomock"
)

// MockConnection is a mock of Connection interface
type MockConnection struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionMockRecorder
}

// MockConnectionMockRecorder is the mock recorder for MockConnection
type MockConnectionMockRecorder struct {
	mock *MockConnection
}

// NewMockConnection creates a new mock instance
func NewMockConnection(ctrl *gomock.Controller) *MockConnection {
	mock := &MockConnection{ctrl: ctrl}
	mock.recorder = &MockConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConnection) EXPECT() *MockConnectionMockRecorder {
	return m.recorder
}

// Begin connector-mocks base method
func (m *MockConnection) Begin(arg0 []string) (seabolt.RequestHandle, error) {
	ret := m.ctrl.Call(m, "Begin", arg0)
	ret0, _ := ret[0].(seabolt.RequestHandle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Begin indicates an expected call of Begin
func (mr *MockConnectionMockRecorder) Begin(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockConnection)(nil).Begin), arg0)
}

// Close connector-mocks base method
func (m *MockConnection) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockConnectionMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConnection)(nil).Close))
}

// Commit connector-mocks base method
func (m *MockConnection) Commit() (seabolt.RequestHandle, error) {
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(seabolt.RequestHandle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Commit indicates an expected call of Commit
func (mr *MockConnectionMockRecorder) Commit() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockConnection)(nil).Commit))
}

// Data connector-mocks base method
func (m *MockConnection) Data() ([]interface{}, error) {
	ret := m.ctrl.Call(m, "Data")
	ret0, _ := ret[0].([]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Data indicates an expected call of Data
func (mr *MockConnectionMockRecorder) Data() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Data", reflect.TypeOf((*MockConnection)(nil).Data))
}

// DiscardAll connector-mocks base method
func (m *MockConnection) DiscardAll() (seabolt.RequestHandle, error) {
	ret := m.ctrl.Call(m, "DiscardAll")
	ret0, _ := ret[0].(seabolt.RequestHandle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DiscardAll indicates an expected call of DiscardAll
func (mr *MockConnectionMockRecorder) DiscardAll() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiscardAll", reflect.TypeOf((*MockConnection)(nil).DiscardAll))
}

// Fetch connector-mocks base method
func (m *MockConnection) Fetch(arg0 seabolt.RequestHandle) (seabolt.FetchType, error) {
	ret := m.ctrl.Call(m, "Fetch", arg0)
	ret0, _ := ret[0].(seabolt.FetchType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch
func (mr *MockConnectionMockRecorder) Fetch(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockConnection)(nil).Fetch), arg0)
}

// FetchSummary connector-mocks base method
func (m *MockConnection) FetchSummary(arg0 seabolt.RequestHandle) (int, error) {
	ret := m.ctrl.Call(m, "FetchSummary", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchSummary indicates an expected call of FetchSummary
func (mr *MockConnectionMockRecorder) FetchSummary(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchSummary", reflect.TypeOf((*MockConnection)(nil).FetchSummary), arg0)
}

// Flush connector-mocks base method
func (m *MockConnection) Flush() error {
	ret := m.ctrl.Call(m, "Flush")
	ret0, _ := ret[0].(error)
	return ret0
}

// Flush indicates an expected call of Flush
func (mr *MockConnectionMockRecorder) Flush() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockConnection)(nil).Flush))
}

// LastBookmark connector-mocks base method
func (m *MockConnection) LastBookmark() string {
	ret := m.ctrl.Call(m, "LastBookmark")
	ret0, _ := ret[0].(string)
	return ret0
}

// LastBookmark indicates an expected call of LastBookmark
func (mr *MockConnectionMockRecorder) LastBookmark() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastBookmark", reflect.TypeOf((*MockConnection)(nil).LastBookmark))
}

// Metadata connector-mocks base method
func (m *MockConnection) Metadata() (map[string]interface{}, error) {
	ret := m.ctrl.Call(m, "Metadata")
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Metadata indicates an expected call of Metadata
func (mr *MockConnectionMockRecorder) Metadata() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Metadata", reflect.TypeOf((*MockConnection)(nil).Metadata))
}

// PullAll connector-mocks base method
func (m *MockConnection) PullAll() (seabolt.RequestHandle, error) {
	ret := m.ctrl.Call(m, "PullAll")
	ret0, _ := ret[0].(seabolt.RequestHandle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PullAll indicates an expected call of PullAll
func (mr *MockConnectionMockRecorder) PullAll() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PullAll", reflect.TypeOf((*MockConnection)(nil).PullAll))
}

// Reset connector-mocks base method
func (m *MockConnection) Reset() error {
	ret := m.ctrl.Call(m, "Reset")
	ret0, _ := ret[0].(error)
	return ret0
}

// Reset indicates an expected call of Reset
func (mr *MockConnectionMockRecorder) Reset() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockConnection)(nil).Reset))
}

// Rollback connector-mocks base method
func (m *MockConnection) Rollback() (seabolt.RequestHandle, error) {
	ret := m.ctrl.Call(m, "Rollback")
	ret0, _ := ret[0].(seabolt.RequestHandle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Rollback indicates an expected call of Rollback
func (mr *MockConnectionMockRecorder) Rollback() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockConnection)(nil).Rollback))
}

// Run connector-mocks base method
func (m *MockConnection) Run(arg0 string, arg1 *map[string]interface{}) (seabolt.RequestHandle, error) {
	ret := m.ctrl.Call(m, "Run", arg0, arg1)
	ret0, _ := ret[0].(seabolt.RequestHandle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Run indicates an expected call of Run
func (mr *MockConnectionMockRecorder) Run(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockConnection)(nil).Run), arg0, arg1)
}