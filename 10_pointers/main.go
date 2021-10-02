package main

import "fmt"
	

func main() {
	fmt.Println("pointers")
	a := 5
	//annoying like C++, space optional
	pointerOfA := &  a
	valueOfPointerOfA := *pointerOfA

	fmt.Println(a)
	fmt.Println(pointerOfA)
	fmt.Println(valueOfPointerOfA)

	*pointerOfA = 10
	fmt.Println(a)
	
}

