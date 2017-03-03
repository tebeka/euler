/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

Answer: 6857
*/
package main

import (
	"fmt"
	"math"
	"math/big"
)

const (
	n = 600851475143
)

func e3() {
	for start := int64(math.Sqrt(n)); start > 0; start-- {
		if !big.NewInt(start).ProbablyPrime(10) {
			continue
		}
		if n%start == 0 {
			fmt.Println(start)
			break
		}
	}
}

func init() {
    register("3", e3)
}
