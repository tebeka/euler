/*
Starting in the top left corner of a 2×2 grid, there are 6 routes (without
backtracking) to the bottom right corner.

How many routes are there through a 20×20 grid?

Answer: 137846528820
*/

package main

import "fmt"

const (
	MAX_X = 20
	MAX_Y = 20
)

type Point struct {
	X, Y int
}

func (p *Point) Hash() string {
	return fmt.Sprintf("%d, %d", p.X, p.Y)
}

var _cache map[string]uint64

func numPaths(p *Point) uint64 {
	value, ok := _cache[p.Hash()]
	if ok {
		return value
	}
	if p.X == MAX_X || p.Y == MAX_Y {
		return 1
	}

	np := numPaths(&Point{p.X + 1, p.Y}) + numPaths(&Point{p.X, p.Y + 1})
	_cache[p.Hash()] = np

	return np
}

func e15() {
	_cache = make(map[string]uint64)
	fmt.Println(numPaths(&Point{0, 0}))
}

func init() {
    register("15", e15)
}
