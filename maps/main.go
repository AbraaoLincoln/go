package main

import "fmt"

// all the keys has to be the same type
// all the values has to be the same type
// the keys and value can be of different types

func main() { //key   //value types
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00FF00",
	}

	printMap(colors)

	//creates an empty map
	var m1 map[string]string

	fmt.Println(m1)

	//creates an empty map
	m2 := make(map[string]string)

	fmt.Println(m2)

	m2["hello"] = "world"

	fmt.Println(m2)

	//delete an entry on the map with the key
	delete(m2, "hello")

	fmt.Println(m2)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println(key, value)
	}
}
