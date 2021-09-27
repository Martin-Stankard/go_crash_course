package main

import "fmt"

//var str  = "Hello World"


func main(){
	///BEHOLD THE POWER OF INFERRED TYPES!!!!
	//var str string = "Hello World"
	//var str  = "Hello World"
	str := ":= is the phalic method of variable declaration!!"+
	"Can only be used within a function"


	//var age int = 39
	var age uint64 = 39
	//uint8 is the weirdest one and it works.

	const isCool = true

	// this is a float64 size := 1.333
	const size float32 = 1.337

	fmt.Println(str, age, isCool, size)
	fmt.Printf("\"%s\" is a %T\n", str, str)
	fmt.Printf("%d is a %T\n", age, age)
	fmt.Printf("%t is a %T\n", isCool, isCool)
	fmt.Printf("%f is a %T\n", size, size)
}
