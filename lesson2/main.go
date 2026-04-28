package main

import "fmt"

func main() {
	complexNumbers()
	booleanValues()
	runeVaues()
	number, err := atoi("number")
	fmt.Println("String to Integer Conversion:", number, err)
}

func complexNumbers() {
	/*
		Complex numbers in Go are represented using the built-in complex type, which consists of a real part and an imaginary part.
		The real part is a float64, and the imaginary part is also a float64.
		Complex numbers can be created using the complex function, which takes two float64 arguments:
		- The real part
		- The imaginary part.
	*/
	complexNumber1 := complex(10, 20)
	realPart := real(complexNumber1)
	imaginaryPart := imag(complexNumber1)
	fmt.Printf("Complex Number: %v\n", complexNumber1)
	fmt.Printf("Real Part: %v\n", realPart)
	fmt.Printf("Imaginary Part: %v\n", imaginaryPart)
	complexNumber2 := complex(0, 1)
	fmt.Printf("Complex Number: %v\n", complexNumber2)
	fmt.Println("Addition on complex numbers: ", complexNumber1+complexNumber2)
}

func booleanValues() {
	/*
			Operator Precedence:
		1. Parentheses
		2. NOT (!)
		3. AND (&&)
		4. OR (||)
	*/
	formula := true && false || (false && true)
	fmt.Printf("Formula: %v\n", formula)
	var numberComparison bool = 0 == 1
	fmt.Printf("Number Comparison: %v\n", !numberComparison)
	if formula {
		fmt.Println("The formula is true")
	} else {
		fmt.Println("The formula is false")
	}
	if 1 == 0 {
		fmt.Println("1 is equal to 0")
	} else if 1 > 0 {
		fmt.Println("1 is greater than 0")
	} else {
		fmt.Println("1 is less than 0")
	}
}

func runeVaues() {
	japaneseYen := 165
	fmt.Printf("The Unicode code point for the Japanese Yen symbol is: %U\n", japaneseYen)
	fmt.Printf("The character for the Japanese Yen symbol is: %c\n", japaneseYen)
}

func atoi(s string) (int, error) {
	/*
		The atoi function takes a string as input and converts it to an integer.
		It handles negative numbers by checking if the first character is a '-' and adjusting the final result accordingly.
		The function iterates through each character in the string, ensuring that they are valid digits (between '0' and '9').
		If an invalid character is encountered, it returns an error.
		If the conversion is successful, it returns the resulting integer.
	*/
	var number int
	hasNegativeNumbers := false
	if s[0] == '-' {
		hasNegativeNumbers = true
		s = s[1:]
	}
	for _, char := range s {
		if char >= '0' && char <= '9' {
			number = number*10 + int(char-'0')
		} else {
			return 0, fmt.Errorf("Invalid input: %s", s)
		}
	}
	if hasNegativeNumbers {
		number *= -1
	}
	return number, nil
}
