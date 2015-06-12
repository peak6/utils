package validator

import (
	"encoding/json"
	"time"

	"github.com/plimble/utils/errors2"
	"github.com/plimble/utils/strings2"
)

//go:generate mockery -name Former

const (
	requiredMsg = ` is required`
	emailMsg    = ` is not invalid email`
	minCharMsg  = ` is too short`
	maxCharMsg  = ` is too long`
	genderMsg   = ` is not male or female`
	inMsg       = ` is not in list`
	dateMsg     = ` is not iso8601 date format`
	emptyMsg    = ` is empty`
	confirmMsg  = ` doesn't match`
)

type Former interface {
	Close()
	Clear()
	HasError() bool
	AddError(field, err error)
	Messages() map[string]string
	Message() string
	Error() error
	RequiredString(val string, field string, err error)
	RequiredInt(val int, field string, err error)
	RequiredFloat64(val float64, field string, err error)
	RequiredBool(val bool, field string, err error)
	RequiredEmail(val string, field string, err error)
	NotNil(val interface{}, field string, err error)
	RequiredTime(val time.Time, field string, err error)
	MinInt(val int, n int, field string, err error)
	MaxInt(val int, n int, field string, err error)
	MaxFloat64(val float64, n float64, field string, err error)
	MinFloat64(val float64, n float64, field string, err error)
	MinChar(val string, n int, field string, err error)
	MaxChar(val string, n int, field string, err error)
	Email(val, field string, err error)
	Gender(val, field string, err error)
	Confirm(val, confirm, field string, err error)
	ISO8601DataTime(val, field string, err error)
	InString(val string, in []string, field string, err error)
	Length(val int, atleast int, field string, err error)
}

type Form struct {
	messages map[string]string
	isError  bool
}

func NewForm() *Form {
	return formPool.Get().(*Form)
}

func (f *Form) setMsg(field, defaultMsg string, err error) {
	if f.messages == nil {
		f.messages = make(map[string]string)
	}

	if f.messages[field] != "" {
		return
	}

	f.messages[field] = err.Error()
}

func (f *Form) Clear() {
	for k := range f.messages {
		delete(f.messages, k)
	}
	f.isError = false
}

func (f *Form) HasError() bool {
	return f.isError
}

func (f *Form) AddError(field string, err error) {
	if f.messages == nil {
		f.messages = make(map[string]string)
	}
	f.isError = true
	f.messages[field] = err.Error()
}

func (f *Form) Messages() map[string]string {
	return f.messages
}

func (f *Form) Message() string {
	for _, msg := range f.messages {
		return msg
	}

	return ""
}

func (f *Form) Json() ([]byte, error) {
	if f.isError {
		return json.Marshal(f.messages)
	}

	return nil, nil
}

func (f *Form) GetError() error {
	if f.isError {
		data, _ := f.Json()
		return errors2.NewBadReq(string(data))
	}

	return nil
}

func (f *Form) Close() {
	f.messages = nil
	f.isError = false
	formPool.Put(f)
}

func (f *Form) Error() error {
	data, _ := json.Marshal(f.messages)
	return errors2.NewBadReq(string2.BytesToString(data))
}

func (f *Form) RequiredString(val string, field string, err error) {
	if requiredString(val) {
		f.isError = true
		f.setMsg(field, requiredMsg, err)
	}
}

func (f *Form) RequiredInt(val int, field string, err error) {
	if requiredInt(val) {
		f.isError = true
		f.setMsg(field, requiredMsg, err)
	}
}

func (f *Form) RequiredFloat64(val float64, field string, err error) {
	if requiredFloat64(val) {
		f.isError = true
		f.setMsg(field, requiredMsg, err)
	}
}

func (f *Form) RequiredBool(val bool, field string, err error) {
	if requiredBool(val) {
		f.isError = true
		f.setMsg(field, requiredMsg, err)
	}
}

func (f *Form) RequiredEmail(val string, field string, err error) {
	if requireEmail(val) {
		f.isError = true
		f.setMsg(field, requiredMsg, err)
	}
}

func (f *Form) NotNil(val interface{}, field string, err error) {
	if notNil(val) {
		f.isError = true
		f.setMsg(field, requiredMsg, err)
	}
}

func (f *Form) RequiredTime(val time.Time, field string, err error) {
	if requiredTime(val) {
		f.isError = true
		f.setMsg(field, requiredMsg, err)
	}
}

func (f *Form) MinInt(val int, n int, field string, err error) {
	if minInt(val, n) {
		f.isError = true
		f.setMsg(field, minCharMsg, err)
	}
}

func (f *Form) MaxInt(val int, n int, field string, err error) {
	if maxInt(val, n) {
		f.isError = true
		f.setMsg(field, maxCharMsg, err)
	}
}

func (f *Form) MaxFloat64(val float64, n float64, field string, err error) {
	if maxFloat64(val, n) {
		f.isError = true
		f.setMsg(field, maxCharMsg, err)
	}
}

func (f *Form) MinFloat64(val float64, n float64, field string, err error) {
	if minFloat64(val, n) {
		f.isError = true
		f.setMsg(field, minCharMsg, err)
	}
}

func (f *Form) MinChar(val string, n int, field string, err error) {
	if minChar(val, n) {
		f.isError = true
		f.setMsg(field, minCharMsg, err)
	}
}

func (f *Form) MaxChar(val string, n int, field string, err error) {
	if maxChar(val, n) {
		f.isError = true
		f.setMsg(field, maxCharMsg, err)
	}
}

func (f *Form) Email(val, field string, err error) {
	if email(val) {
		f.isError = true
		f.setMsg(field, emailMsg, err)
	}
}

func (f *Form) Gender(val, field string, err error) {
	if gender(val) {
		f.isError = true
		f.setMsg(field, genderMsg, err)
	}
}

func (f *Form) Confirm(val, confirm, field string, err error) {
	if confirmVals(val, confirm) {
		f.isError = true
		f.setMsg(field, confirm, err)
	}
}

func (f *Form) ISO8601DataTime(val, field string, err error) {
	if iso8601DataTime(val) {
		f.isError = true
		f.setMsg(field, dateMsg, err)
	}
}

func (f *Form) InString(val string, in []string, field string, err error) {
	if inString(val, in) {
		f.isError = true
		f.setMsg(field, inMsg, err)
	}
}

func (f *Form) Length(val int, atleast int, field string, err error) {
	if length(val, atleast) {
		f.isError = true
		f.setMsg(field, inMsg, err)
	}
}
