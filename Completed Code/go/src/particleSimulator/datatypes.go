package main

//Particle represents our particle variable.
type Particle struct {
	position         OrderedPair
	name             string
	radius           float64 // prob not necessary
	diffusionRate    float64
	red, green, blue uint8 //coloring particle
}

type Board struct {
	width, height float64
	particles     []*Particle
}

type OrderedPair struct {
	x, y float64
}
