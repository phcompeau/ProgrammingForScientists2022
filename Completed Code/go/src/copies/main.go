package main

import "fmt"

type Node struct {
	name string
	age  float64
}

type Tree struct {
	nodes []Node
	label string
}

func main() {
	//TreeBlah()
	//ObjectEquals()
	SliceEquals()
}

func SliceEquals() {
	// the following gives a compiler error
	/*
		a := []int{1, 5, 2, 1, 3}
		b := []int{1, 5, 2, 1, 3}
		if a == b {
			fmt.Println("Equal!")
		} else {
			fmt.Println("Not equal!")
		}
	*/
}

func ObjectEquals() {
	var v1, v2 Node
	v1.name = "Bro"
	v2.name = "Bro"
	v1.age = 17.0
	v2.age = 17.0
	if v1 == v2 {
		fmt.Println("Equal!")
	} else {
		fmt.Println("Not equal!")
	}
}

func TreeBlah() {
	var v1, v2 Node
	v1.name = "Phillip"
	v2.name = "DJ"
	v2.age = 21.7
	v2.age = 78.2
	var t Tree
	t.nodes = []Node{v1, v2}
	t.label = "This is tree t."

	s := t // fields should get copied!
	s.label = "This is tree s."

	s.nodes[0].name = "Fred"
	s.nodes[1].name = "Sally"
	s.nodes[0].age = 237483789.9
	s.nodes[1].age = 497238978.2

	fmt.Println("s:", s)
	fmt.Println("t:", t)
}

func Nodes() {
	var v1, v2 Node
	v1.name = "Hi"
	v1.age = 68.2

	v2 = v1 //fields get copied?

	fmt.Println(v2.name, v2.age)

	//next thing: let's set fields of v2
	v2.name = "Yo"
	v2.age = 14.6

	fmt.Println("v1", v1.name, v1.age)
}
