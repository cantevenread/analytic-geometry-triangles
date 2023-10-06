package internal

import (
	"fmt"
)

type Point struct {
    X, Y float64
}

func ReadPoint() Point {
	var x, y float64

	fmt.Print("Enter x coordinate: ")
	if _, err := fmt.Scanln(&x); err != nil {
		fmt.Println("Error reading x coordinate:", err)
		return Point{} // Return an empty Point in case of error
	}

	fmt.Print("Enter y coordinate: ")
	if _, err := fmt.Scanln(&y); err != nil {
		fmt.Println("Error reading y coordinate:", err)
		return Point{} // Return an empty Point in case of error
	}

	return Point{X: x, Y: y}
}

