package main

import "testing"

var criteria = map[string]int{
	"uppercaseLetter": 3,
	"lowercaseLetter": 3,
	"digit":           2,
	"specialChar":     2,
}

var numPasswords = 5
var plen = 10

func BenchmarkCountWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generatePasswords(criteria, plen, numPasswords)
	}
}

//func BenchmarkCountWords_Parallel(b *testing.B) {
//	b.ResetTimer()
//	b.RunParallel(func(pb *testing.PB) {
//		for pb.Next() {
//			generatePasswords(criteria, plen, numPasswords)
//		}
//	})
//}
