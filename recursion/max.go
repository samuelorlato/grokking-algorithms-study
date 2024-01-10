package recursion

func MaxValueFromSlice(slice []float64) float64 {
	if len(slice) == 0 {
		return 0
	}

	if slice[0] > MaxValueFromSlice(slice[1:]) {
		return slice[0]
	}

	return MaxValueFromSlice(slice[1:])
}
