package main

import "fmt"

func main(){
	//arrays
	// var fruitArray[2] string

	// fruitArray[0] = "apple"
	// fruitArray[1] = "orange"

	fruitArray := [2] string {"asdf","zxcv"}
	fmt.Println(fruitArray)

	//slices
	fruitSlice := [] string {"qwer", "sdfg", "as many as we like", "one more"}
	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3]) //should be index 1 & 2

}