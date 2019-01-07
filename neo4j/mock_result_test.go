/*
 * Copyright (c) 2002-2019 "Neo4j,"
 * Neo4j Sweden AB [http://neo4j.com]
 *
 * This file is part of Neo4j.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/neo4j/neo4j-go-driver/neo4j (interfaces: Result)

// Package mock_neo4j is a generated GoMock package.
package neo4j

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockResult is a mock of Result interface
type MockResult struct {
	ctrl     *gomock.Controller
	recorder *MockResultMockRecorder
}

// MockResultMockRecorder is the mock recorder for MockResult
type MockResultMockRecorder struct {
	mock *MockResult
}

// NewMockResult creates a new mock instance
func NewMockResult(ctrl *gomock.Controller) *MockResult {
	mock := &MockResult{ctrl: ctrl}
	mock.recorder = &MockResultMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockResult) EXPECT() *MockResultMockRecorder {
	return m.recorder
}

// Consume mocks base method
func (m *MockResult) Consume() (ResultSummary, error) {
	ret := m.ctrl.Call(m, "Consume")
	ret0, _ := ret[0].(ResultSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Consume indicates an expected call of Consume
func (mr *MockResultMockRecorder) Consume() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consume", reflect.TypeOf((*MockResult)(nil).Consume))
}

// Err mocks base method
func (m *MockResult) Err() error {
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err
func (mr *MockResultMockRecorder) Err() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockResult)(nil).Err))
}

// Keys mocks base method
func (m *MockResult) Keys() ([]string, error) {
	ret := m.ctrl.Call(m, "Keys")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Keys indicates an expected call of Keys
func (mr *MockResultMockRecorder) Keys() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockResult)(nil).Keys))
}

// Next mocks base method
func (m *MockResult) Next() bool {
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next
func (mr *MockResultMockRecorder) Next() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockResult)(nil).Next))
}

// Record mocks base method
func (m *MockResult) Record() Record {
	ret := m.ctrl.Call(m, "Record")
	ret0, _ := ret[0].(Record)
	return ret0
}

// Record indicates an expected call of Record
func (mr *MockResultMockRecorder) Record() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Record", reflect.TypeOf((*MockResult)(nil).Record))
}

// Summary mocks base method
func (m *MockResult) Summary() (ResultSummary, error) {
	ret := m.ctrl.Call(m, "Summary")
	ret0, _ := ret[0].(ResultSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Summary indicates an expected call of Summary
func (mr *MockResultMockRecorder) Summary() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Summary", reflect.TypeOf((*MockResult)(nil).Summary))
}
