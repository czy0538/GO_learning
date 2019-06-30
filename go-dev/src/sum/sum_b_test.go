package sum

import "testing"

func Benchmark_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(i)
	}
}
