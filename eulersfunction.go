package main

import "fmt"

func main() {
	for n := range 400 {
		fmt.Printf("Number %4d has %4d coprime numbers.\n", n, eulersFunction(n))
	}
}

func eulersFunction(n int) int {
	if n <= 0 {
		return 0
	}

	count := 0
	for i := 1; i <= n; i++ {
		if gcd(i, n) == 1 {
			count++
		}
	}

	return count
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
