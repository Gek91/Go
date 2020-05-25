package main

import "fmt"

func myFunc1(x int) (result int) {
	
	result = x +1;

	return //return result
}

func myFunc2(x int) (result int, odd bool) {

	odd = x%2 == 1
	result = x-1

	return result, odd
}

func squares() func() int { //return a function
	var x int 
	return func() int { //anonymous function
			x++ //literal function can access variable in the enviroment in witch is defined
			return x * x
		}	
}

//variadic function
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
	total += val
	}
	return total
}

func main() {

	fmt.Println(myFunc1(1))

	fmt.Println(myFunc2(1))

	//first class value
	f := myFunc1
	//f = myFunc2 //error, no same type of func(int) int

	fmt.Printf("%T\n", f) // "func(int) int"
	fmt.Println(myFunc1(1))


	fmt.Println(squares()()) //1
	fmt.Println(squares()()) //1

	function := squares() //return the anonymous inner function

	//clouser!
	fmt.Println(function()) //x=1
	fmt.Println(function()) //x=2
	fmt.Println(function()) //x=3
	fmt.Println(function()) //x=4


	fmt.Println(sum(1))
	fmt.Println(sum(1,2))
	fmt.Println(sum(1,2,3))
}