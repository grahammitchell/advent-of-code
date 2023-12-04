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
		card := Card{winning, have}
		cards = append(cards, card)
	}
		
	return cards
}

func scoreCard(c Card) int {
	score := 0
	for _, n := range c.have {
		if c.winning[n] {
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
		}
	}
	// fmt.Println(c)
	return score
}

func main() {
	lines := read_input("input/4")
	cards := deserialize(lines)
	sum := 0
	for _, card := range cards {
		sum += scoreCard(card)
	}
	fmt.Println("Total: ", sum)
}
