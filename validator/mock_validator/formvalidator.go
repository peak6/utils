package mock_validator

import "github.com/stretchr/testify/mock"

import "time"

type MockFormValidator struct {
	mock.Mock
}

func NewMockFormValidator() *MockFormValidator {
	return &MockFormValidator{}
}

func (m *MockFormValidator) Clear() {
	m.Called()
}
func (m *MockFormValidator) HasError() bool {
	ret := m.Called()

	r0 := ret.Get(0).(bool)

	return r0
}
func (m *MockFormValidator) AddError(field string, msg string) {
	m.Called(field, msg)
}
func (m *MockFormValidator) Messages() map[string]string {
	ret := m.Called()

	var r0 map[string]string
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(map[string]string)
	}

	return r0
}
func (m *MockFormValidator) Message() string {
	ret := m.Called()

	r0 := ret.Get(0).(string)

	return r0
}
func (m *MockFormValidator) Error() error {
	ret := m.Called()

	r0 := ret.Error(0)

	return r0
}
func (m *MockFormValidator) RequiredString(val string, field string, msg ...string) {
	m.Called(val, field, msg)
}
func (m *MockFormValidator) RequiredInt(val int, field string, msg ...string) {
	m.Called(val, field, msg)
}
func (m *MockFormValidator) RequiredFloat64(val float64, field string, msg ...string) {
	m.Called(val, field, msg)
}
func (m *MockFormValidator) RequiredBool(val bool, field string, msg ...string) {
	m.Called(val, field, msg)
}
func (m *MockFormValidator) RequiredEmail(val string, field string, msg ...string) {
	m.Called(val, field, msg)
}
func (m *MockFormValidator) NotNil(val interface{}, field string, msg ...string) {
	m.Called(val, field, msg)
}
func (m *MockFormValidator) RequiredTime(val time.Time, field string, msg ...string) {
	m.Called(val, field, msg)
}
func (m *MockFormValidator) MinInt(val int, n int, field string, msg ...string) {
	m.Called(val, n, field, msg)
}
func (m *MockFormValidator) MaxInt(val int, n int, field string, msg ...string) {
	m.Called(val, n, field, msg)
}
func (m *MockFormValidator) MaxFloat64(val float64, n float64, field string, msg ...string) {
	m.Called(val, n, field, msg)
}
func (m *MockFormValidator) MinFloat64(val float64, n float64, field string, msg ...string) {
	m.Called(val, n, field, msg)
}
func (m *MockFormValidator) MinChar(val string, n int, field string, msg ...string) {
	m.Called(val, n, field, msg)
}
func (m *MockFormValidator) MaxChar(val string, n int, field string, msg ...string) {
	m.Called(val, n, field, msg)
}
func (m *MockFormValidator) Email(val string, field string, msg ...string) {
	m.Called(val, field, msg)
}
func (m *MockFormValidator) Gender(val string, field string, msg ...string) {
	m.Called(val, field, msg)
}
func (m *MockFormValidator) Confirm(val string, confirm string, field string, msg ...string) {
	m.Called(val, confirm, field, msg)
}
func (m *MockFormValidator) ISO8601DataTime(val string, field string, msg ...string) {
	m.Called(val, field, msg)
}
func (m *MockFormValidator) InString(val string, in []string, field string, msg ...string) {
	m.Called(val, in, field, msg)
}
