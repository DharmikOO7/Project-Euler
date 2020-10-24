package main

import (
	"fmt"
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a []int) int {
	res := a[0]
	for i := 1; i < len(a); i++ {
		res = (a[i] * res) / (gcd(res, a[i]))
	}
	return res
}

func main() {
	var T int
	fmt.Scanf("%d\n", &T)
	for i := 0; i < T; i++ {
		var n int
		fmt.Scanf("%d\n", &n)
		arr1toN := make([]int, n)
		for i := 0; i < n; i++ {
			arr1toN[i] = i + 1
		}
		smallestMultiple1toN := lcm(arr1toN)
		fmt.Println(smallestMultiple1toN)
	}
}
