package internal

import "strconv"

type Slope struct {
	Rise, Run float64
}

func FormatSlope(slope Slope) string {
	formatted := strconv.FormatFloat(slope.Rise, 'f', -1, 64) + "/" + strconv.FormatFloat(slope.Run, 'f', -1, 64)
	return formatted
}

func FormatNegativeSlope(s Slope) Slope {
    if s.Rise < 0 {
        if s.Run < 0 {
            newslope :=  Slope{
                Rise: -s.Rise,
                Run: -s.Run,
            }
            return newslope
        }
    }
    return Slope{Rise: s.Rise, Run: s.Run}
}

func FindSlope(p1, p2 Point) Slope {
	return Slope{Rise: p2.Y - p1.Y, Run: p2.X - p1.X}
}

func SlopeReciprocal(slope Slope) Slope {
    return Slope{
        Rise: -slope.Run,
        Run: slope.Rise,
    }
}
