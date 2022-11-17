package benchmark_demo2

import "testing"

func BenchmarkFooSerial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dosomething()
	}

}
