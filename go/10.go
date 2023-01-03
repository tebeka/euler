/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.

Answer: 142913828922
*/

package main

import (
	"fmt"
	"math/big"
)

func e10() {
	sum := int64(0)
	for i := int64(0); i < 2000000; i++ {
		if big.NewInt(i).ProbablyPrime(10) {
			sum += i
		}
	}
	fmt.Println(sum)
}

func init() {
    register("10", e10)
}
