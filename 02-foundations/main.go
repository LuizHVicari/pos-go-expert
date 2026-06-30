package main

import (
	"errors"
	"fmt"
)

// const is used to declare a constant value that cannot be changed (must be assigned at compile time)
const a = "Hello, World!"

// custom types can be used for expressiveness
type ID int

/*
string defaults to an empty string if not set
int defaults to 0 if not set
bool defaults to false if not set
*/
var (
	b bool
	c int
	d bool = true // can also declare and assign a variable in one line
	f ID   = 1
)

func main() {
	var d string = "Hello, Go!" // scoped variables must be used, otherwise they will throw an error
	println(d)
	e := "Hello, Go! (short declaration)" // short variable declaration, declares type and assigns value, can only be used inside a function
	println(e)
	e = "Hello, Go! (short declaration) - updated" // can reassign a value to a variable, but cannot change its type, cannot use := to reassign a value to an existing variable, must use = instead
	println(e)

	fmt.Printf("The type of E is %T\n", e) // %T is used to print the type of a variable

	var myArray [3]int // arrays are fixed size, cannot be resized, must declare size at compile time
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3

	fmt.Println("My array:", myArray)
	fmt.Println("Last element of my array:", myArray[len(myArray)-1]) // len() is used to get the length of an array, slice, or string)

	for i, v := range myArray { // i is the index, v is the value, range is used to iterate over an array, slice, or map, like for i in array in python
		println("Index:", i, "Value:", v)
	}

	mySlice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100} // slices are dynamic arrays, can be resized, can be created using a slice literal or the make() function
	fmt.Println("My slice:", mySlice)
	fmt.Printf("len=%d cap=%d %v\n", len(mySlice), cap(mySlice), mySlice)
	fmt.Printf("len=%d cap=%d %v\n", len(mySlice[:0]), cap(mySlice[:0]), mySlice[:0]) // slicing a slice returns a new slice, but does not copy the underlying array, so the capacity is still the same as the original slice

	mySlice = append(mySlice, 110, 120, 130) // append() is used to add elements to a slice, returns a new slice with the added elements, if the underlying array is full, a new array will be allocated and the old array will be copied to the new array
	fmt.Println("My slice:", mySlice)
	fmt.Printf("len=%d cap=%d %v\n", len(mySlice), cap(mySlice), mySlice)
	fmt.Printf("len=%d cap=%d %v\n", len(mySlice[:0]), cap(mySlice[:0]), mySlice[:0])

	payments := map[string]int{"Luiz": 1000, "John": 2000, "Jane": 3000} // maps are key-value pairs, can be created using a map literal or the make() function
	fmt.Println("Payments:", payments)
	fmt.Println("John's payment:", payments["John"]) // access a value in a map using its key
	delete(payments, "John")                         // delete() is used to remove a key-value pair from a map
	fmt.Println("Payments:", payments)
	payments["Paul"] = 4000 // add a new key-value pair to a map
	fmt.Println("Payments:", payments)

	payments2 := make(map[string]int) // make() is used to create a map, can also be used to create slices and channels
	payments2["Luiz"] = 1000
	payments2["John"] = 2000
	payments2["Jane"] = 3000
	fmt.Println("Payments2:", payments2)

	for name, payment := range payments { // iterate over a map using range, name is the key, payment is the value
		fmt.Printf("Name: %s, Payment: %d\n", name, payment)
	}

	for _, payment := range payments { // if you don't need the key, you can use _ to ignore it
		fmt.Printf("Payment: %d\n", payment)
	}

	val, err := sum(1, 2) // functions can return multiple values, but you can only access the first value using [0]
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Sum:", val)
	}
}

func sum(a, b int) (int, error) { // functions can have multiple parameters, and can return a value
	if a+b >= 50 {
		return a + b, errors.New("Sum is greater or equal to 50")
	}
	return a + b, nil
}
