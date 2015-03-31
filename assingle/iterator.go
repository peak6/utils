package assingle

import (
	"github.com/aerospike/aerospike-client-go"
	"github.com/plimble/utils/errors2"
	"github.com/tinylib/msgp/msgp"
)

type Iterator struct {
	record []*aerospike.Record
	index  int
	size   int
	err    error
}

func NewIterator(records []*aerospike.Record) *Iterator {
	return &Iterator{records, 0, len(records), nil}
}

func (it *Iterator) Scan(val msgp.Unmarshaler) bool {
	if it.size < it.index+1 {
		return false
	}

	_, err := val.UnmarshalMsg(it.record[it.index].Bins[""].([]byte))
	if err != nil {
		it.err = errors2.NewInternal(err.Error())
		return false
	}

	it.index++

	return true
}

func (it *Iterator) GetError() error {
	return it.err
}

func (it *Iterator) Size() int {
	return it.size
}
