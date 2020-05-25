package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

type Circle struct {
	Point //anonymous field -> no name point
	Radius int
}

type Wheel struct {
	Circle //anonymous field -> no name circle
	Spokes int
}

func main() {

	var w Wheel
	w.X = 8 //point x -> instead of w.circle.point.x
	w.Y = 8 //point y -> insted of w.circle.point.y
	w.Radius = 5 //circle radius -> instead of w.circle.radius
	w.Spokes = 20 //wheel spokes

	dist := w.Distance(Point{1,1}) //call distance of point with the point 1,1

	fmt.Println(dist)

}