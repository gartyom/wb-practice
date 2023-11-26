package main

import "testing"

func Test_findAnagrams(t *testing.T) {
}

func Benchmark_findAnagrams(b *testing.B) {
	array := []string{"пятка", "пятка", "тяпка", "пятак", "листок", "слиток"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findAnagrams(&array)
	}
}
