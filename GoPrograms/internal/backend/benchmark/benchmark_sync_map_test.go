package backend

import "testing"

func BenchmarkSyncMapWithIndvKeys(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SyncMapWithIndvKeys()
	}
}

func BenchmarkSyncMapWithCombKeys(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SyncMapWithCombKeys()
	}
}
