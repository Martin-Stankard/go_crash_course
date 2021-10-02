package main

import "fmt"
	

func main() {
	fmt.Println("mapsssss")

	emails := make(map[string]string)

	emails["bob"] = "bob@gmail.com"
	emails["will"] = "will@gmail.com"
	emails["mike"] = "mike@gmail.com"

	fmt.Println(emails)
	
	if val, ok := emails["asdf"]; ok {
		fmt.Println(val)
		} else {
		fmt.Println("it isn't there")
	}

	if _, ok := emails["bob"]; ok {
		fmt.Println("it IS there and because of _ instead of val , I don't need to use it")
		} else {
		fmt.Println("it isn't there")
	}
}
