package twodslices

import "fmt"

func Twodslices(rows, cols int) {
	var twodslices = make([][]int, rows)
	var i int
	for i = range twodslices {
		twodslices[i] = make([]int, cols)
	}
	fmt.Println(twodslices)
}
