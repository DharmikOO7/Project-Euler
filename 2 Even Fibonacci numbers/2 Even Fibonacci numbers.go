package main

//This program's C++ equivalent runs faster

import "fmt"

func main() {
	var T int
	fmt.Scanf("%d\n", &T)
	for i := 0; i < T; i++ {
		var n int64
		fmt.Scanf("%d\n", &n)
		fmt.Println(sumOfEvenFibonnaci(n))
	}
}

func sumOfEvenFibonnaci(n int64) int64 {
	var a, b int64 = 2, 8
	var sum int64 = 2
	for b <= n {
		sum += b
		a, b = b, 4*b+a
	}
	return sum
}
