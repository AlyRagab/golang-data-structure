/*
This is to accept an input number and search for it in a slice.
OUTPUT :
# go run SearchInSlice.go 1
> The number 1 is existed :)

# go run SearchInSlice.go 10
> The number 10 is Not There :(
*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	list := []int{1, 2, 3, 4, 5}
	args, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	input := args

	for _, v := range list {
		if v == input {
			fmt.Printf("The number %v is existed :) \n", input)
			break
		} else {
			fmt.Printf("The number %v is Not There :( \n", input)
			break
		}
	}
}
