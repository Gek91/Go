package main 

import "fmt"

func incr(pointer *int) {
	*pointer++
}

func main() {

	x := 1
	p := &x //pointer

	fmt.Println(*p); //deferentiation

	*p = 2
	fmt.Println(x);

	p1 := p;
	fmt.Println(p1 == p); //comparable

	//passing pointer to function
	incr(&x)

	fmt.Println(x);

}