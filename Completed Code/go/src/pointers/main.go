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
	DeleteExample()
}

func DeleteExample() {
	a := make([]int, 5)
	for i := range a {
		a[i] = 2*i + 1
	}
	index := 2

	a = BetterDeleteElement(a, index)
	// this is why it's a = append(a, item) and not just append(a, item)

	fmt.Println("a is", a)

	fmt.Println("length of a is", len(a))

	b := a[3:5]
	fmt.Println("It's alive:", b)
}

func DeleteElement(list []int, index int) {
	list = append(list[:index], list[index+1:]...)
	fmt.Println("list is", list)

	fmt.Println(len(list))
}

func BetterDeleteElement(list []int, index int) []int {
	list = append(list[:index], list[index+1:]...)

	return list
}

func Slices() {
	a := make([]int, 5)
	ChangeFirst(a)
	fmt.Println(a)

	//a slice is a pointer to a location in an array
	//(it's more than that; it's a pointer + a length)

	var b []int
	b = make([]int, 10, 20)
	//create an array of length 20 and b points at the first 10 elements
	//so, b = make([]int, 10) is shorthand for b = make([]int, 10, 10)

	//let's set some elements of b
	for i := 0; i < 10; i++ {
		b[i] = -i - 1
	}

	fmt.Println(b)

	var q []int = b[8:15]

	fmt.Println("q is", q)

	b[9] = 2375908

	fmt.Println("q is", q)

	q[0] = -25837
	fmt.Println(b)
}

func ChangeFirst(list []int) {
	list[0] = 1
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
