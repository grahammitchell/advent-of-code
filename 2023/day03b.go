package main

import (
	"bufio"
	"fmt"
	// "io/ioutil"
	//"math"
	"os"
	"strconv"
)

type PartNumber struct {
	num int
	row int
	start, end int
	isLegal bool
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

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func isSymbol(ch rune) bool {
	if ch == '.' || isDigit(ch) {
		return false
	}
	return true
}

func (p PartNumber) contains(row, col int) bool {
	return p.row == row && p.start <= col && col <= p.end 
}

func (p PartNumber) containsAnyNeighborOf(row, col int) bool {
	return p.contains(row-1, col-1) || // NW
		p.contains(row-1, col)   || // N
		p.contains(row-1, col+1) || // NE
		p.contains(row, col-1)   || // W
		p.contains(row, col+1)   || // E
		p.contains(row+1, col-1) || // SW
		p.contains(row+1, col)   || // S
		p.contains(row+1, col+1)    // SE
}

func deserialize(lines []string) []PartNumber {
	// 467..114..

	allParts := make([]PartNumber, 0)
	for row, line := range lines {
		part := PartNumber{0, row, -1, -1, false}
		for i, c := range line {
			if isDigit(c) {
				digit := atoi(string(c))
				part.num = (10*part.num) + digit
				part.end = i
				if part.start == -1 {
					part.start = i
				}
			} else {
				if part.num != 0 {
					//fmt.Println("NEW PART: ", part)
					allParts = append(allParts, part)
					part = PartNumber{0, row, -1, -1, false}
				}
			}	
		}
		if part.num != 0 {
			allParts = append(allParts, part)
			part = PartNumber{0, row, -1, -1, false}
		}
	}
		
	return allParts
}

func adjacentPartList(row, col int, allParts []PartNumber) []PartNumber {
	adjacentParts := make([]PartNumber, 0)
	for _, part := range allParts {
		if part.containsAnyNeighborOf(row, col) {
			adjacentParts = append(adjacentParts, part)
		}
	}
	return adjacentParts
}

func main() {
	lines := read_input("input/3")
	allParts := deserialize(lines)
	sum := 0
	for row, line := range lines {
		for col, ch := range line {
			if ch == '*' {
				adjacentParts := adjacentPartList(row, col, allParts)
				if len(adjacentParts) == 2 {
					//fmt.Println(adjacentParts)
					sum += adjacentParts[0].num * adjacentParts[1].num
				}
			}	
		}
	}
	fmt.Println("Total: ", sum)
}
