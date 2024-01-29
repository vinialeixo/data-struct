package slices

func TwiceValues(slice []int) int {
	var i int
	var value int

	for i, value = range slice {
		slice[i] = 2 * value
	}
	return value
}
