package process

import "testing"

func Benchmark_Process(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Process(1, 100, 1)
	}
}
