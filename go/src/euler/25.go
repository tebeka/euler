/*
The Fibonacci sequence is defined by the recurrence relation:

    Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.

Hence the first 12 terms will be:

    F1 = 1
    F2 = 1
    F3 = 2
    F4 = 3
    F5 = 5
    F6 = 8
    F7 = 13
    F8 = 21
    F9 = 34
    F10 = 55
    F11 = 89
    F12 = 144

The 12th term, F12, is the first term to contain three digits.

What is the first term in the Fibonacci sequence to contain 1000 digits?

Answer: 4782
*/
package main

import (
	"fmt"
	"math/big"
)

func bigFibs() chan *big.Int {
	a, b := big.NewInt(1), big.NewInt(1)
	out := make(chan *big.Int)
	go func() {
		for {
			f := big.NewInt(0)
			f.Add(f, a)
			out <- f
			a, b = b, a.Add(a, b)
		}
	}()

	return out
}

func e25() {
	i := 0
	for f := range bigFibs() {
		i++
		if len(f.String()) >= 1000 {
			break
		}
	}
	fmt.Println(i)
}

func init() {
    register("25", e25)
}
