package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

// from https://groups.google.com/g/golang-nuts/c/oPuBaYJ17t4/m/PCmhdAyrNVkJ
func reverse(input string) string {
	n := 0
	rune := make([]rune, len(input))
	for _, r := range input {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	return string(rune)
}

func firstNumber(input string) (int, error) {
	for _, character := range input {
		if unicode.IsNumber(character) {
			return strconv.Atoi(string(character))
		}
	}

	return 0, fmt.Errorf("'%s' does not contain any numbers", input)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: day01 <file>")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Unexpected error: %v\n", err)
			os.Exit(1)
		}
	}()

	if os.IsNotExist(err) {
		fmt.Printf("File '%s' does not exist\n", os.Args[1])
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("Unexpected error: %v\n", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var calibrationValue int

	for scanner.Scan() {
		if n, err := firstNumber(scanner.Text()); err != nil {
			fmt.Printf("Unexpected error: %v\n", err)
			os.Exit(1)
		} else {
			calibrationValue += n * 10
		}

		if n, err := firstNumber(reverse(scanner.Text())); err != nil {
			fmt.Printf("Unexpected error: %v\n", err)
			os.Exit(1)
		} else {
			calibrationValue += n
		}
	}

	fmt.Println(calibrationValue)
}
