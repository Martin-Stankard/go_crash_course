package main

import "fmt"
	

func main() {
	fmt.Println("ranges")

	ids :=[] int {1,2,3}

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	sum := 0

	for _, value := range ids {
		sum+=value
	}
	fmt.Println(sum)

	emails := map[string] int {"a":1,"c":2,"e":3,"g":4 }

	for key, value  := range emails {
		fmt.Println(key, value)
	}
}

