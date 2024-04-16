package structs

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

// Area method for Rectangle type
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

// Area method for Rectangle type
func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func Perimeter(r Rectangle) float64 {
	return (r.Width + r.Height) * 2
}

func Area(r Rectangle) float64 {
	return r.Height * r.Width
}
