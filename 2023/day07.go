package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Hand struct {
	cards []int
	bid int 
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

func deserialize(lines []string) ([]Hand) {
	cardLookup := map[rune]int{'2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
		'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14}

	hands := make([]Hand, 0)
	
	for _, line := range lines {
		parts := strings.Split(line, " ")
		bid := atoi(parts[1])
		cards := make([]int, 0)
		for _, ch := range parts[0] {
			value := cardLookup[ch]
			cards = append(cards, value)
		}
		hand := Hand{cards, bid}
		hands = append(hands, hand)
	}
		
	return hands
}


func main() {
	lines := read_input("input/7s")
	hands := deserialize(lines)
	for _, hand := range hands {
		fmt.Println(hand)
	}
	
	fmt.Println("Lowest location is ")
	
}
