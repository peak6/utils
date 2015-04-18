package mock_validator

import "github.com/stretchr/testify/mock"

import "time"

type MockCheckValidator struct {
	mock.Mock
}

func NewMockCheckValidator() *MockCheckValidator {
	return &MockCheckValidator{}
}

func (m *MockCheckValidator) Close() {
	m.Called()
}
func (m *MockCheckValidator) GetError() error {
	ret := m.Called()

	r0 := ret.Error(0)

	return r0
}
func (m *MockCheckValidator) Clear() {
	m.Called()
}
func (m *MockCheckValidator) CheckErr(err error) {
	m.Called(err)
}
func (m *MockCheckValidator) RequiredString(val string, err error) {
	m.Called(val, err)
}
func (m *MockCheckValidator) RequiredInt(val int, err error) {
	m.Called(val, err)
}
func (m *MockCheckValidator) RequiredFloat64(val float64, err error) {
	m.Called(val, err)
}
func (m *MockCheckValidator) RequiredBool(val bool, err error) {
	m.Called(val, err)
}
func (m *MockCheckValidator) RequiredEmail(val string, err error) {
	m.Called(val, err)
}
func (m *MockCheckValidator) NotNil(val interface{}, err error) {
	m.Called(val, err)
}
func (m *MockCheckValidator) RequiredTime(val time.Time, err error) {
	m.Called(val, err)
}
func (m *MockCheckValidator) MinInt(val int, n int, err error) {
	m.Called(val, n, err)
}
func (m *MockCheckValidator) MaxInt(val int, n int, err error) {
	m.Called(val, n, err)
}
func (m *MockCheckValidator) MinFloat64(val float64, n float64, err error) {
	m.Called(val, n, err)
}
func (m *MockCheckValidator) MaxFloat64(val float64, n float64, err error) {
	m.Called(val, n, err)
}
func (m *MockCheckValidator) MinChar(val string, n int, err error) {
	m.Called(val, n, err)
}
func (m *MockCheckValidator) MaxChar(val string, n int, err error) {
	m.Called(val, n, err)
}
func (m *MockCheckValidator) Email(val string, err error) {
	m.Called(val, err)
}
func (m *MockCheckValidator) Gender(val string, err error) {
	m.Called(val, err)
}
func (m *MockCheckValidator) Confirm(val string, confirm string, err error) {
	m.Called(val, confirm, err)
}
func (m *MockCheckValidator) ISO8601DataTime(val string, err error) {
	m.Called(val, err)
}
func (m *MockCheckValidator) InString(val string, in []string, err error) {
	m.Called(val, in, err)
}
