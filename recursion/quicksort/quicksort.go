package quicksort

import (
	"math/rand"

	"golang.org/x/exp/constraints"
)

// O(n log n)
func Quicksort[T constraints.Ordered](slice []T) []T {
	if len(slice) <= 1 {
		return slice
	}

	pivot := slice[rand.Intn(len(slice))]

	var lesser []T
	var equal []T
	var bigger []T

	for _, e := range slice {
		if e < pivot {
			lesser = append(lesser, e)
		}

		if e == pivot {
			equal = append(equal, e)
		}

		if e > pivot {
			bigger = append(bigger, e)
		}
	}

	return append(append(Quicksort[T](lesser), equal...), Quicksort[T](bigger)...)
}
