package mock_validator

import "github.com/stretchr/testify/mock"

import "time"

type MockFormInterface struct {
	mock.Mock
}

func NewMockFormInterface() *MockFormInterface {
	return &MockFormInterface{}
}

func (m *MockFormInterface) RequiredString(val string, field string, msg ...string) {
	m.Called(val, field, msg)
}
func (m *MockFormInterface) RequiredInt(val int, field string, msg ...string) {
	m.Called(val, field, msg)
}
func (m *MockFormInterface) RequiredFloat64(val float64, field string, msg ...string) {
	m.Called(val, field, msg)
}
func (m *MockFormInterface) RequiredBool(val bool, field string, msg ...string) {
	m.Called(val, field, msg)
}
func (m *MockFormInterface) RequiredEmail(val string, field string, msg ...string) {
	m.Called(val, field, msg)
}
func (m *MockFormInterface) NotNil(val interface{}, field string, msg ...string) {
	m.Called(val, field, msg)
}
func (m *MockFormInterface) RequiredTime(val time.Time, field string, msg ...string) {
	m.Called(val, field, msg)
}
func (m *MockFormInterface) MinInt(val int, n int, field string, msg ...string) {
	m.Called(val, n, field, msg)
}
func (m *MockFormInterface) MaxInt(val int, n int, field string, msg ...string) {
	m.Called(val, n, field, msg)
}
func (m *MockFormInterface) MaxFloat64(val float64, n float64, field string, msg ...string) {
	m.Called(val, n, field, msg)
}
func (m *MockFormInterface) MinFloat64(val float64, n float64, field string, msg ...string) {
	m.Called(val, n, field, msg)
}
func (m *MockFormInterface) MinChar(val string, n int, field string, msg ...string) {
	m.Called(val, n, field, msg)
}
func (m *MockFormInterface) MaxChar(val string, n int, field string, msg ...string) {
	m.Called(val, n, field, msg)
}
func (m *MockFormInterface) Email(val string, field string, msg ...string) {
	m.Called(val, field, msg)
}
func (m *MockFormInterface) Gender(val string, field string, msg ...string) {
	m.Called(val, field, msg)
}
func (m *MockFormInterface) Confirm(val string, confirm string, field string, msg ...string) {
	m.Called(val, confirm, field, msg)
}
func (m *MockFormInterface) ISO8601DataTime(val string, field string, msg ...string) {
	m.Called(val, field, msg)
}
func (m *MockFormInterface) InString(val string, in []string, field string, msg ...string) {
	m.Called(val, in, field, msg)
}
