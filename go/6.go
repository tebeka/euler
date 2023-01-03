/*
The sum of the squares of the first ten natural numbers is,
1^2 + 2^2 + ... + 10^2 = 385

The square of the sum of the first ten natural numbers is,
(1 + 2 + ... + 10)^2 = 552 = 3025

Hence the difference between the sum of the squares of the first ten natural
numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred
natural numbers and the square of the sum.

Answer: 25164150
*/

package main

import "fmt"

func sumOfSquares(n uint64) uint64 {
	sum := uint64(0)
	for i := uint64(1); i <= n; i++ {
		sum += (i * i)
	}

	return sum
}

func squareOfSums(n uint64) uint64 {
	sum := uint64(0)
	for i := uint64(1); i <= n; i++ {
		sum += i
	}

	return sum * sum
}

func e6() {
	fmt.Println(squareOfSums(100) - sumOfSquares(100))
}

func init() {
    register("6", e6)
}
