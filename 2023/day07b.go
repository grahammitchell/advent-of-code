package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var handTypes = map[string]int {
	"FIVE_OF_A_KIND": 600,
	"FOUR_OF_A_KIND": 500,
	"FULL_HOUSE": 400,
	"THREE_OF_A_KIND": 300,
	"TWO_PAIR": 200,
	"ONE_PAIR": 100,
	"HIGH_CARD": 0,
}

var cardLetters = []string{"0", "J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "j", "Q", "K", "A"}

type Hand struct {
	cards []int
	bid int
	groups map[int]int
	score int
}

func newHand(cards []int, bid int) Hand {
	h := Hand{cards, bid, make(map[int]int), 0}

	for n := 1 ; n <= 14 ; n++ {
		count := 0
		for _, c := range h.cards {
			if c == n {
				count++;
			}
		}
		h.groups[n] = count
	}

	if h.has_a(5) {
		h.score = handTypes["FIVE_OF_A_KIND"]
	} else if h.has_a(4) {
		h.score = handTypes["FOUR_OF_A_KIND"]
	} else if h.has_full_house() {
		h.score = handTypes["FULL_HOUSE"]
	} else if h.has_a(3) {
		h.score = handTypes["THREE_OF_A_KIND"]
	} else if h.has_two_pair() {
		h.score = handTypes["TWO_PAIR"]
	} else if h.has_a(2) {
		h.score = handTypes["ONE_PAIR"]
	}

	h.score += h.first_card()

	return h
}

func (h Hand) has_a(count int) bool {
	for n := 2 ; n <= 14 ; n++ {
		if h.groups[n] == count {
			return true
		} else if h.groups[n] + h.jokers() == count {
			return true
		}
	}
	return false
}

func (h Hand) jokers() int {
	return h.groups[1]
}

func (h Hand) has_without_jokers(count int) bool {
	for n := 2 ; n <= 14 ; n++ {
		if h.groups[n] == count {
			return true
		}
	}
	return false
}

func (h Hand) has_full_house() bool {
	if h.has_without_jokers(3) && h.has_without_jokers(2) {
		return true
	} else if h.has_two_pair() && h.jokers() == 1 {
		return true
	}
	return false
}

func (h Hand) has_two_pair() bool {
	// no need to check for jokers here - if you can make two pair with jokers, you could have made three of a kind instead
	pairs := 0
	for n := 2 ; n <= 14 ; n++ {
		if h.groups[n] == 2 {
			pairs++
		}
	}
	return pairs == 2
}

func (h Hand) display() {
	label := "unknown"
	max_v := -1
	for k, v := range handTypes {
		if h.score > v && v > max_v {
			label = k
			max_v = v
		}
	}
	for _, c := range h.cards {
		fmt.Print(cardLetters[c])
	}
	fmt.Println(": ", h.score, " - ", label)
}

func (h Hand) first_card() int {
	return h.cards[0]
}

func (h Hand) beats(other Hand) bool {
	if h.score == other.score {
		for i, c := range h.cards {
			if c > other.cards[i] {
				return true
			} else if c < other.cards[i] {
				return false
			}
		}
	}
	return h.score > other.score
}

func sort_by_rank(hands []Hand) {
	sort.Slice(hands, func(i, j int) bool { return hands[j].beats(hands[i])})
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
		'T': 10, 'J': 1, 'Q': 12, 'K': 13, 'A': 14}

	hands := make([]Hand, 0)
	
	for _, line := range lines {
		parts := strings.Split(line, " ")
		bid := atoi(parts[1])
		cards := make([]int, 0)
		for _, ch := range parts[0] {
			value := cardLookup[ch]
			cards = append(cards, value)
		}
		hand := newHand(cards, bid)
		hands = append(hands, hand)
	}
		
	return hands
}


func main() {
	lines := read_input("input/7")
	hands := deserialize(lines)
	sort_by_rank(hands)
	winnings := 0
	for i, hand := range hands {
		hand.display()
		winnings += (i+1) * hand.bid
	}
	
	fmt.Println("Total winnings: ", winnings)
	
}
