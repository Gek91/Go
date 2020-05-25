package main

import (
	"fmt"
)

func main() {
	
	p := new(int) //pointer to int variable
	p2 := new(int) //pointer to another variable

	fmt.Println(p == p2)

	*p = 2

	fmt.Println(*p)

}