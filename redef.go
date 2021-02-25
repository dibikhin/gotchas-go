// It shows possible redefining of predefined identificators in Go
package main

func main() {
	nil := 123

	true := 123
	false := 321

	bool := 123

	// func := 123 syntax error: unexpected :=, expecting (

	make := 123
	// ss := make([]string, 3) // cannot call non-function make (type int), declared at
	// type []string is not an expression

	string := 123
	// var s string // string is not a type
	s := "zxcv" // still works

	int := 123
	// var i int // int is not a type
	i := 123 // still works

	// if := 123 // syntax error: unexpected :=, expecting expression and so on
	println(nil)

	// println(!true) // invalid operation: ! int
	// println(!false) // invalid operation: ! int

	println(true)
	println(false)

	println(bool)

	println(make)

	println(string)
	println(s)

	println(int + 123)
	println(i)
}
