package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
	//greetings()
	//Rectangle()
	//random()
	converter()
}

func greetings(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Printf("Hello, %s", name)
}

func Rectangle(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the lenght: ")
	lengthInput, _ := reader.ReadString('\n')
	lengthInput = strings.TrimSpace(lengthInput)
	length, err1 := strconv.ParseFloat(lengthInput, 64)

	fmt.Print("Enter the width: ")
	widthInput, _ := reader.ReadString('\n')
	widthInput = strings.TrimSpace(widthInput)
	width, err2 := strconv.ParseFloat(widthInput, 64)

	if err1 != nil || err2 != nil {
		fmt.Println("Invalid Inputs. Try Again")
		return
	}

	area := length * width

	fmt.Printf("The area of the reactangle is: %v", area)
}

func random(){
	var attempts int
	target := rand.Intn(100)+1

	reader := bufio.NewReader(os.Stdin)
	for{
		fmt.Print("Enter your guess: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		number, err := strconv.Atoi(input)

		if err != nil{
			fmt.Println("Please enter a valid number")
			continue
		}

		attempts++

		if number < target {
			fmt.Println("Too low!")
		} else if number > target {
			fmt.Println("Too high!")
		} else {
			fmt.Printf("Correct! You guessed in %v attempts", attempts)
			break
		}
	}

}

