package main

import "fmt"

/*
tuple is a finite sorted list of elements.
It is a data structure that groups data
Tuples are typically immutable sequential collections
the only way to modify a tuple is to change the fields
operators such as + abd * can be applied to tuples
*/

// gets the power series of integer a and returns tuple of square of a
// and cube of a
func powerSeries(a int) (int, int, error) {
	square := a * a
	cube := square * a
	return square, cube, nil
}

func main() {

	square, cube, _ := powerSeries(3)
	fmt.Println("Square ", square, "Cube", cube)
}
