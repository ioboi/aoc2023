package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var numberRegex = regexp.MustCompile(`one|two|three|four|five|six|seven|eight|nine|\d`)

func mapValue(input string) int {
	switch input {
	case "1":
		return 1
	case "2":
		return 2
	case "3":
		return 3
	case "4":
		return 4
	case "5":
		return 5
	case "6":
		return 6
	case "7":
		return 7
	case "8":
		return 8
	case "9":
		return 9
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	}
	panic("should never happen!!!")
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

	var calibrationValue int

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		matches := numberRegex.FindAllString(scanner.Text(), -1)
		v := mapValue(matches[0])*10 + mapValue(matches[len(matches)-1])
		calibrationValue += v
	}

	fmt.Println(calibrationValue)
}
