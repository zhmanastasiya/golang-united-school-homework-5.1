package square

type Point struct {
	x, y int
}

type Square struct {
	Point
	a uint
}

func (s Square) End() Point {
	Point := Point{s.x + s.a, s.y - s.a}
	return Point
}

func (s Square) Area() uint {
	return uint(s.a * s.a)
}

func (s Square) Perimeter() uint {
	return uint(s.a * 4)
}
