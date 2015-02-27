package errors2

func ToError2(err error) *Errors2 {
	if e, ok := err.(*Errors2); ok {
		return e
	}

	return nil
}

func RespError(err error) (int, string) {
	e := ToError2(err)
	return e.code, e.message
}

func IsNotFound(err error) bool {
	if e := ToError2(err); e != nil {
		return e.ty == NotFound
	}

	return false
}

func IsInternal(err *Errors2) bool {
	if e := ToError2(err); e != nil {
		return e.ty == Internal
	}

	return false
}

func IsBadReq(err *Errors2) bool {
	if e := ToError2(err); e != nil {
		return e.ty == BadReq
	}

	return false
}

func IsUnauthorized(err *Errors2) bool {
	if e := ToError2(err); e != nil {
		return e.ty == Unauthorized
	}

	return false
}

func IsForbidden(err *Errors2) bool {
	if e := ToError2(err); e != nil {
		return e.ty == Forbidden
	}

	return false
}
