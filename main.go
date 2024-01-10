package main

import (
	"fmt"

	"github.com/samuelorlato/grokking-algorithms-study/recursion/quicksort"
)

func main() {
	slice := []float64{5, 2, 62, 654, 75, 34632, 6, 547, 546, 345, 236, 43, 62, 5, 345, 32, 43, 6, 43, 2345, 657, 65, 7, 35, 23, 5, 2}
	orderedSlice := quicksort.Quicksort[float64](slice)
	fmt.Println(orderedSlice)
}
