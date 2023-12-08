package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type LR struct {
	left, right string
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

func deserialize(lines []string) (string, map[string]LR) {

	lr := lines[0]

	nodes := make(map[string]LR)
	
	for _, line := range lines[2:] {
		// AAA = (BBB, CCC)
		re := regexp.MustCompile(`^(.*) = \((.*), (.*)\)$`)
		matches := re.FindStringSubmatch(line)
		label := matches[1]
		nodes[label] = LR{matches[2], matches[3]}
	}
		
	return lr, nodes
}

func navigate(lr string, nodes map[string]LR) int {
	steps := 0
	cur := "AAA"
	fmt.Print(cur)
	for {
		for _, c := range lr {
			steps++

			if c == 'L' {
				cur = nodes[cur].left
			} else {
				cur = nodes[cur].right
			}
			if cur == "ZZZ" {
				return steps
			}
			fmt.Print(" -> ", cur)
		}
	}
	return steps
}

func navigate_multi(lr string, nodes map[string]LR) int {
	steps := 0
	curs := make([]string,0)
	// find all the keys ending in A
	for k, _ := range nodes {
		if strings.HasSuffix(k, "A") {
			curs = append(curs, k)
		}
	}
	fmt.Println(curs)
	for {
		for _, c := range lr {
			steps++
			
			if c == 'L' {
				for i, cur := range curs {
					curs[i] = nodes[cur].left
				}
			} else {
				for i, cur := range curs {
					curs[i] = nodes[cur].right
				}
			}
			zeds := 0
			for _, cur := range curs {
				if strings.HasSuffix(cur, "Z") {
					zeds++
				}
			}
			if zeds == len(curs) {
				return steps
			}
		}
	}
	return steps
}


func main() {
	lines := read_input("input/8")
	lr, nodes := deserialize(lines)
	
	fmt.Println(lr)
	for k,v := range nodes {
		fmt.Printf("%s = (%s, %s)\n", k, v.left, v.right)
	}
	steps := 0
	steps = navigate_multi(lr, nodes)
	fmt.Println("\nSteps: ", steps)
	
}
