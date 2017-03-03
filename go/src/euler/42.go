/*
The nth term of the sequence of triangle numbers is given by, tn = ½n(n+1); so
the first ten triangle numbers are:

1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...

By converting each letter in a word to a number corresponding to its
alphabetical position and adding these values we form a word value. For example,
the word value for SKY is 19 + 11 + 25 = 55 = t10. If the word value is a
triangle number then we shall call the word a triangle word.

Using words.txt (right click and 'Save Link/Target As...'), a 16K text file
containing nearly two-thousand common English words, how many are triangle
words?

Answer: 162
*/
package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func triangleNum(n int) int {
	f := float64(n)
	return int((f / 2) * (f + 1))
}

func wordToNumber(word []byte) int {
	num := 0
	offset := byte('A' - 1)
	for _, c := range word {
		num += int(c - offset)
	}
	return num
}

func triangles(max int) map[int]bool {
	ts := make(map[int]bool)
	for i := 1; i <= max; i++ {
		ts[triangleNum(i)] = true
	}

	return ts
}

func e42() {
	data, _ := ioutil.ReadFile("words.txt")
	re := regexp.MustCompile("[A-Z]+")

	ts := triangles(1000)

	count := 0
	for _, word := range re.FindAll(data, -1) {
		n := wordToNumber(word)
		if ts[n] {
			count++
		}
	}

	fmt.Println(count)

}

func init() {
    register("42", e42)
}
