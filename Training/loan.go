package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculatemonthlyPay(amount float64, term int)float64{
	monthyPay := amount / float64(term)
	return monthyPay
}

func loan() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the principle amount: ")
	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimSpace(input1)
	amount, err1 := strconv.ParseFloat(input1, 64)

	fmt.Print("Enter the interest rate: ")
	input2, _ := reader.ReadString('\n')
	input2 = strings.TrimSpace(input2)
	rate, err2 := strconv.ParseFloat(input2, 64)

	fmt.Print("Enter the loan term(in months): ")
	input3, _ := reader.ReadString('\n')
	input3 = strings.TrimSpace(input3)
	term, err3:= strconv.ParseFloat(input3, 64)

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("Invalid Inputs. Try again")
	}

	fullAmount := amount + ( (amount/100) * rate )

	monthyPayment := calculatemonthlyPay(fullAmount, int(term))

	fmt.Printf("The monthly payment is: $%.2f\n", monthyPayment)
}