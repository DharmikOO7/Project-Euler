package main

import (
	"fmt"
	"math"
)

func main() {
	var T int
	fmt.Scanf("%d\n", &T)
	for i := 0; i < T; i++ {
		var n int64
		fmt.Scanf("%d\n", &n)
		fmt.Println(largestPrimeFactor(n))
	}
}

func largestPrimeFactor(n int64) int64 {
	pf := int64(1)
	if n%2 == 0 {
		pf = 2
		for n%2 == 0 {
			n /= 2
		}
	}
	for i := int64(3); i <= int64(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			if i > pf {
				pf = i
			}
			for n%i == 0 {
				n /= i
			}
		}
	}
	if n > 2 {
		if n > pf {
			pf = n
		}
	}
	return pf
}
