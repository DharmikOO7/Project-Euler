package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(n int) bool {
	s := strconv.Itoa(int(n))
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var T int
	fmt.Scanf("%d\n", &T)
	for i := 0; i < T; i++ {
		var n int64
		fmt.Scanf("%d\n", &n)
		multiplesOf11 := make([]int, 81)
		multiplesOf11[0] = 110
		for i := 1; i < 81; i++ {
			multiplesOf11[i] = multiplesOf11[i-1] + 11
		}
		largestPalindrome := 101101
		for i := 80; i > -1; i-- {
			for j := 999; j > 99; j-- {
				num := multiplesOf11[i] * j
				if num < int(n) {
					if isPalindrome(num) && num > largestPalindrome {
						largestPalindrome = num
					}
				}
			}
		}
		fmt.Println(largestPalindrome)
	}
}
