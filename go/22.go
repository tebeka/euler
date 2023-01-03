/*
Using names.txt (right click and 'Save Link/Target As...'), a 46K text file
containing over five-thousand first names, begin by sorting it into alphabetical
order. Then working out the alphabetical value for each name, multiply this
value by its alphabetical position in the list to obtain a name score.

For example, when the list is sorted into alphabetical order, COLIN, which is
worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would
obtain a score of 938 × 53 = 49714.

What is the total of all the name scores in the file?

Answer: 871198282
*/

package main

import (
	"io/ioutil"
	"regexp"
	"sort"
)

func score(name string, position int) int {
	score := 0
	for _, c := range name {
		score += int(c) - 'A' + 1
	}

	return position * score
}

func readNamesSorted() []string {
	data, _ := ioutil.ReadFile("names.txt")
	re := regexp.MustCompile("[A-Z]+")
	names := re.FindAll(data, -1)

	// FIXME: There's probably a better way
	strs := make([]string, len(names))
	for i, name := range names {
		strs[i] = string(name)
	}
	sort.Strings(strs)

	return strs
}

func e22() {
	sum := 0
	for i, name := range readNamesSorted() {
		sum += score(name, i+1)
	}
	println(sum)
}

func init() {
    register("22", e22)
}
