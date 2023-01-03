/*
Work out the first ten digits of the sum of the following one-hundred 50-digit numbers.

[13-digits.txt]

Answer: 5537376230
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
)

func numbers(filename string) chan *big.Int {
	out := make(chan *big.Int)
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("error: can't open %s - %s", filename, err))
	}
	reader := bufio.NewReader(file)
	go func() {
		for {
			line, _, err := reader.ReadLine()
			if err == io.EOF {
				close(out)
				break
			} else if err != nil {
				panic(fmt.Sprintf("error: %s", err))
			}

			i := big.NewInt(0)
			i.SetString(string(line), 10)

			out <- i
		}
	}()

	return out
}

func e13() {
	sum := big.NewInt(0)
	for i := range numbers("13-digits.txt") {
		sum.Add(sum, i)
	}

	fmt.Println(sum.String()[:10])
}

func init() {
    register("13", e13)
}
