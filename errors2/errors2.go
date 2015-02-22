package errors2

type errors struct {
	Message string `json:"error"`
}

type validateError struct {
	Message interface{} `json:"error"`
}

type errNotFound errors

func (e errNotFound) Error() string {
	return e.Message
}

func IsNotFound(err error) bool {
	_, ok := err.(errNotFound)
	return ok
}

func NewNotFound(msg string) error {
	return errNotFound{msg}
}

type errExist errors

func (e errExist) Error() string {
	return e.Message
}

func IsExist(err error) bool {
	_, ok := err.(errExist)
	return ok
}

func NewExist(msg string) error {
	return errExist{msg}
}

type errValidation validateError

func (e errValidation) Error() string {
	return "error validation"
}

func IsValidation(err error) bool {
	_, ok := err.(errValidation)
	return ok
}

func NewValidation(msg interface{}) error {
	return errValidation{msg}
}

func (e errors) Error() string {
	return e.Message
}

func IsErrors(err error) bool {
	_, ok := err.(errors)
	return ok
}

func NewErrors(msg string) error {
	return errors{msg}
}

type errUnauthorized errors

func (e errUnauthorized) Error() string {
	return e.Message
}

func IsUnauthorized(err error) bool {
	_, ok := err.(errUnauthorized)
	return ok
}

func NewUnauthorized(msg string) error {
	return errUnauthorized{msg}
}

type errForbidden errors

func (e errForbidden) Error() string {
	return e.Message
}

func IsForbidden(err error) bool {
	_, ok := err.(errForbidden)
	return ok
}

func NewForbidden(msg string) error {
	return errForbidden{msg}
}

func HttpCode(err error) int {
	switch err.(type) {
	case errNotFound:
		return 404
	case errExist, errValidation:
		return 400
	case errors:
		return 500
	case errUnauthorized:
		return 401
	case errForbidden:
		return 403
	}

	return 500
}

func JsonRes(err error) (int, error) {
	return HttpCode(err), err
}

func Sql(err error) error {
	if err == nil {
		return nil
	}

	switch err.Error() {
	case "sql: no rows in result set":
		return NewNotFound("not found")
	}

	return NewErrors(err.Error())
}

func Mgo(err error) error {
	if err == nil {
		return nil
	}

	if err.Error() == "not found" {
		return NewNotFound(err.Error())
	}

	return NewErrors(err.Error())
}
