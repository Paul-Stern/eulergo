// task: https://projecteuler.net/problem=4

package main

import "fmt"

func main() {
	fmt.Println(91 * 9999)
}

func isPalindrome(n int) bool {
	b := fmt.Sprintf("%d", n)
	e := fmt.Sprintf("%d", n)
	return b == e
}
