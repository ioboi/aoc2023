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

func getAdjacentGears(e []string, column, row int) []gearPosition {
	var result []gearPosition
	minX := max(column-1, 0)
	minY := max(row-1, 0)

	maxX := min(len(e[row])-1, column+1)
	maxY := min(len(e)-1, row+1)

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			c := rune(e[y][x])
			if c == '*' {
				result = append(result, pos(x, y))
			}
		}
	}
	return result
}

type gearPosition struct {
	x int
	y int
}

func pos(x, y int) gearPosition {
	return gearPosition{x, y}
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

	gearPositions := make(map[gearPosition][]int)

	for y := range engineSchematic {
		for x := range engineSchematic[y] {
			t := engineSchematic[y][x]
			if rune(t) == '*' {
				gearPositions[pos(x, y)] = make([]int, 0)
			}
		}
	}

	for y := range engineSchematic {
		gears := make(map[gearPosition]bool)
		for x := range engineSchematic[y] {
			t := engineSchematic[y][x]

			if unicode.IsDigit(rune(t)) {
				builder.WriteByte(t)
				for _, gear := range getAdjacentGears(engineSchematic, x, y) {
					gears[gear] = true
				}
			}

			if !unicode.IsDigit(rune(t)) || (x+1) >= len(engineSchematic[y]) {
				if builder.Len() > 0 {
					partNumber, err := strconv.Atoi(builder.String())
					check(err)
					for gear := range gears {
						gearPositions[gear] = append(gearPositions[gear], partNumber)
					}
				}

				builder.Reset()
				gears = make(map[gearPosition]bool)
			}
		}
	}

	var sum int
	for _, gear := range gearPositions {
		if len(gear) == 2 {
			sum += gear[0] * gear[1]
		}
	}
	fmt.Println(sum)
}
