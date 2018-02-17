package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var n int
	fmt.Println("Enter a number:")
	fmt.Scan(&n)
	fmt.Println(factorial(n))
}
