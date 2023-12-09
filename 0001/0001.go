package main

import(
	"fmt"
	"heywood/euler/0001/functions"
)

const(
	n = 1000
)

func main() {
	fmt.Printf("For n = %d,\nsum = %d.\n", n, functions.Summarise(n))
}
