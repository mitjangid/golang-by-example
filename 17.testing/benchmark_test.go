package testing

import (
	"testing"
)

// ============================================
// BENCHMARKS
// ============================================
// Benchmarks measure the performance of code
// Run with: go test -bench=. -benchmem
// -bench=. runs all benchmarks
// -benchmem shows memory allocation statistics

func BenchmarkAdd(b *testing.B) {
	// b.N is the number of iterations the benchmark should run
	// The benchmark runs the loop b.N times and measures the time
	for i := 0; i < b.N; i++ {
		Add(2, 3)
	}
}

// Benchmark with different input sizes
func BenchmarkAddLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1000000, 2000000)
	}
}

// Benchmark comparing two approaches
func BenchmarkSubtract(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Subtract(100, 50)
	}
}
