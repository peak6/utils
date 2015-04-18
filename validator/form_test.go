package validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRequiredString(t *testing.T) {
	assert := assert.New(t)
	v := NewForm()
	defer v.Close()
	v.RequiredString(``, "name")
	assert.True(v.HasError())

	v.Clear()
	v.RequiredString("aaa", "name")
	assert.False(v.HasError())
}

func TestRequiredInt(t *testing.T) {
	assert := assert.New(t)
	v := NewForm()
	defer v.Close()
	v.RequiredInt(0, "name")
	assert.True(v.HasError())

	v.Clear()
	v.RequiredInt(10, "name")
	assert.False(v.HasError())
}

func TestRequiredBool(t *testing.T) {
	assert := assert.New(t)
	v := NewForm()
	defer v.Close()
	v.RequiredBool(false, "name")
	assert.True(v.HasError())

	v.Clear()
	v.RequiredBool(true, "name")
	assert.False(v.HasError())
}

func TestRequiredTime(t *testing.T) {
	assert := assert.New(t)
	v := NewForm()
	defer v.Close()
	v.RequiredTime(time.Time{}, "name")
	assert.True(v.HasError())

	v.Clear()
	now := time.Now()
	v.RequiredTime(now, "name")
	assert.False(v.HasError())
}

func TestNotNil(t *testing.T) {
	assert := assert.New(t)
	v := NewForm()
	defer v.Close()
	v.NotNil(nil, "name")
	assert.True(v.HasError())

	v.Clear()
	v.NotNil(struct{}{}, "name")
	assert.False(v.HasError())
}

func TestMinChar(t *testing.T) {
	assert := assert.New(t)
	v := NewForm()
	defer v.Close()
	v.MinChar("ssss", 10, "name")
	assert.True(v.HasError())

	v.Clear()
	v.MinChar("ssss", 1, "name")
	assert.False(v.HasError())
}

func TestEmail(t *testing.T) {
	assert := assert.New(t)
	v := NewForm()
	defer v.Close()
	v.Email("rerer@ddsds", "name")
	assert.True(v.HasError())

	v.Clear()
	v.NotNil("test@test.com", "name")
	assert.False(v.HasError())
}

func TestGender(t *testing.T) {
	assert := assert.New(t)
	v := NewForm()
	defer v.Close()
	v.Gender("ddd", "name")
	assert.True(v.HasError())

	v.Clear()
	v.Gender("male", "name")
	assert.False(v.HasError())
}

func TestConfirm(t *testing.T) {
	assert := assert.New(t)
	v := NewForm()
	defer v.Close()
	v.Gender("ddd", "name")
	assert.True(v.HasError())

	v.Clear()
	v.Gender("male", "name")
	assert.False(v.HasError())
}

func TestISO8601DataTime(t *testing.T) {
	assert := assert.New(t)
	v := NewForm()
	defer v.Close()
	v.ISO8601DataTime("2014-04-18", "name")
	assert.True(v.HasError())

	v.Clear()
	v.ISO8601DataTime("2014-04-18T10:47:23+07:00", "name")
	assert.False(v.HasError())

	v.Clear()
	v.ISO8601DataTime("2014-04-18T10:47:23Z", "name")
	assert.False(v.HasError())
}

func TestInString(t *testing.T) {
	assert := assert.New(t)
	v := NewForm()
	defer v.Close()
	v.InString("ddd", []string{"a", "b", "c"}, "name")
	assert.True(v.HasError())

	v.Clear()
	v.InString("c", []string{"a", "b", "c"}, "name")
	assert.False(v.HasError())
}

func TestErrorMessages(t *testing.T) {
	assert := assert.New(t)
	v := NewForm()
	defer v.Close()
	v.RequiredString(``, "name")
	v.RequiredString(``, "jack")
	v.Email("tttt", "email")
	assert.True(v.HasError())
	assert.Equal(map[string]string{
		"email": "email is not invalid email",
		"name":  "name is required",
		"jack":  "jack is required",
	}, v.Messages())
}

func TestAddError(t *testing.T) {
	assert := assert.New(t)
	v := NewForm()
	defer v.Close()
	v.AddError("name", "name is error")
	assert.True(v.HasError())
	assert.Equal("name is error", v.Message())
}
