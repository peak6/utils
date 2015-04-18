package validator

import (
	"sync"
)

var formPool sync.Pool

func init() {
	formPool = sync.Pool{}
	formPool.New = func() interface{} {
		return &Form{
			messages: make(map[string]string),
		}
	}
}
