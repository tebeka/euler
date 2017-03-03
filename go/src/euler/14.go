/*
The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:
13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1

It can be seen that this sequence (starting at 13 and finishing at 1) contains
10 terms. Although it has not been proved yet (Collatz Problem), it is thought
that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million
*/

package main

import (
	"fmt"
	"math/big"
)

var one *big.Int = big.NewInt(1)
var two *big.Int = big.NewInt(2)
var three *big.Int = big.NewInt(3)

func collatz(n int) int {
	// use big.Int since intermediate results overflow ints
	value := big.NewInt(int64(n))
	for length := 1; ; length++ {
		if value.Cmp(one) == 0 {
			return length
		}

		if value.Bit(0) == 0 {
			value.Div(value, two)
		} else {
			value.Mul(value, three)
			value.Add(value, one)
		}

	}
	// make the compiler happy
	return -1
}

func e14() {
	n, value := 1, 1

	for i := 2; i < 1000000; i++ {
		c := collatz(i)
		if c > value {
			n, value = i, c
		}
	}

	fmt.Println(n)
}

func init() {
    register("14", e14)
}
