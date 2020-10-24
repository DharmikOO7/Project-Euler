package main

import "fmt"

func main() {
	var T int
	fmt.Scanf("%d\n", &T)
	for i := 0; i < T; i++ {
		var n int
		fmt.Scanf("%d\n", &n)
		fmt.Println(sumOfMultiples3or5(n))
	}
}

func sumOfMultiples3or5(n int) int {
	s3 := sumOfMultiplesLessThanN(n, 3)
	s5 := sumOfMultiplesLessThanN(n, 5)
	s15 := sumOfMultiplesLessThanN(n, 15)
	return s3 + s5 - s15
}

func sumOfMultiplesLessThanN(n, x int) int {
	noOfMultiples := (n - 1) / x
	//1x + 2x + 3x + .... + yx
	return x * (noOfMultiples * (noOfMultiples + 1)) / 2
}
