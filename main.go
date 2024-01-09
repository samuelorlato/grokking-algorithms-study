package main

import (
	"fmt"

	"github.com/samuelorlato/grokking-algorithms-study/binary_search"
)

func main() {
	orderedList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}
	foundIndex, err := binary_search.BinarySearch[int](orderedList, 10)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	fmt.Println("Found index:", *foundIndex)
}