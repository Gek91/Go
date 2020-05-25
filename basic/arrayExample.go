package main

import "fmt"

type Currency int

const (
	USD Currency = iota //underling type int
	EUR
	GBP
	RMB
)

func main() {
	
	var array [3]int;
	array[0] = 1
	array[1] = 11

	fmt.Println(array[0])
	fmt.Println(array[1])
	fmt.Println(array[2]) //zero value

	var array1 [3]int = [3]int{1, 2 ,3} //array literal
	fmt.Println(array1[2])

	//constand indexing
	symbol := [...]string{USD: "dollar", EUR: "euro", GBP: "sterlina"}
	fmt.Println(symbol[GBP])

	//array comparation
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // "true false false"
}
