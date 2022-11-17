package benchmark_demo3

// 斐波那契数列
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
func sum(a, b int) int {
	return a + b
}
