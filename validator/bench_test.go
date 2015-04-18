package validator

import (
	"errors"
	"testing"
)

var errTest = errors.New("error")

func BenchmarkRequire(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := NewCheck()
		v.RequiredString("val", errTest)
	}
}

func BenchmarkMinChar(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := NewCheck()
		v.MinChar("s", 2, errTest)
	}
}

func BenchmarkEmail(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := NewCheck()
		v.Email("test@test.com", errTest)
	}
}

func BenchmarkGender(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := NewCheck()
		v.Gender("test", errTest)
	}
}

func BenchmarkConfirm(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := NewCheck()
		v.Confirm("123", "321", errTest)
	}
}

func BenchmarkISO8601DataTime(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := NewCheck()
		v.ISO8601DataTime(`2014-04-18T10:47:23+07:00`, errTest)
	}
}

func BenchmarkInString(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := NewCheck()
		v.InString("222", []string{"1", "2", "3", "4"}, errTest)
	}
}
