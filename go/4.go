/*
A palindromic number reads the same both ways. The largest palindrome made from
the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.

Answer: 906609
*/
package main

import "fmt"

func isPalindrom(n int) bool {
	str := fmt.Sprintf("%d", n)
	size := len(str)
	for i := 0; i < size/2; i++ {
		if str[i] != str[size-i-1] {
			return false
		}
	}
	return true
}

func e4() {
	max := 0
	for i := 1; i < 1000; i++ {
		for j := 1; j < 1000; j++ {
			n := i * j
			if isPalindrom(n) && n > max {
				max = n
			}
		}
	}

	fmt.Println(max)
}

func init() {
    register("4", e4)
}
