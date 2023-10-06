package main

import (
	"github.com/cantevenread/agt/internal"
	"fmt"
)
func main() {
	fmt.Println("Enter coordinates for Point A:")
	pointA := internal.ReadPoint()
	fmt.Println("Enter coordinates for Point B:")
	pointB := internal.ReadPoint()
	fmt.Println("Enter coordinates for Point C:")
	pointC := internal.ReadPoint()

    distanceAB := internal.FindDistance(pointA, pointB)
    distanceBC := internal.FindDistance(pointB, pointC)
    distanceCA := internal.FindDistance(pointC, pointA)

    fmt.Println("------------------------------------------------------------")

    fmt.Println("Distance between A and B:","√", distanceAB.Answer)
    fmt.Println("Distance between B and C:","√", distanceBC.Answer)
    fmt.Println("Distance between C and A:","√", distanceCA.Answer)
}
