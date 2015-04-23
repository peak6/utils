package ashelper

import (
	"github.com/aerospike/aerospike-client-go"
	"github.com/plimble/utils/errors2"
	"github.com/tinylib/msgp/msgp"
)

var (
	ErrKeyExist    = errors2.NewInternal("Key already exists")
	ErrIndexExist  = errors2.NewInternal("Index already exists")
	ErrKeyNotExist = errors2.NewInternal("Key is not exists")
	ErrNotFound    = errors2.NewNotFound("not found")
)

func MarshalMsgPack(v msgp.Marshaler) []byte {
	b, err := v.MarshalMsg(nil)
	if err != nil {
		panic(errors2.NewInternal(err.Error()))
	}

	return b
}

func UnmarshalMsgPack(data []byte, v msgp.Unmarshaler) {
	if _, err := v.UnmarshalMsg(data); err != nil {
		panic(errors2.NewInternal(err.Error()))
	}
}

func Err(err error) error {
	if err == nil {
		return nil
	}

	return errors2.NewInternal(err.Error())
}

func ErrIndex(err error) error {
	if err == nil {
		return nil
	}

	switch err.Error() {
	case "Index already exists":
		return ErrIndexExist
	}

	return errors2.NewInternal(err.Error())
}

func ErrGet(record *aerospike.Record, err error) error {
	switch {
	case err == nil && record != nil:
		return nil
	case err != nil:
		return errors2.NewInternal(err.Error())
	case record == nil:
		return ErrNotFound
	}

	return nil
}

func ErrDel(exist bool, err error) error {
	if !exist {
		return ErrKeyNotExist
	}

	if err != nil {
		return errors2.NewInternal(err.Error())
	}

	return nil
}

func ErrPut(err error) error {
	if err == nil {
		return nil
	}

	switch err.Error() {
	case "Key already exists":
		return ErrKeyExist
	default:
		return errors2.NewInternal(err.Error())
	}

	return nil
}
