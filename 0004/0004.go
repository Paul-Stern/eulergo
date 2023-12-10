// task: https://projecteuler.net/problem=4

package main

import "fmt"

const t = 91 * 99

func main() {
	fmt.Println()
	fmt.Println(isPalindrome(t))
	// fmt.Println(isPalindrome(90 * 99))
	// fmt.Printf("%.*s", 2, fmt.Sprint(t))
	// fmt.Printf("%.*s", 2, fmt.Sprint(t))
	s := fmt.Sprint(t)
	fmt.Println(s, s[0])
}

func isPalindrome(n int) bool {
	b := fmt.Sprint(n / 100)
	e := fmt.Sprintf("%02d", n%100)
	// fmt.Printf("%s: %s", b, e)
	fmt.Println(n, b, e)
	for i := 0; i < 3; i++ {
		if b[i] != e[len(e)-1] {
			return false
		}
	}
	return true
}
