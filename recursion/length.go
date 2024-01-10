package recursion

func LengthFromSlice[T any](slice []T) int {
	if cap(slice) == 0 {
		return 0
	}

	return 1 + LengthFromSlice(slice[1:])
}
