package main

import "fmt"

type shape interface {
	getArea() float64
}
type square struct {
	side float64
}
type rectangle struct {
	length  float64
	breadth float64
}

func main() {
	s := square{side: 2}
	r := rectangle{length: 2, breadth: 3}
	printArea(s)
	printArea(r)

}

func (s square) getArea() float64 {
	return s.side * s.side
}
func (r rectangle) getArea() float64 {
	return r.length * r.breadth
}
func printArea(s shape) {
	fmt.Println(s.getArea())
}
