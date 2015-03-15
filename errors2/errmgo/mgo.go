package errmgo

func Err(err error) error {
	if err == nil {
		return nil
	}

	if err.Error() == "not found" {
		return NewNotFound("not found")
	}

	return NewInternal(err.Error())
}
