package main

import "fmt"

func sum_to_n_a(n int) int {
	if n == 0 {
		return n
	} else {
		return n + sum_to_n_a(n-1)
	}
}

func sum_to_n_b(n int) int {
	var total int
	for i := 0; i <= n; i++ {
		total += i
	}

	return total
}

func sum_to_n_c(n float64) float64 {
	summation := (n / 2) * (n + 1)
	return summation
}

func main() {
	result1 := sum_to_n_a(5)
	fmt.Println("result of Recursive Function is:", result1)
	// Time Complexity O(n) Perform recusrive calls that is dependent on size of n

	result2 := sum_to_n_b(5)
	fmt.Println("result of For Loop is:", result2)
	// Time Complexity O(n) Similarly, iterates over each number form 0 to n,
	// but it is more efficint compared to resusive approach

	result3 := sum_to_n_c(5)
	fmt.Println("result of Fummation Formula is:", result3)
	// Time Complexity O(1) Directly computes the sum using formula, does not depend on the size of n

}
