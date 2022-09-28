package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	x1, y1        float64
	width, height float64
	rotation      float64
}

type Circle struct {
	x1, y1 float64
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r *Rectangle) Translate(a, b float64) {
	//a copy is not created because we pass in a pointer to the function
	//instead of the rectangle object itself
	r.x1 += a
	r.y1 += b
}

func (c *Circle) Translate(a, b float64) {
	c.x1 += a
	c.y1 += b
}

//write methods for rectangles and circles that translate each shape by a units in the x direction and b units in the y direction

func main() {

	Rectangles()
}

func Rectangles() {
	var r Rectangle
	r.width = 3.0
	r.height = 5.0
	fmt.Println(r.Area())

	r.Translate(1000, 2) // don't copy the rectangle
	fmt.Println(r.x1, r.y1)
}

func CirclePointers() {
	var c Circle
	var pointerToC *Circle

	//pointerToC is nil
	pointerToC = &c // it now points to c

	//Note: we could have done pointerToC := &c

	//with the pointer I can change the object itself
	(*pointerToC).x1 = -1.7
	(*pointerToC).y1 = 2.3

	//this is called pointer dereference

	fmt.Println(c.x1, c.y1)

	//Go does not care if you dereference pointers :)
	pointerToC.x1 = 1000
	pointerToC.y1 = 2

	fmt.Println(c.x1, c.y1)
}

func Pointers() {
	//a pointer is a reference to an address in RAM
	var a int = 14

	var b *int // b's type is a pointer to an integer
	//default value of any pointer is nil

	fmt.Println(a)

	fmt.Println(b)

	//so, let's make b refer to the LOCATION in RAM of a
	b = &a
	// && means "AND". & means "location of"

	fmt.Println(b)
}
