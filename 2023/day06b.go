package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	time, distance int64
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
	str = strings.TrimSpace(str)
	nums_str := strings.Split(str, " ")
	for _, s := range nums_str {
		if s == "" {
			continue
		}
		nums = append(nums, atoi(s))
	}
	return nums
}

func concat(nums []int) int64 {
	strs := make([]string, 0)
	for _, n := range nums {
		strs = append(strs, strconv.Itoa(n))
	}
	val, _ := strconv.ParseInt(strings.Join(strs, ""), 10, 64)
	return val
}

func deserialize(lines []string) []Race {
	// Time:      7  15   30
	// Distance:  9  40  200

	time := concat(wsNumsToArr(strings.Split(lines[0], ": ")[1]))
	dist := concat(wsNumsToArr(strings.Split(lines[1], ": ")[1]))
	
	races := make([]Race, 0)
	r := Race{time, dist}
	races = append(races, r)

		
	return races
}

func dumbWins(r Race) int {
	wins := 0
	for i := int64(1) ; i < r.time ; i++ {
		attempt := i * (r.time-i)
		if attempt > r.distance {
			wins++
		}
	}
	return wins
}


func main() {
	lines := read_input("input/6")
	races := deserialize(lines)
	sum := 1
	for _, race := range races {
		wins := dumbWins(race)
		if wins > 0 {
			sum *= wins
		}
	}
	fmt.Println("Total: ", sum)
}
