package mock_validator

import "github.com/stretchr/testify/mock"

import "time"

type MockCheckInterface struct {
	mock.Mock
}

func NewMockCheckInterface() *MockCheckInterface {
	return &MockCheckInterface{}
}

func (m *MockCheckInterface) RequiredString(val string, err error) {
	m.Called(val, err)
}
func (m *MockCheckInterface) RequiredInt(val int, err error) {
	m.Called(val, err)
}
func (m *MockCheckInterface) RequiredFloat64(val float64, err error) {
	m.Called(val, err)
}
func (m *MockCheckInterface) RequiredBool(val bool, err error) {
	m.Called(val, err)
}
func (m *MockCheckInterface) RequiredEmail(val string, err error) {
	m.Called(val, err)
}
func (m *MockCheckInterface) NotNil(val interface{}, err error) {
	m.Called(val, err)
}
func (m *MockCheckInterface) RequiredTime(val time.Time, err error) {
	m.Called(val, err)
}
func (m *MockCheckInterface) MinInt(val int, n int, err error) {
	m.Called(val, n, err)
}
func (m *MockCheckInterface) MaxInt(val int, n int, err error) {
	m.Called(val, n, err)
}
func (m *MockCheckInterface) MinFloat64(val float64, n float64, err error) {
	m.Called(val, n, err)
}
func (m *MockCheckInterface) MaxFloat64(val float64, n float64, err error) {
	m.Called(val, n, err)
}
func (m *MockCheckInterface) MinChar(val string, n int, err error) {
	m.Called(val, n, err)
}
func (m *MockCheckInterface) MaxChar(val string, n int, err error) {
	m.Called(val, n, err)
}
func (m *MockCheckInterface) Email(val string, err error) {
	m.Called(val, err)
}
func (m *MockCheckInterface) Gender(val string, err error) {
	m.Called(val, err)
}
func (m *MockCheckInterface) Confirm(val string, confirm string, err error) {
	m.Called(val, confirm, err)
}
func (m *MockCheckInterface) ISO8601DataTime(val string, err error) {
	m.Called(val, err)
}
func (m *MockCheckInterface) InString(val string, in []string, err error) {
	m.Called(val, in, err)
}
