package main

// const is used to declare a constant value that cannot be changed (must be assigned at compile time)
const a = "Hello, World!"

/*
string defaults to an empty string if not set
int defaults to 0 if not set
bool defaults to false if not set
*/
var (
	b bool
	c int
	d bool = true // can also declare and assign a variable in one line
)

func main() {
	var d string = "Hello, Go!" // scoped variables must be used, otherwise they will throw an error
	println(d)
	e := "Hello, Go! (short declaration)" // short variable declaration, declares type and assigns value, can only be used inside a function
	println(e)
	e = "Hello, Go! (short declaration) - updated" // can reassign a value to a variable, but cannot change its type, cannot use := to reassign a value to an existing variable, must use = instead
	println(e)
}
