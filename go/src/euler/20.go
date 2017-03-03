/*j
n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!

Answer: 648
*/

package main

import (
	"fmt"
	"math/big"
)

// sumDigits is defined in 16

func e20() {
	fact := big.NewInt(1)
	fact.MulRange(1, 100)

	fmt.Println(sumDigits(fact.String()))
}

func init() {
    register("20", e20)
}
