// This is to sort in a map by key
package main

import (
	"fmt"
	"sort"
)

func sortingMap(m map[int]string) map[int]string {
	keys := make([]int, 0, len(m))
	for v := range m {
		keys = append(keys, v)
	}
	sort.Ints(keys)
	for _, i := range keys {
		fmt.Println("Keys :=>", i, "Value :=> ", m[i])
	}
	return m
}

func main() {
	m := map[int]string{
		1: "A",
		3: "F",
		0: "C",
		2: "B",
		4: "D",
	}
	sortingMap(m)
}
