package selectionsort

import "golang.org/x/exp/constraints"

func findLowestValueIndex[T constraints.Ordered](slice []T) int {
	lowestIndex := 0

	for i := range slice {
		if slice[i] < slice[lowestIndex] {
			lowestIndex = i
		}
	}

	return lowestIndex
}

func remove[T constraints.Ordered](slice []T, i int) []T {
	// replace the element to delete with the last
	slice[i] = slice[len(slice)-1]
	// return slice without the last
	return slice[:len(slice)-1]
}

// O(n x n)
func SelectionSort[T constraints.Ordered](slice []T) []T {
	newSlice := []T{}

	for range slice {
		lowestIndex := findLowestValueIndex[T](slice)
		newSlice = append(newSlice, slice[lowestIndex])

		slice = remove[T](slice, lowestIndex)
	}

	return newSlice
}
