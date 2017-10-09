package main

import "fmt"

//Initialize a MAP using a composite literal where the key is a string
// and the value is an int; print out the map; range over the map printing
// out just the key; range over the map printing out both the key and the value

func main() {
	mapInstance := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println(mapInstance)
	for k := range mapInstance {
		fmt.Println(k)
	}
	for k, v := range mapInstance {
		fmt.Printf("Key: %s, Value: %d\n", k, v)
	}
}
