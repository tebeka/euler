/*
Let d(n) be defined as the sum of proper divisors of n (numbers less than n
which divide evenly into n).  If d(a) = b and d(b) = a, where a ≠ b, then a and
b are an amicable pair and each of a and b are called amicable numbers.

For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55
and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and
142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.

Answer: 31626
*/
package main

import "fmt"

func sumFactors(n int) int {
	max := n/2 + 1
	count := 0
	for i := 1; i <= max; i++ {
		if n%i == 0 {
			count += i
		}
	}

	return count
}

func e21() {
	seen := map[int]bool{}
	for i := 1; i < 10000; i++ {
		sf := sumFactors(i)
		if sf != i && sumFactors(sf) == i {
			seen[i] = true
		}
	}

	sum := 0
	for a := range seen {
		sum += a
	}

	fmt.Println(sum)
}

func init() {
    register("21", e21)
}
