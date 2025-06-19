package backend

import "testing"

func BenchmarkFullEvent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TestFullEvent()
	}
}

func BenchmarkInternalEvent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TestInternalEvent()
	}
}
