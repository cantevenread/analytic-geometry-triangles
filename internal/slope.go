package internal

import "strconv"

type Slope struct {
	Rise, Run float64
}

func FormatSlope(slope Slope) string {
	formatted := strconv.FormatFloat(slope.Rise, 'f', -1, 64) + "/" + strconv.FormatFloat(slope.Run, 'f', -1, 64)
	return formatted
}

func FindSlope(p1, p2 Point) Slope {
	return Slope{Rise: p2.Y - p1.Y, Run: p2.X - p1.X}
}
