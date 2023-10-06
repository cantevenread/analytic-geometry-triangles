package internal

import "math"

type Distance struct {
    Answer float64
    SqrAnswer float64
}

// Gets distance between two points
// NUMBER NEEDS TO BE SQR
func FindDistance(p1, p2 Point) Distance {
    eq1 := p2.X - p1.X
    eq2 := p2.Y - p1.Y
    eq3 := eq1*eq1 + eq2*eq2
    return Distance{Answer: eq3, SqrAnswer: math.Sqrt(eq3)}
}
