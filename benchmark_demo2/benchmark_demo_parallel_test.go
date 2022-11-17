package benchmark_demo2

import "testing"

func BenchmarkFooParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			dosomething()
		}
	})
}
