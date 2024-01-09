package main

import (
	"fmt"

	selectionsort "github.com/samuelorlato/grokking-algorithms-study/selection_sort"
)

func main() {
	array := []int{7, 10, 1, 2, 99, 205, 15, 17, 901, 532}
	orderedArray := selectionsort.SelectionSort[int](array)
	fmt.Println(orderedArray)
}
