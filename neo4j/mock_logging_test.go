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
// Source: github.com/neo4j/neo4j-go-driver (interfaces: Logging)

// Package mocks is a generated GoMock package.
package neo4j

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLogging is a mock of Logging interface
type MockLogging struct {
	ctrl     *gomock.Controller
	recorder *MockLoggingMockRecorder
}

// MockLoggingMockRecorder is the mock recorder for MockLogging
type MockLoggingMockRecorder struct {
	mock *MockLogging
}

// NewMockLogging creates a new mock instance
func NewMockLogging(ctrl *gomock.Controller) *MockLogging {
	mock := &MockLogging{ctrl: ctrl}
	mock.recorder = &MockLoggingMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLogging) EXPECT() *MockLoggingMockRecorder {
	return m.recorder
}

// DebugEnabled mocks base method
func (m *MockLogging) DebugEnabled() bool {
	ret := m.ctrl.Call(m, "DebugEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// DebugEnabled indicates an expected call of DebugEnabled
func (mr *MockLoggingMockRecorder) DebugEnabled() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DebugEnabled", reflect.TypeOf((*MockLogging)(nil).DebugEnabled))
}

// Debugf mocks base method
func (m *MockLogging) Debugf(arg0 string, arg1 ...interface{}) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf
func (mr *MockLoggingMockRecorder) Debugf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockLogging)(nil).Debugf), varargs...)
}

// ErrorEnabled mocks base method
func (m *MockLogging) ErrorEnabled() bool {
	ret := m.ctrl.Call(m, "ErrorEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// ErrorEnabled indicates an expected call of ErrorEnabled
func (mr *MockLoggingMockRecorder) ErrorEnabled() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ErrorEnabled", reflect.TypeOf((*MockLogging)(nil).ErrorEnabled))
}

// Errorf mocks base method
func (m *MockLogging) Errorf(arg0 string, arg1 ...interface{}) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf
func (mr *MockLoggingMockRecorder) Errorf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockLogging)(nil).Errorf), varargs...)
}

// InfoEnabled mocks base method
func (m *MockLogging) InfoEnabled() bool {
	ret := m.ctrl.Call(m, "InfoEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// InfoEnabled indicates an expected call of InfoEnabled
func (mr *MockLoggingMockRecorder) InfoEnabled() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InfoEnabled", reflect.TypeOf((*MockLogging)(nil).InfoEnabled))
}

// Infof mocks base method
func (m *MockLogging) Infof(arg0 string, arg1 ...interface{}) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infof", varargs...)
}

// Infof indicates an expected call of Infof
func (mr *MockLoggingMockRecorder) Infof(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infof", reflect.TypeOf((*MockLogging)(nil).Infof), varargs...)
}

// WarningEnabled mocks base method
func (m *MockLogging) WarningEnabled() bool {
	ret := m.ctrl.Call(m, "WarningEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// WarningEnabled indicates an expected call of WarningEnabled
func (mr *MockLoggingMockRecorder) WarningEnabled() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WarningEnabled", reflect.TypeOf((*MockLogging)(nil).WarningEnabled))
}

// Warningf mocks base method
func (m *MockLogging) Warningf(arg0 string, arg1 ...interface{}) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warningf", varargs...)
}

// Warningf indicates an expected call of Warningf
func (mr *MockLoggingMockRecorder) Warningf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warningf", reflect.TypeOf((*MockLogging)(nil).Warningf), varargs...)
}
