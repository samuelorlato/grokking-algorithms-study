package main

import (
	"fmt"

	"github.com/samuelorlato/grokking-algorithms-study/recursion"
)

func main() {
	slice := []float64{5, 2, 3, 50, 4, 4, 1}
	sum := recursion.MaxValueFromSlice(slice)
	fmt.Println(sum)
}
