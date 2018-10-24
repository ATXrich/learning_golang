package shapes

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	newCircle := Circle{20.0}
	newRectangle := Rectangle{10.0, 10.0}

	fmt.Printf("newCircle Area: %2.f\n", newCircle.Area())
	fmt.Printf("newRectangle Area: %2.f\n", newRectangle.Area())
}
