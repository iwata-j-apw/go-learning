package main

import (
	"fmt"
	"godemy/internal/vertex"
)

type Vertex struct {
	X, Y int
}

func main() {
	v := vertex.New(3, 4)
	fmt.Println(AreaFunc(*v))                     // => 12
	fmt.Println(v.AreaMethod())                  // => 12
	fmt.Println(v.ScaleMethod(10).AreaMethod())  // => 1200
}

func AreaFunc(v Vertex) int {
	return v.X * v.Y
}

func (v Vertex) AreaMethod() int {
	return v.X * v.Y
}

func (v *Vertex) ScaleMethod(i int) *Vertex {
	v.X *= i
	v.Y *= i
	return v
}
