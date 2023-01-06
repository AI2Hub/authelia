// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/authelia/authelia/v4/internal/trust (interfaces: Provider)

// Package mocks is a generated GoMock package.
package mocks

import (
	tls "crypto/tls"
	x509 "crypto/x509"
	reflect "reflect"

	schema "github.com/authelia/authelia/v4/internal/configuration/schema"
	gomock "github.com/golang/mock/gomock"
)

// MockTrust is a mock of Provider interface.
type MockTrust struct {
	ctrl     *gomock.Controller
	recorder *MockTrustMockRecorder
}

// MockTrustMockRecorder is the mock recorder for MockTrust.
type MockTrustMockRecorder struct {
	mock *MockTrust
}

// NewMockTrust creates a new mock instance.
func NewMockTrust(ctrl *gomock.Controller) *MockTrust {
	mock := &MockTrust{ctrl: ctrl}
	mock.recorder = &MockTrustMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrust) EXPECT() *MockTrustMockRecorder {
	return m.recorder
}

// AddTrustedCertificate mocks base method.
func (m *MockTrust) AddTrustedCertificate(arg0 *x509.Certificate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTrustedCertificate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTrustedCertificate indicates an expected call of AddTrustedCertificate.
func (mr *MockTrustMockRecorder) AddTrustedCertificate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTrustedCertificate", reflect.TypeOf((*MockTrust)(nil).AddTrustedCertificate), arg0)
}

// AddTrustedCertificateFromPath mocks base method.
func (m *MockTrust) AddTrustedCertificateFromPath(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTrustedCertificateFromPath", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTrustedCertificateFromPath indicates an expected call of AddTrustedCertificateFromPath.
func (mr *MockTrustMockRecorder) AddTrustedCertificateFromPath(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTrustedCertificateFromPath", reflect.TypeOf((*MockTrust)(nil).AddTrustedCertificateFromPath), arg0)
}

// GetTLSConfiguration mocks base method.
func (m *MockTrust) GetTLSConfiguration(arg0 *schema.TLSConfig) *tls.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTLSConfiguration", arg0)
	ret0, _ := ret[0].(*tls.Config)
	return ret0
}

// GetTLSConfiguration indicates an expected call of GetTLSConfiguration.
func (mr *MockTrustMockRecorder) GetTLSConfiguration(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTLSConfiguration", reflect.TypeOf((*MockTrust)(nil).GetTLSConfiguration), arg0)
}

// GetTrustedCertificates mocks base method.
func (m *MockTrust) GetTrustedCertificates() *x509.CertPool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTrustedCertificates")
	ret0, _ := ret[0].(*x509.CertPool)
	return ret0
}

// GetTrustedCertificates indicates an expected call of GetTrustedCertificates.
func (mr *MockTrustMockRecorder) GetTrustedCertificates() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrustedCertificates", reflect.TypeOf((*MockTrust)(nil).GetTrustedCertificates))
}

// StartupCheck mocks base method.
func (m *MockTrust) StartupCheck() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartupCheck")
	ret0, _ := ret[0].(error)
	return ret0
}

// StartupCheck indicates an expected call of StartupCheck.
func (mr *MockTrustMockRecorder) StartupCheck() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartupCheck", reflect.TypeOf((*MockTrust)(nil).StartupCheck))
}
