// Write a function which takes a string as input and write it out as a reverse string
// like  "Aly" ==> "ylA" :)

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	input := args[1:]
	fmt.Printf("The Reversed string is: %q \n", reverseString(strings.Join(input, "")))

}

func reverseString(rs string) string {
	var reverse string
	for i := len(rs) - 1; i >= 0; i-- {
		reverse += string(rs[i])
	}
	return reverse
}
