package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
	winning map[int]bool
	have []int
	score, copies int
}

func read_input(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	
	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	return lines
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func wsNumsToArr(str string) []int {
	nums := make([]int, 0)
	nums_str := strings.Split(str, " ")
	for _, s := range nums_str {
		if s == "" {
			continue
		}
		nums = append(nums, atoi(s))
	}
	return nums
}

func arrToSet(nums []int) map[int]bool {
	winning := make(map[int]bool)
	for _, n := range nums {
		winning[n] = true
	}
	return winning
}


func deserialize(lines []string) []Card {
	// Card  1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
	r, _ := regexp.Compile("Card[ ]+\\d+: (.*)$")

	cards := make([]Card, 0)
	for _, line := range lines {
		// chop off Card N
		parts := r.FindStringSubmatch(line)
		rest := parts[1]

		groups := strings.Split(rest, " | ")
		winning := arrToSet(wsNumsToArr(groups[0]))
		have := wsNumsToArr(groups[1])
		card := Card{winning, have, 0, 1}
		card.scoreCard()
		cards = append(cards, card)
	}
		
	return cards
}

func (c *Card) scoreCard() {
	for _, n := range c.have {
		if c.winning[n] {
			c.score++
		}
	}
}

func addDupes(cards []Card, pos int) {
	for times := 0 ; times < cards[pos].copies ; times++ {
		for i := 1 ; i <= cards[pos].score ; i++ {
			cards[pos+i].copies++
		}
	}
}

func main() {
	lines := read_input("input/4")
	cards := deserialize(lines)
	for i, card := range cards {
		if card.score > 0 {
			addDupes(cards, i)
		}
	}
	// count cards
	sum := 0
	for _, card := range cards {
		sum += card.copies
	}
	fmt.Println("Total: ", sum)
}
