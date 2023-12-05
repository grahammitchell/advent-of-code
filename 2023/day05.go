package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Map struct {
	title string
	dest_starts []int64
	src_starts []int64
	range_lens []int64
}

func (m Map) lookup(src_num int64) int64 {
	for i := 0; i < len(m.dest_starts); i++ {
		if m.src_starts[i] <= src_num && src_num < m.src_starts[i]+m.range_lens[i] {
			diff := src_num - m.src_starts[i]
			return m.dest_starts[i] + diff
		}
	}
	return src_num
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

func wsNumsToArr(str string) []int64 {
	nums := make([]int64, 0)
	nums_str := strings.Split(str, " ")
	for _, s := range nums_str {
		if s == "" {
			continue
		}
		nums = append(nums, int64(atoi(s)))
	}
	return nums
}

func deserialize(lines []string) ([]int64, []Map) {
	// seeds: 79 14 55 13

	// seed-to-soil map:
	// 50 98 2
	// 52 50 48

	// soil-to-fertilizer map:
	// 0 15 37
	// 37 52 2
	// 39 0 15

	// fertilizer-to-water map:
	// 49 53 8
	// 0 11 42
	// 42 0 7
	// 57 7 4

	// water-to-light map:
	// 88 18 7
	// 18 25 70

	// light-to-temperature map:
	// 45 77 23
	// 81 45 19
	// 68 64 13

	// temperature-to-humidity map:
	// 0 69 1
	// 1 0 69

	// humidity-to-location map:
	// 60 56 37
	// 56 93 4
	seeds := wsNumsToArr(strings.Split(lines[0], ": ")[1])

	maps := make([]Map, 0)
	curMap := Map{"", make([]int64,0), make([]int64,0), make([]int64,0)}
	
	for _, line := range lines[2:] {
		if strings.Contains(line, " map:") {
			curMap.title = line
			continue
		} else if strings.TrimSpace(line) == "" {
			maps = append(maps, curMap)
			curMap = Map{"", make([]int64,0), make([]int64,0), make([]int64,0)}
		} else {
			nums := wsNumsToArr(line)
			curMap.dest_starts = append(curMap.dest_starts, nums[0])
			curMap.src_starts = append(curMap.src_starts, nums[1])
			curMap.range_lens = append(curMap.range_lens, nums[2])
		}
	}
	maps = append(maps, curMap)
		
	return seeds, maps
}

func chain_lookup(seed int64, maps []Map) int64 {
	cur := seed
	var next int64 = -1
	//fmt.Print(seed)
	for _, m := range maps {
		next = m.lookup(cur)
		//fmt.Print(" -> ", next)
		cur = next
	}
	//fmt.Println()
	return next
}

func main() {
	lines := read_input("input/5")
	seeds, maps := deserialize(lines)
	lowest := int64(math.MaxInt64)
	for _, seed := range seeds {
		location := chain_lookup(seed, maps)
		if location < lowest {
			lowest = location
		}
	}
	
	fmt.Println("Lowest location is ", lowest)
	
}
