// Another way of Sorting an integer slice values , by insertion the value with its previous one :
// We will do 2 iteration:
// First for every elements of the slice , Second for comparing the elements and go forward
// The second iteration starting with the value of j , then go back one step
// >> go run insertionSort.go
// Befor Sorting :[55 88 11 44 22 99]
// After Sorting: [11 22 44 55 88 99]

package main

import "fmt"

func insertionSort(insert []int) []int {
	for i := 1; i < len(insert); i++ {
		value := insert[i]
		j := i - 1
		for j >= 0 && insert[j] > value {
			insert[j+1] = insert[j]
			j--
		}
		insert[j+1] = value
	}
	return insert

}

func main() {
	value := []int{55, 88, 11, 44, 22, 99}
	fmt.Printf("Befor Sorting :%v\n", value)
	fmt.Printf("After Sorting: %v\n", insertionSort(value))
}
