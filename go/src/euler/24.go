/*
A permutation is an ordered arrangement of objects. For example, 3124 is one
possible permutation of the digits 1, 2, 3 and 4. If all of the permutations are
listed numerically or alphabetically, we call it lexicographic order. The
lexicographic permutations of 0, 1 and 2 are:

012   021   102   120   201   210

What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5,
6, 7, 8 and 9?
*/
package main

import (
	"fmt"
	"sort"
)

type Reversed []int

func (r Reversed) Len() int {
	return len(r)
}

func (r Reversed) Less(i, j int) bool {
	return r[i] > r[j]
}

func (r Reversed) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func isLast(items []int) bool {
	return sort.IsSorted(Reversed(items))
}

func lexicalPermutations(items []int) chan []int {
	out := make(chan []int)
	sort.Ints(items)
	go func() {
		for {
			out <- items
			if isLast(items) {
				close(out)
				return
			}
		}
	}()

	return out
}

var items = []int{0, 1, 2}

func e24() {
	var i = []int{3, 1, 2}
	fmt.Printf("%v\n", isLast(i))
	sort.Sort(Reversed(i))
	fmt.Printf("%v\n", i)
	fmt.Printf("%v\n", isLast(i))
}

func init() {
    register("24", e24)
}
