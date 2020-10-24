package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scanf("%d\n", &T)
	for i := 0; i < T; i++ {
		var n int
		fmt.Scanf("%d\n", &n)
		sumOfSquare := 0
		sum := 0
		for i := 1; i <= n; i++ {
			sumOfSquare += i * i
			sum += i
		}
		squareOfSum := sum * sum
		fmt.Println(squareOfSum - sumOfSquare)
	}
}
