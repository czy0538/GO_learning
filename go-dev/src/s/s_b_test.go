package s

import "testing"

func Benchmark_S(b *testing.B) {
	for i := 1; i < b.N; i++ {
		S(1, 100, 1)
	}

}
