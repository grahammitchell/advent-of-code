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

func first_last(line string) string {
	first := -1
	last := -1
	for i, ch := range line {
		if '0' <= ch && ch <= '9' {
			last = i
			if first == -1 {
				first = i
			}
		}
	}
	return string(line[first]) + string(line[last])
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
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
