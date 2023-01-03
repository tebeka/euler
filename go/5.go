/*
2520 is the smallest number that can be divided by each of the numbers from 1 to
10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the
numbers from 1 to 20?

Answer: 232792560
*/

package main

import "fmt"

func isDivisible(n int) bool {
	for i := 1; i < 21; i++ {
		if n%i != 0 {
			return false
		}
	}
	return true
}

func e5() {
	for i := 20; true; i++ {
		if isDivisible(i) {
			fmt.Println(i)
			break
		}
	}
}

func init() {
    register("5", e5)
}
