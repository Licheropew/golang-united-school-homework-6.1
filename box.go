package golang_united_school_homework

import (
	"errors"
)

const ERRORINDEXMESSAGE string = "index doesn't exist or index went out of the range"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if b.shapesCapacity <= len(b.shapes) {
		return errors.New("shapes capacity is full")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if len(b.shapes) <= i || i < 0 {
		return nil, errors.New(ERRORINDEXMESSAGE)
	}
	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if len(b.shapes) <= i || i < 0 {
		return nil, errors.New(ERRORINDEXMESSAGE)
	}
	res := b.shapes[i]
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	return res, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if len(b.shapes) <= i || i < 0 {
		return nil, errors.New(ERRORINDEXMESSAGE)
	}
	res := b.shapes[i]
	b.shapes[i] = shape
	return res, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var res float64
	for _, p := range b.shapes {
		res += p.CalcPerimeter()
	}
	return res
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var res float64
	for _, p := range b.shapes {
		res += p.CalcArea()
	}
	return res
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	circleIndex := make([]int, 0)
	for i, c := range b.shapes {
		switch c.(type) {
		case *Circle:
			circleIndex = append(circleIndex, i)
		}
	}
	if len(circleIndex) == 0 {
		return errors.New("no circles in the list")
	}
	x := 0
	for _, i := range circleIndex {
		i -= x
		b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
		x++
	}
	return nil
}
