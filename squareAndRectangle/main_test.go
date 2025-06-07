package main

import (
	"fmt"
	"testing"
)

func TestAreaOfShapes(t *testing.T) {
	s := square{side: 2}
	r := rectangle{length: 2, breadth: 3}

	fmt.Println(s, r)

	//areaOfSquare := s.getArea()
	//
	//if areaOfSquare != 4 {
	//	t.Errorf("Expected area of square is 4 but got , %v", areaOfSquare)
	//}
	//
	//areaOfRectangle := r.getArea()
	//
	//if areaOfRectangle != 4 {
	//	t.Errorf("Expected area of square is 6 but got , %v", areaOfRectangle)
	//}

}
