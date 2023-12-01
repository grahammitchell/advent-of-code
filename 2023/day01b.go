package main

import (
	"bufio"
	"fmt"
	// "io/ioutil"
	//"math"
	"os"
	"strconv"
)

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

func is_spelled_out(line string, i int) string {
	words := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for _, word := range words {
		if i+len(word) <= len(line) && line[i:i+len(word)] == word {
			return word2i(word)
		}
	}
	return ""
}


func first_last(line string) string {
	first := -1
	last := -1
	a := ""
	b := ""
	// check for single digits
	for i, ch := range line {
		if '0' <= ch && ch <= '9' {
			last = i
			b = string(ch)
			if first == -1 {
				first = i
				a = string(ch)
			}
		}
	}
	// check for spelled-out digits
	for i := 0; i<len(line); i++ {
		if val := is_spelled_out(line, i); val != "" {
			if i > last {
				last = i
				b = val
			}
			if first == -1 || i < first {
				first = i
				a = val
			}
		}
	}
	return a+b
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func word2i(s string) string {
	words := map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}
	return words[s]
}


func main() {
	lines := read_input("input/1")
	sum := 0
	for _, line := range lines {
		s := first_last(line)
		sum += atoi(s)
	}
	fmt.Println("Total: ", sum)
}

