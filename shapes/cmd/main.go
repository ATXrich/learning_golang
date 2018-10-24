package main

import (
	"fmt"

	"github.com/learn-go-with-tests/shapes"
)

func main() {
	newCircle := shapes.Circle{radius: 20.0}
	newRectangle := shapes.Rectangle{10.0, 10.0}

	fmt.Printf("newCircle Area: %2.f\n", newCircle.Area())
	fmt.Printf("newRectangle Area: %2.f\n", newRectangle.Area())
}
