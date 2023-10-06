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
    fmt.Printf("Perimiter of triangle: %.2f\n", internal.FindPerimeter(distanceAB, distanceBC, distanceCA))
    fmt.Printf("Distance between Midpoint AB to Midpoint BC: √%f or %.2f\n", internal.FindDistance(internal.FindMidpoint(pointA, pointB),internal.FindMidpoint(pointB, pointC)).Answer, internal.FindDistance(internal.FindMidpoint(pointA, pointB),internal.FindMidpoint(pointB, pointC)).SqrAnswer)
    fmt.Printf("Distance between Midpoint AB to Midpoint CA: √%f or %.2f\n", internal.FindDistance(internal.FindMidpoint(pointA, pointB),internal.FindMidpoint(pointC, pointA)).Answer, internal.FindDistance(internal.FindMidpoint(pointA, pointB),internal.FindMidpoint(pointC, pointA)).SqrAnswer)
    fmt.Printf("Distance between Midpoint BC to Midpoint CA: √%f or %.2f\n", internal.FindDistance(internal.FindMidpoint(pointB, pointC),internal.FindMidpoint(pointC, pointA)).Answer, internal.FindDistance(internal.FindMidpoint(pointB, pointC),internal.FindMidpoint(pointC, pointA)).SqrAnswer)

	fmt.Println("------------------------------------------------------------")

	fmt.Println("Midpoint of line A and B:", internal.FindMidpoint(pointA, pointB))
	fmt.Println("Midpoint of line B and C:", internal.FindMidpoint(pointB, pointC))
	fmt.Println("Midpoint of line C and A:", internal.FindMidpoint(pointC, pointA))

	fmt.Println("------------------------------------------------------------")

    fmt.Println("Slopes are not reduced or account for double negatives")
	fmt.Println("Slope of A to B: ", internal.FormatSlope(internal.FindSlope(pointA, pointB)))
	fmt.Println("Slope of B to C: ", internal.FormatSlope(internal.FindSlope(pointB, pointC)))
	fmt.Println("Slope of C to A: ", internal.FormatSlope(internal.FindSlope(pointC, pointA)))
    fmt.Println("Slope of Midpoint AB to Midpoint BC: ", internal.FormatSlope(internal.FindSlope(internal.FindMidpoint(pointA, pointB), internal.FindMidpoint(pointB, pointC))))
    fmt.Println("Slope of Midpoint AB to Midpoint CA: ", internal.FormatSlope(internal.FindSlope(internal.FindMidpoint(pointA, pointB), internal.FindMidpoint(pointC, pointA))))
    fmt.Println("Slope of Midpoint BC to Midpoint CA: ", internal.FormatSlope(internal.FindSlope(internal.FindMidpoint(pointB, pointC), internal.FindMidpoint(pointC, pointA))))

    fmt.Println("Negative reciprocal of slope A to B", internal.FormatSlope(internal.SlopeReciprocal(internal.Slope{Rise: -1, Run: 1})))

    test := internal.Slope{Rise: -1, Run: -1}
    fmt.Println("Test", internal.FormatSlope(internal.FormatNegativeSlope(test)))
}
