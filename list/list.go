package main

//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book

// importing fmt and container list packages
import (
	"container/list"
	"fmt"
)

/*
A list is a sequence of elements. Each element can be connected to another with a link in a
forward or backward direction
*/
// main method
func main() {
	var intList list.List
	intList.PushBack(11)
	intList.PushBack(23)
	intList.PushBack(34)

	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}
}

//The list is iterated through the for loop, and the element's value is accessed through the
//Value method.
