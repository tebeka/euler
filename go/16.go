/*
215 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 21000?

Answer: 1366
*/

package main

import (
	"fmt"
	"math/big"
)

func sumDigits(s string) int {
	sum := 0
	for _, c := range s {
		sum += int(c) - '0'
	}

	return sum
}

func e16() {
	i := big.NewInt(2)
	i.Exp(big.NewInt(2), big.NewInt(1000), nil)
	fmt.Println(sumDigits(i.String()))
}

func init() {
    register("16", e16)
}
