package main

import "fmt"

func main() {

	// Initialize a "vector" via a slice.
	var myVector []float64

	// Add a couple of components to the vector.
	myVector = append(myVector, 11.0)
	myVector = append(myVector, 5.2)

	fmt.Println(myVector)
	fmt.Println(myVector[1])
}
