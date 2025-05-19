package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Analyzer(){
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the text: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	characterCount := len(input)

	words := strings.Fields(input)
	wordcount := len(words) 

	

	fmt.Print("Enter the specific letter: ")
	letter, _ := reader.ReadString('\n')
	letter = strings.TrimSpace(letter)

	character := letter[0]

	
	fmt.Printf("\nNumber of words: %v", wordcount)
	fmt.Printf("\nNumber of Characters: %v", characterCount)

	specNumber := 0

	for i := 0; i < len(input); i++ {
		if input[i] == character {
			specNumber++
		} 
	}
	fmt.Printf("\nNumber of Specific letters (%v): %v", letter, specNumber)
}