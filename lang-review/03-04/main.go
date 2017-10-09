package main

import "fmt"

//Create a new type called person which is STRUCT that stores fName and lName
//Using the STRUCT “person”, using a composite literal create a value of type
//person and assign it to a variable with the identifier “p1”; print out
//“p1”; print out just the field fName for “p1”
//Take the STRUCT “person” in the previous exercise and add a field called
//“favFood” that stores a slice of string; for the variable “p1” use a
//composite literal to add values to the field favFood; print out the values
//in favFood; also print out the values for “favFood” by using a for range loop

type person struct {
	fName string
	lName string
}
type personFood struct {
	fName   string
	lName   string
	favFood []string
}

func main() {

	p1 := person{
		fName: "Bill",
		lName: "Gates",
	}
	fmt.Println(p1)
	fmt.Println(p1.fName)
	fmt.Println("-------------------------------")
	p2 := personFood{"Bill", "Gates", []string{"Hotdog", "Chicken"}}
	fmt.Println(p2.favFood)
	for _, food := range p2.favFood {
		fmt.Println(food)
	}
}
