package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s *Square) End() Point {
	Point := Point{s.start.x + int(s.a), s.start.y + int(s.a)}
	return Point
}

func (s *Square) Area() uint {
	return uint(s.a * s.a)
}

func (s *Square) Perimeter() uint {
	return uint(s.a * 4)
}
