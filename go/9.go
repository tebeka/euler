/*
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
a^2 + b^2 = c^2

For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.

Answer: 31875000
*/

package main

import "fmt"

func isPythagorean(a, b, c int) bool {
	return a*a+b*b == c*c
}

func e9() {
	for a := 1; a < 1000; a++ {
		for b := 1; b < 1000; b++ {
			for c := 1; c < 1000; c++ {
				if a+b+c == 1000 && isPythagorean(a, b, c) {
					fmt.Println(a * b * c)
					return
				}
			}
		}
	}
	fmt.Println(isPythagorean(3, 4, 5))
}

func init() {
    register("9", e9)
}
