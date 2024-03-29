package maps

import "fmt"

var languages = map[int]string{
	3: "English",
	4: "French",
	5: "Spanish",
}

func ValidateMap() {

	var products = make(map[int]string)
	products[1] = "chair"
	products[2] = "table"

	var value string
	var i int
	for i, value = range languages {
		fmt.Println("language", i, ":", value)
	}
	fmt.Println("product with key 2", products[2])

	delete(products, 1)
	fmt.Println("products", products)
}
