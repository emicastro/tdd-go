package shapes

import "math"

// Shape is an interface
type Shape interface {
	Area() float64
}

// Rectangle is a rectangle shape
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle is a spheric shape
type Circle struct {
	Radius float64
}

// Triangle is a shape with 3 sides
type Triangle struct {
	Base   float64
	Height float64
}

// Perimeter receives a height and a width and returns the perimeter of a rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area receives a width and a height and returns the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter receives a height and a width and returns the perimeter of a circle
func (c Circle) Perimeter() float64 {
	return c.Radius * (2 * math.Pi)
}

// Area receives a width and a height and returns the area of a circle
func (c Circle) Area() float64 {
	return math.Pi * (c.Radius * c.Radius)
}

// Area receives a width and a height and returns the area of a triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}
