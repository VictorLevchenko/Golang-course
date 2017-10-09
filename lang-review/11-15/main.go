package main

import "fmt"

//Create a new type called “gator”. The underlying type of “gator” is an int.
//Using var, declare an identifier “g1” as type gator (var g1 gator). Assign a
//value to “g1”. Print out “g1”. Print the type of “g1” using fmt.Printf
//(“%T\n”, g1)

//Can you assign “g1” to “x”? Why or why not?

//Create a value of type gator. Call the greeting() method from that value.
type gator int

func (g gator) greeting()  {
	fmt.Println("Hello, I'm gator")
}

func main() {
	var g1 gator
	g1 = 1
	fmt.Printf("%T\n", g1)
	var x int
	x = 2
	fmt.Printf("%T\n", x)
	//x = g1 can't assign , types are uncompatible
	x = int(g1) //right
	fmt.Println("------------------------------")
	g1.greeting()
}
