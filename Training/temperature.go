package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func celsiusToFahrenheit(c float64) float64 {
	return (c * 9 / 5) + 32
}

func fahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func converter() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Temperature Converter")
	fmt.Println("Enter 'C' to convert from Celsius to Fahrenheit")
	fmt.Println("Enter 'F' to convert from Fahrenheit to Celsius")
	fmt.Print("Enter your choice ( C / F ): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	choice := strings.ToUpper(input)

	fmt.Print("Enter the temperature: ")
	tempInput, _ := reader.ReadString('\n')
	tempInput = strings.TrimSpace(tempInput)
	temp, err := strconv.ParseFloat(tempInput, 64)

	if err != nil {
		fmt.Println("Invalid temperature input. Try again")
	}

	if choice == "C" {
		result := celsiusToFahrenheit(temp)
		fmt.Printf("%v°C is %v°F\n", temp, result)
	} else if choice == "F" {
		result := fahrenheitToCelsius(temp)
		fmt.Printf("%v°F is %v°C\n", temp, result)
	} else {
		fmt.Println("invalid choice. Please enter C or F")
	}

}