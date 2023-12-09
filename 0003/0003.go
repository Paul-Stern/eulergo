// task: https://projecteuler.net/problem=3

package main

import "fmt"

const num = 600851475143

func main() {
	fmt.Println(factor(num, 2))
}

func isPrime(n int) bool {
	for m := 2; m < n; m++ {
		if n%m == 0 {
			return false
		}
	}
	return true
}

func factor(n, f int) int {
	if isPrime(f) && n%f == 0 {
		return n / f
	} else {
		return factor(n, f+1)
	}
}
