package split

import (
	"testing"
)

func Benchmark_Split(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split((i + b.N) / 2)
	}
}
