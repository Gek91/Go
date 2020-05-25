package main

import (
	"fmt"
	"math"
)

//struct
type Point struct{ 
	X, Y float64 
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

//named type
type PointDistance float64

func (p PointDistance) Distance() float64 {
	return float64(p)
}


func main() {

	var p Point = Point{1.0,2.0}

	var p1 Point = Point{2.0,1.0}

	fmt.Println(p.Distance(p1))
	fmt.Println(p.Distance(p))

	pointer := &p
	pointer.ScaleBy(1.0)
	fmt.Println(p.Distance(p1))
	fmt.Println((*pointer).Distance(p1))
	fmt.Println(pointer.Distance(p1)) //equivalent to previous

	(&p).ScaleBy(1.0)
	fmt.Println(p.Distance(p1))

	//equivalent to previus, valid only on variables
	p.ScaleBy(2.0)
	fmt.Println(p.Distance(p1))



	var d PointDistance = 0.2
	fmt.Println(d.Distance())

	distanceMethod := p.Distance //method value, variable containing reference to method distance being executed on the variable p
	fmt.Println(distanceMethod(p1))

	distance := Point.Distance //method expression, variable containing method reference
	fmt.Println(distance(p, p1))
}