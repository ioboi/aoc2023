package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func check(err error) {
	if err == nil {
		return
	}

	fmt.Printf("Unexpected error: %v\n", err)
	os.Exit(1)
}

func isAdjacentToSymbol(e []string, column, row int) bool {
	minX := max(column-1, 0)
	minY := max(row-1, 0)

	maxX := min(len(e[row])-1, column+1)
	maxY := min(len(e)-1, row+1)

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			c := rune(e[y][x])
			if !unicode.IsDigit(c) && c != '.' {
				return true
			}
		}
	}
	return false
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: day03 <file>")
		os.Exit(1)
	}

	b, err := os.ReadFile(os.Args[1])
	defer func() {
		check(err)
	}()

	if os.IsNotExist(err) {
		fmt.Printf("File '%s' does not exist\n", os.Args[1])
		os.Exit(1)
	}

	check(err)

	engineSchematic := strings.Split(string(b), "\n")

	var builder strings.Builder
	var isAdjacent bool

	var sum int

	for y := range engineSchematic {
		for x := range engineSchematic[y] {
			t := engineSchematic[y][x]

			if unicode.IsDigit(rune(t)) {
				isAdjacent = isAdjacent || isAdjacentToSymbol(engineSchematic, x, y)
				builder.WriteByte(t)
			}

			if !unicode.IsDigit(rune(t)) || (x+1) >= len(engineSchematic[y]) {
				if builder.Len() > 0 && isAdjacent {
					partNumber, err := strconv.Atoi(builder.String())
					check(err)
					sum += partNumber
				}

				builder.Reset()
				isAdjacent = false
			}
		}
	}

	fmt.Println(sum)
}
