package benchmark_demo

import "testing"

func BenchmarkIsPaling(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPaling("aaaaa")
	}
}
