/*
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that
the 6th prime is 13.

What is the 10001st prime number?

Answer: 104743
*/

package main

import (
	"fmt"
)

func e7() {
	i := 0
	for p := range Primes() {
		i++
		if i > 10000 {
			fmt.Println(p)
			break
		}
	}

}

func init() {
    register("7", e7)
}
