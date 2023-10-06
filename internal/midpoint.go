package internal

func FindMidpoint(p1, p2 Point) Point {
    return Point{
        X: (p1.X + p2.X) / 2,
        Y: (p1.Y + p2.Y) / 2,
    }
}
