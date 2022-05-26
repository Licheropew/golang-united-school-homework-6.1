package golang_united_school_homework

import (
	"math"
)

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (t Triangle) CalcArea() float64 {
	return math.Sqrt(3) * math.Pow(t.Side, 2) / 4
}

func (t Triangle) CalcPerimeter() float64 {
	return 3 * t.Side
}
