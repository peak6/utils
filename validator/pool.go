package validator

import (
	"sync"
)

var checkPool sync.Pool
var formPool sync.Pool

func init() {
	checkPool = sync.Pool{}
	checkPool.New = func() interface{} {
		return &Check{}
	}

	formPool = sync.Pool{}
	formPool.New = func() interface{} {
		return &Form{
			messages: make(map[string]string),
		}
	}
}
