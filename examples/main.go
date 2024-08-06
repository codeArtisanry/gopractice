package main

import (
	"fmt"
	"math"
)

func main() {
	var choice int
	fmt.Println("Enter 1 to calculate area and circumference of a circle")
	fmt.Println("Enter 2 to convert celsius into fahrenheit")
	fmt.Println("Enter 3 to use simple calculator")
	fmt.Println("Enter 4 to swap two variables values")
  fmt.Println("Enter 5 to find the largest of three numbers")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		calcAreaAndCircumference()
	case 2:
		convertCelsiusIntoFahrenheit()
	case 3:
		simpleCalculator()
	case 4:
		swapTwoVariablesValuesWithoutTemp()
  case 5:
    largestOfThreeNumbers()
	default:
		fmt.Println("Invalid choice")
	}
}

// +++++++++++++++++++++++++++++++++ Day 1 +++++++++++++++++++++++++++++++++

func calcAreaAndCircumference() {
	var radius float64
	var area, circumference float64

	fmt.Printf("Enter the radius of the circle:")
	fmt.Scanln(&radius)

	area = math.Pi * radius * radius
	circumference = 2 * math.Pi * radius

	fmt.Printf("Aream of the circle: %.2f\n", area)
	fmt.Printf("circumference of the circle: %.2f\n", circumference)
}

func convertCelsiusIntoFahrenheit() {
	var celsius, fahrenheit float64

	fmt.Println("Please enter the temperature in celsius: ")
	fmt.Scanln(&celsius)

	fahrenheit = (celsius * 9 / 5) + 32

	fmt.Printf("%.2f celsius is equal to %.2f fahrenheit\n", celsius, fahrenheit)

	fmt.Println("Enter the temperature in fahrenheit: ")
	fmt.Scanln(&fahrenheit)

	celsius = (fahrenheit - 32) * 5 / 9

	fmt.Printf("%.2f fahrenheit = %.2f Celsius\n", fahrenheit, celsius)
}

func simpleCalculator() {
	var num1, num2 float64
	var operator string

	fmt.Println("Enter the first number: ")
	fmt.Scanln(&num1)

	fmt.Println("Enter the second number: ")
	fmt.Scanln(&num2)

	fmt.Println("Enter the operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, num1+num2)
	case "-":
		fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, num1-num2)
	case "*":
		fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, num1*num2)
	case "/":
		fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, num1/num2)
	default:
		fmt.Println("Invalid operator")
	}
}

func swapTwoVariablesValuesWithoutTemp() {
	var a, b float64

	fmt.Println("Enter the value of a: ")
	fmt.Scanln(&a)

	fmt.Println("Enter the value of b: ")
	fmt.Scanln(&b)

	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("After swapping, a = %.2f and b = %.2f\n", a, b)

}

// +++++++++++++++++++++++++++++++++ Day 2 +++++++++++++++++++++++++++++++++

func largestOfThreeNumbers() {
	var num1, num2, num3 float64

	fmt.Println("Enter the first number: ")
	fmt.Scanln(&num1)
	fmt.Println("Enter the Second number: ")
	fmt.Scanln(&num2)
	fmt.Println("Enter the Three number: ")
	fmt.Scanln(&num3)

	largest := num1

	if num2 > largest {
		largest = num2
	}

	if num3 > largest {
		largest = num3
	}

	fmt.Printf("The largest number is %.2f\n", largest)
}
