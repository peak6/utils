package validator

import (
	"testing"
)

func BenchmarkRequire(b *testing.B) {
	v := NewForm()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.RequiredString("", `name`)
		v.Messages()
	}
}

func BenchmarkMinChar(b *testing.B) {
	v := NewForm()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MinChar("s", 2, "email")
	}
}

func BenchmarkEmail(b *testing.B) {
	v := NewForm()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.Email(`test@test.com`, `email`)
	}
}

func BenchmarkGender(b *testing.B) {
	v := NewForm()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.Gender(`test`, `name`)
	}
}

func BenchmarkConfirm(b *testing.B) {
	v := NewForm()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.Confirm("123", "321", "sss")
	}
}

func BenchmarkISO8601DataTime(b *testing.B) {
	v := NewForm()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.ISO8601DataTime(`2014-04-18T10:47:23+07:00`, "ddd")
	}
}

func BenchmarkInString(b *testing.B) {
	v := NewForm()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.InString("222", []string{"1", "2", "3", "4"}, "asd")
	}
}
