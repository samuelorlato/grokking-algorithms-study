package recursion

func SumFromSlice(numericSlice []float64) float64 {
	if len(numericSlice) == 0 {
		return 0
	}

	return numericSlice[0] + SumFromSlice(numericSlice[1:])
}
