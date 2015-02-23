package errors2

func ToError2(err error) *Errors2 {
	if e, ok := err.(*Errors2); ok {
		return e
	}

	return nil
}

func IsNotFound(err *Errors2) bool {
	return err.ty == NotFound
}

func IsInternal(err *Errors2) bool {
	return err.ty == Internal
}

func IsBadReq(err *Errors2) bool {
	return err.ty == BadReq
}

func IsUnauthorized(err *Errors2) bool {
	return err.ty == Unauthorized
}

func IsForbidden(err *Errors2) bool {
	return err.ty == Forbidden
}
