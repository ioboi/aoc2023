package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var lineRegex = regexp.MustCompile(`Game (\d+): (.*)`)
var setRegex = regexp.MustCompile(`(\d+) (red|green|blue)`)

type set struct {
	red   int
	green int
	blue  int
}

type game struct {
	id   int
	sets []set
}

func (g game) power() int {
	var maxRed, maxGreen, maxBlue = 1, 1, 1
	for _, set := range g.sets {
		if set.red > maxRed {
			maxRed = set.red
		}
		if set.green > maxGreen {
			maxGreen = set.green
		}
		if set.blue > maxBlue {
			maxBlue = set.blue
		}
	}

	return maxRed * maxGreen * maxBlue
}

func parseGame(line string) game {
	var game game
	matches := lineRegex.FindStringSubmatch(line)
	gameId, err := strconv.Atoi(matches[1])
	if err != nil {
		fmt.Printf("Unexpected error: %v\n", err)
		os.Exit(1)
	}
	game.id = gameId

	for _, singleSet := range strings.Split(matches[2], ";") {
		var set set
		for _, cubes := range strings.Split(singleSet, ",") {
			singleSetMatches := setRegex.FindStringSubmatch(cubes)
			cubeCount, err := strconv.Atoi(singleSetMatches[1])
			if err != nil {
				fmt.Printf("Unexpected error: %v\n", err)
				os.Exit(1)
			}

			switch singleSetMatches[2] {
			case "red":
				set.red = cubeCount
			case "green":
				set.green = cubeCount
			case "blue":
				set.blue = cubeCount
			}
		}
		game.sets = append(game.sets, set)
	}

	return game
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: day02 <file>")
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

	var sum int
	for scanner.Scan() {
		game := parseGame(scanner.Text())
		sum += game.power()
	}

	fmt.Println(sum)
}
