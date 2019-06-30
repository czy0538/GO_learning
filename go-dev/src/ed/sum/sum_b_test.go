package sum

import "testing"

func Benchmark_Sum(b *testing.B) {
	for i := 1; i <= b.N; i++ {
		Sum(1, 100, 1)
	}
}
