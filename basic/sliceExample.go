package main

import "fmt"

func main() {


	s := []int{2,3,5,7,11,13}
	fmt.Println(len(s)) //6
	fmt.Println(cap(s))	//6

	s = s[:4]
	fmt.Println(len(s))//4
	fmt.Println(cap(s))//6
	
	s = s[2:]
	fmt.Println(len(s)) //2
	fmt.Println(cap(s)) //4

	var k = make([]int, 3)
	fmt.Println(len(k)) 
	fmt.Println(cap(k))

	var z []int
	fmt.Println(z == nil) 
	fmt.Println(len(z)) 
	fmt.Println(cap(z)) 

	var y []int = []int{}
	fmt.Println(y == nil) 
	fmt.Println(len(y)) 
	fmt.Println(cap(y)) 

	//append
	var x []int
	x = append(x, 1)
	x = append(x, 2, 3)
	x = append(x, 4, 5, 6)
	x = append(x, x...) // append the slice x
	fmt.Println(x)
}
