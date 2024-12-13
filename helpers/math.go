package helpers

import "math"

func AbsInt(x int) int {
	return AbsDiffInt(x, 0)
}

func AbsDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func Sgn(x float64) int {
	if math.Signbit(x) {
		return -1
	} else {
		return 1
	}
}

type Coordinate struct {
	X, Y int
}

type Vector struct {
	X, Y float64
}

func (v Vector) Build(v2 Vector) Vector {
	return Vector{v2.X - v.X, v2.Y - v.Y}
}

func (v Vector) Add(v2 Vector) Vector {
	return Vector{v2.X + v.X, v2.Y + v.Y}
}

type Line struct {
	A, B, AB Vector
}

func (l1 Line) CalcIntersection(l2 Line) (float64, float64) {
	num := (l1.A.X-l2.A.X)*(l2.A.Y-l2.B.Y) - (l1.A.Y-l2.A.Y)*(l2.A.X-l2.B.X)
	den := (l1.A.X-l1.B.X)*(l2.A.Y-l2.B.Y) - (l1.A.Y-l1.B.Y)*(l2.A.X-l2.B.X)
	// parallel lines
	if den == 0 {
		return 0, -1
	}

	t := float64(num) / float64(den)
	num = (l1.A.X-l1.B.X)*(l1.A.Y-l2.A.Y) - (l1.A.Y-l1.B.Y)*(l1.A.X-l2.A.X)
	u := -(float64(num) / float64(den))
	return t, u
}

func (l Line) GetVector(t float64) Vector {
	return Vector{X: l.A.X + t*l.AB.X, Y: l.A.Y + t*l.AB.Y}
}
