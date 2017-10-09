package main

import "fmt"

//Create a new type: vehicle. The underlying type is a struct. The fields:
//doors, color. Create two new types: truck & sedan. The underlying type of
//each of these new types is a struct. Embed the “vehicle” type in both truck
//& sedan. Give truck the field “fourWheel” which will be set to bool. Give
//sedan the field “luxury” which will be set to bool.

//Using the vehicle, truck, and sedan structs: using a composite literal,
//create a value of type truck and assign values to the fields; using a
//composite literal, create a value of type sedan and assign values to the
//fields. Print out each of these values. Print out a single field from each
//of these values.

//Give a method to both the “truck” and “sedan” types with the following
//signature transportationDevice() string .Have each func return a string
// saying what they do. Create a value of type truck and populate the fields.
// Create a value of type sedan and populate the fields. Call the method for
// each value.

//Create a new type called “transportation”. The underlying type of this new
//type is interface. An interface defines functionality. Said another way, an
//interface defines behavior. For this interface, any other type that has a
//method with this signature …
//transportationDevice() string
//… will automatically, implicitly implement this interface. Create a func
//called “report” that takes a value of type “transportation” as an argument.
//The func should print the string returned by “transportationDevice()” being
//called for any type that implements the “transportation” interface. Call
//“report” passing in a value of type truck. Call “report” passing in a value
//of type sedan.
func report(t transportation) {
	fmt.Println(t.transportationDevice())
}

type transportation interface {
	transportationDevice() string
}
type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	fourWheel bool
}

func (t *truck) transportationDevice() string {
	return fmt.Sprintln("Truck caring cargo")
}

type sedan struct {
	vehicle
	luxury bool
}

func (s *sedan) transportationDevice() string {
	return "Sedan transporting people"
}
func main() {
	aSedan := sedan{vehicle{4, "white"}, false}
	aTruck := truck{vehicle{2, "black"}, true}

	fmt.Println(aSedan.transportationDevice())
	fmt.Println(aTruck.transportationDevice())
	fmt.Println("-------------------------------")

	report(&aTruck)
	report(&aSedan)
}
