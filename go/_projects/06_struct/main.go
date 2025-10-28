package main

import "fmt"

func main() {
	point := Point{1, 1}

	point.SetX(100)
	fmt.Println(point)
}

type Point struct {
	X int
	Y int
}

func (p *Point) SetX(v int) {
	p.X = v
}
