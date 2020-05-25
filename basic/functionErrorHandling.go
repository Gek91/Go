package main 

import "fmt"

func panicRecover() (result int, err error)  {

	defer func() {
		if p:= recover(); p != nil {
			err = fmt.Errorf("%v", p)
		}
	}()

	panic(fmt.Sprintf("errore"))

	result = 10

	return
}
	

func main() {

	_, err := panicRecover()
	fmt.Println(err)
}