// Get the duplicated elements in a slice and count it.
// So we need to record the items in a map and compare it with the original slice
// if equals then remove the duplicated , if no then go forward :)

package main

import "fmt"

func main() {
	duplicatedSlice := []string{"Aly", "Loves", "Cat", "Cat", "Loves", "Aly"}
	fmt.Printf("The Slice is : %q \n", duplicatedSlice)
	duplication := count(duplicatedSlice)

	for index, value := range duplication {
		fmt.Printf("The Item : %s , Counts: %d \n", index, value)
	}
	remove := removeDuplicated(duplicatedSlice)
	fmt.Printf("The Slice after removing the duplication is %q \n", remove)

}

// This function returns a map of string with int returned values as we need to count.
func count(elements []string) map[string]int {
	isDuplicated := make(map[string]int)

	// Loop over the elements :
	for _, val := range elements {
		_, res := isDuplicated[val]
		if res {
			isDuplicated[val]++
		} else {
			isDuplicated[val] = 1
		}

	}
	return isDuplicated
}

// This function is to remove the duplicated item/s:
func removeDuplicated(elements []string) []string {
	checkDuplication := make(map[string]bool)
	newValues := []string{}

	for val := range elements {
		if checkDuplication[elements[val]] == true {
			break
		} else {
			checkDuplication[elements[val]] = true
			newValues = append(newValues, elements[val])
		}
	}
	return newValues
}
