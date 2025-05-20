package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"os"
)

func Calculator()float64{
	reader := bufio.NewReader(os.Stdin)
	var result float64 = 0

	fmt.Print("Enter the first number: ")
	number1, _ := reader.ReadString('\n')
	number1 = strings.TrimSpace(number1)
	number_01, err1 := strconv.ParseFloat(number1, 64)

	fmt.Print("Enter the second number: ")
	number2, _ := reader.ReadString('\n')
	number2 = strings.TrimSpace(number2)
	number_02, err2 := strconv.ParseFloat(number2, 64)

	fmt.Print("Enter the operation ( + , - , *, / ): ")
	operation, _ := reader.ReadString('\n')
	operation = strings.TrimSpace(operation)

	if err1 != nil || err2 != nil {
		fmt.Println("Invalid inputs")
		return 0
	}

	if operation == "+" {
		result = number_01 + number_02
	} else if operation == "-" {
		result = number_01 - number_02
	} else if operation == "*"{
		result = number_01 * number_02
	} else {
		result = number_01 / number_02
	}
	return result
}