package main

import (
	"fmt"
)

//interface
type MyInterface interface {

	MyFuncInt() int

	MyFuncString() string
}

//concrete type
type MySimpleType int

func (t MySimpleType) MyFuncInt() int {
	return 1
}

func (t MySimpleType) MyFuncString() string {
	return "a"
}

type MySimpleType2 int 

func (t MySimpleType2) MyFuncInt() int {
	return 2
}	


type MySimpleType3 int 

func (t MySimpleType3) aFunction() int {
	return 3
} 


//empty interface
type emptyInterface interface {}

//main
func main() {

	var interf MyInterface

	myType := MySimpleType(1)
	myType2 := MySimpleType2(2)
	myType3 := MySimpleType3(3)
	interf = myType; //implements interface
	//interf = myType2 //error, doesn't implements all methods

	fmt.Println(interf.MyFuncString())

	var emptyInterf emptyInterface

	emptyInterf = myType3
	emptyInterf = myType2
	emptyInterf = myType
}