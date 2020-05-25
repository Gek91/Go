package main

import (
	"fmt"
)

func main() {
	
	var m map[string]int = make(map[string]int)
	m["uno"] = 1
	m["due"] = 2

	fmt.Println(m["uno"])
	fmt.Println(m["due"])

	m = map[string]int {
		"uno" : 1,
		"due" : 2,
	}

	
	fmt.Println(m["due"])

	//remove
	delete(m, "uno")

	//safe operation
	fmt.Println(m["uno"])

	m = map[string]int {
		"uno" : 1,
		"due" : 2,
		"tre" : 3,
	}

	//loop on map
	for key, value := range m {
		fmt.Printf("%s\t%d\n", key, value)
	}


	//zero value

	var zeroMap map[string]int

	fmt.Println(zeroMap == nil)
	fmt.Println(len(zeroMap))


	if value, ok := m["uno"]; !ok {
		fmt.Println(value)
		fmt.Println(ok)
	}
	
	if value, ok := m["cinque"]; !ok {
		//Not executed
		fmt.Println(value)
		fmt.Println(ok)
	}
}