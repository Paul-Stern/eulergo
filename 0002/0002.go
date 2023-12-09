// task: https://projecteuler.net/problem=2
package main

import (
	"fmt"
)

const max = 4000000

func main() {
	fmt.Printf("sum: %d\n", evenFibSum(max))
}

func fib(n int) int {
	a, b := 0, 1
	for n > 0 {
		a, b = b, b+a
		n = n - 1
	}
	return b
}
func evenFibSum(m int) (sum int) {
	a, b := 0, 1
	for b <= m {
		if a%2 == 0 {
			sum += a
		}
		a, b = b, b+a
	}
	return
}

func out(n int) {
	fmt.Printf("fib(%d): %d\n", n, fib(n))
}
