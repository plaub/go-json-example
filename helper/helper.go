package helper

// Generic function to sum up all values in an array.
func GetArrayValueSum[V int64 | float64](m []V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
