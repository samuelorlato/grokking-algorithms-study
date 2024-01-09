package binarysearch

import (
	"errors"

	"golang.org/x/exp/constraints"
)

// O(log n)
func BinarySearch[T constraints.Ordered](orderedArray []T, elementToFind T) (*int, error) {
	start := 0
	end := len(orderedArray) - 1

	for start <= end {
		middle := (start + end) / 2
		try := orderedArray[middle]

		if try == elementToFind {
			foundIndex := middle
			return &foundIndex, nil
		}

		if try < elementToFind {
			start = middle + 1
		}
		if try > elementToFind {
			end = middle - 1
		}
	}

	return nil, errors.New("element not found")
}
