package main

import "testing"

func BenchmarkCountWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountWords("sample.txt")
	}
}

func BenchmarkCountWords_Parallel(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			CountWords("sample.txt")
		}
	})
}
