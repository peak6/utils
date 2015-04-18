package assingle

import (
	"github.com/aerospike/aerospike-client-go"
	"github.com/plimble/utils/errors2"
	"github.com/tinylib/msgp/msgp"
	"sync"
)

type Iterator struct {
	records   []*aerospike.Record
	recordLen int
	index     int
	size      int
	err       error
	itrPool   *sync.Pool
}

func (it *Iterator) Scan(val msgp.Unmarshaler) bool {
	if it.err != nil {
		return false
	}

	for i := it.index; i < it.recordLen; i++ {
		if it.records[i] != nil {
			it.index = i
			break
		}
	}

	_, err := val.UnmarshalMsg(it.records[it.index].Bins[""].([]byte))
	if err != nil {
		it.err = errors2.NewInternal(err.Error())
		return false
	}

	it.index++

	return true
}

func (it *Iterator) Close() error {
	defer it.itrPool.Put(it)
	return it.err
}

func (it *Iterator) Size() int {
	return it.size
}
