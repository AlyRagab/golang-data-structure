/*Starting with the simple sorting "The swapping of elements"
   The result will be :
>> go run bubbleSortInt.go
   Before Sorting : [22 44 77 11 55 66 33]
   After Sorting  : [11 22 33 44 55 66 77]
*/
package main

import "fmt"

func bubbleSortInt(numbers []int) []int {
	// Starting the loop with the length of the slice
	for i := len(numbers); i > 0; i-- {
		// In this next iteration will be added n-1 then n-2 ...etc for the swap action
		for x := 1; x < i; x++ {
			if numbers[x-1] > numbers[x] {
				save := numbers[x]
				numbers[x] = numbers[x-1]
				numbers[x-1] = save
			}
		}
	}
	return numbers

}

func main() {
	givenSlice := []int{22, 44, 77, 11, 55, 66, 33}
	fmt.Printf("Before Sorting : %v\n", givenSlice)
	fmt.Printf("After Sorting  : %v\n", bubbleSortInt(givenSlice))

}
