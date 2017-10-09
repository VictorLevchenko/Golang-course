package main

import "fmt"

func main() {
	intList := []int{1, 2, 3, 4, 5}

	fmt.Println(intList)
	for i := range intList {
		fmt.Println(i)
	}
	for i, v := range intList {
		fmt.Printf("Index = %d, Value = %d\n", i, v)
	}
}
