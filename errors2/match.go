package errors2

func Sql(err error) error {
	if err == nil {
		return nil
	}

	switch err.Error() {
	case "sql: no rows in result set":
		return NewNotFound("not found")
	}

	return NewInternal(err.Error())
}

func Mgo(err error) error {
	if err == nil {
		return nil
	}

	if err.Error() == "not found" {
		return NewNotFound("not found")
	}

	return NewInternal(err.Error())
}
