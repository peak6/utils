package errsqlx

func Err(err error) error {
	if err == nil {
		return nil
	}

	switch err.Error() {
	case "sql: no rows in result set":
		return NewNotFound("not found")
	}

	return NewInternal(err.Error())
}
