package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ProcessFile(inputFile string, outputFile string, targetWord string) error {
	file, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("failed to open input file: %v", err)
	}
	defer file.Close()

	lineCount := 0
	wordCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineCount++
		line := scanner.Text()
		wordCount += strings.Count(strings.ToLower(line), strings.ToLower(targetWord))
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error while reading file: %v", err)
	}

	result := fmt.Sprintf("Total lines: %d\nOccurrences of '%s': %d\n", lineCount, targetWord, wordCount)

	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		return fmt.Errorf("failed to write to output file: %v", err)
	}

	return nil
}
