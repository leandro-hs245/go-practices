package main

import "fmt"

const version = "0.0.1"

func main() {
	/* Go is a statically typed language, which means that the type of a variable is determined at compile time.
	* `:=` is a shorthand for declaring and initializing a variable. It can only be used inside functions. The type of the variable is inferred from the value on the right-hand side.
	* `var` is used to declare a variable. It can be used both inside and outside functions. When used outside functions, it creates a package-level variable. When used inside functions, it creates a local variable.
	 */
	fmt.Println("Running version: ", version)
	var number int = 1
	fmt.Println(number)
	cpp, python, golang := true, false, "Golang!"
	fmt.Println("Languages: ", cpp, python, golang)
	zeroValue()
}

func zeroValue() {
	// Show zero values for different data types
	var i int
	var f float64
	var b bool
	var s string
	var c complex64
	var bt byte
	var r rune
	var m map[string][]int
	fmt.Println("Zero values for int: ", i)
	fmt.Println("Zero values for float64: ", f)
	fmt.Println("Zero values for bool: ", b)
	fmt.Println("Zero values for string: ", s)
	fmt.Println("Zero values for complex64: ", c)
	fmt.Println("Zero values for byte: ", bt)
	fmt.Println("Zero values for rune: ", r)
	fmt.Println("Zero values for map: ", m)
}
