package main

import (
	"fmt"
)

type Point struct {
	x, y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {

	var w Wheel
	w.x = 8 		//same with w.Circle.Point.x
	w.y = 8			//same with w.Circle.Point.y
	w.Radius = 5	//same with W.Circle.Radius
	w.Spokes = 20
	fmt.Printf("x:%d, y:%d, radius:%d, spokes:%d\n", w.x, w.y, w.Radius, w.Spokes)

	w2 := Wheel{Circle{Point{8,8}, 5}, 20}
	fmt.Printf("x:%d, y:%d, radius:%d, spokes:%d\n", w2.x, w2.y, w2.Radius, w2.Spokes)

	w3 := Wheel {
		Circle: Circle {
			Point: Point {x: 8, y:8},
			Radius: 5,
		},
		Spokes: 20,
	}
	fmt.Printf("x:%d, y:%d, radius:%d, spokes:%d\n", w3.x, w3.y, w3.Radius, w3.Spokes)

	fmt.Printf("%#v\n", w)
	fmt.Printf("%#v\n",w2)
	fmt.Printf("%#v\n",w3)
}
