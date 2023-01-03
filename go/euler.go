package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var solutions = map[string]func(){}

func register(name string, fn func()) {
	solutions[name] = fn
}

type Strings []string
func (s Strings) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s Strings) Len() int { return len(s) }
func (s Strings) Less(i, j int) bool {
	ival, _ := strconv.Atoi(s[i])
	jval, _ := strconv.Atoi(s[j])
	return ival < jval
}


func main() {
	list := flag.Bool("list", false, "list solutions")
	flag.Parse()

	if *list {
		var keys = Strings{}
		for key := range solutions {
			keys = append(keys, key)
		}
		sort.Sort(keys)
		for _, name := range keys {
			fmt.Println(name)
		}
	}

	for _, name := range flag.Args() {
		if fn, ok := solutions[name]; !ok {
			fmt.Fprintf(os.Stderr, "error: %s not found\n", name)
			os.Exit(1)
		} else {
			fn()
		}
	}
}
