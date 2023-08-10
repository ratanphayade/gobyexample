package main

import "testing"

var digits = "23"

func BenchmarkCountWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		letterCombinations(digits)
	}
}

func BenchmarkCountWords_Parallel(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			letterCombinations(digits)
		}
	})
}
