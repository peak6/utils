package validator

import (
	"encoding/json"
	"time"

	"github.com/plimble/utils/errors2"
	"github.com/plimble/utils/strings2"
)

//go:generate mockery -name FormValidator

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

type FormValidator interface {
	Clear()
	HasError() bool
	AddError(field, msg string)
	Messages() map[string]string
	Message() string
	Error() error
	RequiredString(val string, field string, msg ...string)
	RequiredInt(val int, field string, msg ...string)
	RequiredFloat64(val float64, field string, msg ...string)
	RequiredBool(val bool, field string, msg ...string)
	RequiredEmail(val string, field string, msg ...string)
	NotNil(val interface{}, field string, msg ...string)
	RequiredTime(val time.Time, field string, msg ...string)
	MinInt(val int, n int, field string, msg ...string)
	MaxInt(val int, n int, field string, msg ...string)
	MaxFloat64(val float64, n float64, field string, msg ...string)
	MinFloat64(val float64, n float64, field string, msg ...string)
	MinChar(val string, n int, field string, msg ...string)
	MaxChar(val string, n int, field string, msg ...string)
	Email(val, field string, msg ...string)
	Gender(val, field string, msg ...string)
	Confirm(val, confirm, field string, msg ...string)
	ISO8601DataTime(val, field string, msg ...string)
	InString(val string, in []string, field string, msg ...string)
}

type Form struct {
	messages map[string]string
	isError  bool
}

func NewForm() *Form {
	return &Form{}
}

func (f *Form) setMsg(field, defaultMsg string, msg []string) {
	if f.messages == nil {
		f.messages = make(map[string]string)
	}

	if f.messages[field] != "" {
		return
	}

	if len(msg) == 0 {
		f.messages[field] = field + defaultMsg
	} else {
		f.messages[field] = msg[0]
	}
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

func (f *Form) AddError(field, msg string) {
	if f.messages == nil {
		f.messages = make(map[string]string)
	}
	f.isError = true
	f.messages[field] = msg
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

func (f *Form) Error() error {
	data, _ := json.Marshal(f.messages)
	return errors2.NewBadReq(string2.BytesToString(data))
}

func (f *Form) RequiredString(val string, field string, msg ...string) {
	if requiredString(val) {
		f.isError = true
		f.setMsg(field, requiredMsg, msg)
	}
}

func (f *Form) RequiredInt(val int, field string, msg ...string) {
	if requiredInt(val) {
		f.isError = true
		f.setMsg(field, requiredMsg, msg)
	}
}

func (f *Form) RequiredFloat64(val float64, field string, msg ...string) {
	if requiredFloat64(val) {
		f.isError = true
		f.setMsg(field, requiredMsg, msg)
	}
}

func (f *Form) RequiredBool(val bool, field string, msg ...string) {
	if requiredBool(val) {
		f.isError = true
		f.setMsg(field, requiredMsg, msg)
	}
}

func (f *Form) RequiredEmail(val string, field string, msg ...string) {
	if requireEmail(val) {
		f.isError = true
		f.setMsg(field, requiredMsg, msg)
	}
}

func (f *Form) NotNil(val interface{}, field string, msg ...string) {
	if notNil(val) {
		f.isError = true
		f.setMsg(field, requiredMsg, msg)
	}
}

func (f *Form) RequiredTime(val time.Time, field string, msg ...string) {
	if requiredTime(val) {
		f.isError = true
		f.setMsg(field, requiredMsg, msg)
	}
}

func (f *Form) MinInt(val int, n int, field string, msg ...string) {
	if minInt(val, n) {
		f.isError = true
		f.setMsg(field, minCharMsg, msg)
	}
}

func (f *Form) MaxInt(val int, n int, field string, msg ...string) {
	if maxInt(val, n) {
		f.isError = true
		f.setMsg(field, maxCharMsg, msg)
	}
}

func (f *Form) MaxFloat64(val float64, n float64, field string, msg ...string) {
	if maxFloat64(val, n) {
		f.isError = true
		f.setMsg(field, maxCharMsg, msg)
	}
}

func (f *Form) MinFloat64(val float64, n float64, field string, msg ...string) {
	if minFloat64(val, n) {
		f.isError = true
		f.setMsg(field, minCharMsg, msg)
	}
}

func (f *Form) MinChar(val string, n int, field string, msg ...string) {
	if minChar(val, n) {
		f.isError = true
		f.setMsg(field, minCharMsg, msg)
	}
}

func (f *Form) MaxChar(val string, n int, field string, msg ...string) {
	if maxChar(val, n) {
		f.isError = true
		f.setMsg(field, maxCharMsg, msg)
	}
}

func (f *Form) Email(val, field string, msg ...string) {
	if email(val) {
		f.isError = true
		f.setMsg(field, emailMsg, msg)
	}
}

func (f *Form) Gender(val, field string, msg ...string) {
	if gender(val) {
		f.isError = true
		f.setMsg(field, genderMsg, msg)
	}
}

func (f *Form) Confirm(val, confirm, field string, msg ...string) {
	if confirmVals(val, confirm) {
		f.isError = true
		f.setMsg(field, confirm, msg)
	}
}

func (f *Form) ISO8601DataTime(val, field string, msg ...string) {
	if iso8601DataTime(val) {
		f.isError = true
		f.setMsg(field, dateMsg, msg)
	}
}

func (f *Form) InString(val string, in []string, field string, msg ...string) {
	if inString(val, in) {
		f.isError = true
		f.setMsg(field, inMsg, msg)
	}
}
