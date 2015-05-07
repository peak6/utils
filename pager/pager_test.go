package pager

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPage(t *testing.T) {
	c, p := Page(5, 0, 10)
	assert.Equal(t, 1, c)
	assert.Equal(t, 2, p)
	c, p = Page(10, 21, 30)
	assert.Equal(t, 3, c)
	assert.Equal(t, 3, p)
	c, p = Page(10, 21, 0)
	assert.Equal(t, 0, c)
	assert.Equal(t, 0, p)
}

func TestPageToRange(t *testing.T) {
	o := Offset(3, 2)
	assert.Equal(t, 3, o)
}

func BenchmarkGetPage(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Page(5, 6, 10)
	}
}
