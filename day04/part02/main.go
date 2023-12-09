package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err == nil {
		return
	}

	fmt.Printf("Unexpected error: %v\n", err)
	os.Exit(1)
}

type card struct {
	id  int
	win map[int]bool
	num []int
}

func (c card) matches() int {
	var count int
	for _, n := range c.num {
		if _, winning := c.win[n]; winning {
			count += 1
		}
	}

	return count
}

func parseCard(id int, line string) card {
	winAndNum := strings.Split(strings.Split(line, ":")[1], "|")
	win := make(map[int]bool)
	for _, n := range numbers(winAndNum[0]) {
		win[n] = true
	}

	return card{
		id:  id,
		win: win,
		num: numbers(winAndNum[1]),
	}
}

func numbers(list string) []int {
	var result []int
	for _, num := range strings.Split(list, " ") {
		raw := strings.Trim(num, " ")
		if len(raw) > 0 {
			n, err := strconv.Atoi(raw)
			check(err)
			result = append(result, n)
		}
	}
	return result
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: day04 <file>")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	defer func() {
		check(err)
	}()

	if os.IsNotExist(err) {
		fmt.Printf("File '%s' does not exist\n", os.Args[1])
		os.Exit(1)
	}

	check(err)

	s := bufio.NewScanner(file)
	s.Split(bufio.ScanLines)

	var originalCards []card
	var id int
	for s.Scan() {
		originalCards = append(originalCards, parseCard(id, s.Text()))
		id += 1
	}

	maxLen := len(originalCards)

	for i := 0; i < len(originalCards); i++ {
		originalCard := originalCards[i]
		from := originalCard.id + 1
		toMax := min(originalCard.id+originalCard.matches()+1, maxLen)
		for j := from; j < toMax; j++ {
			originalCards = append(originalCards, originalCards[j])
		}
	}

	fmt.Println(len(originalCards))
}
