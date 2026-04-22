package main

import "fmt"

var name = "Golang"

const version = "0.0.1"

const (
	monday = iota
	tuesday
	wednesday
	thursday
	friday
	saturday
	sunday
)

const (
	kb = 1 << (10 * iota)
	mb
	gb
	tb
	pb
	eb
)
const (
	_          = iota
	debugLevel = iota << 1
	infoLevel
	warningLevel
	errorLevel
)

func main() {
	/* Go is a statically typed language, which means that the type of a variable is determined at compile time.
	* `:=` is a shorthand for declaring and initializing a variable. It can only be used inside functions. The type of the variable is inferred from the value on the right-hand side.
	* `var` is used to declare a variable. It can be used both inside and outside functions. When used outside functions, it creates a package-level variable. When used inside functions, it creates a local variable.
	 */
	fmt.Println("Running version: ", version)
	var number int = 1
	fmt.Println(number)
	cpp, python, golang := true, false, "Golang!"
	fmt.Println("Languages: ", cpp, python, golang, name)
	zeroValue()
	printIota()
	shadowing()
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

func printIota() {
	/*
		`iota` is a predeclared identifier that represents successive untyped integer constants.
		It is reset to 0 whenever the `const` keyword appears in the source code and
		increments by 1 for each constant specification in the same block. It is often used to create enumerated constants.
	*/
	fmt.Println("Days of the week: ", monday, tuesday, wednesday, thursday, friday, saturday, sunday)
	fmt.Println("Data sizes: ", kb, mb, gb, tb, pb, eb)
	fmt.Println("Log levels: ", debugLevel, infoLevel, warningLevel, errorLevel)
}

func shadowing() {
	/*
		Variable shadowing occurs when a variable declared within a certain scope (e.g., an inner block) has the same name as a variable declared in an outer scope.
		The inner variable "shadows" the outer variable, meaning that within the inner scope,
		references to that variable name will refer to the inner variable rather than the outer one.
	*/
	x := 10
	fmt.Println("Outer x: ", x)
	{
		x := 11
		fmt.Println("Inner x: ", x)
	}
	fmt.Println("Outer x after inner block: ", x)
}
