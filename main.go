package main

import (
	"fmt"

	"github.com/cantevenread/agt/internal"
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

	fmt.Printf("Distance between A and B: √%f or %.2f\n", distanceAB.Answer, distanceAB.SqrAnswer)
	fmt.Printf("Distance between B and C: √%f or %.2f\n", distanceBC.Answer, distanceBC.SqrAnswer)
	fmt.Printf("Distance between C and A: √%f or %.2f\n", distanceCA.Answer, distanceCA.SqrAnswer)

	fmt.Println("------------------------------------------------------------")

	fmt.Println("Midpoint of line A and B:", internal.FindMidpoint(pointA, pointB))
	fmt.Println("Midpoint of line B and C:", internal.FindMidpoint(pointB, pointC))
	fmt.Println("Midpoint of line C and A:", internal.FindMidpoint(pointC, pointA))

	fmt.Println("------------------------------------------------------------")

    fmt.Println("Slopes are not reduced or account for double negatives or denominator negatives.")
	fmt.Println("Slope of A to B: ", internal.FormatSlope(internal.FindSlope(pointA, pointB)))
	fmt.Println("Slope of B to C: ", internal.FormatSlope(internal.FindSlope(pointB, pointC)))
	fmt.Println("Slope of C to A: ", internal.FormatSlope(internal.FindSlope(pointC, pointA)))
}
