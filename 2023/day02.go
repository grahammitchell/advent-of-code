package main

import (
	"bufio"
	"fmt"
	// "io/ioutil"
	//"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type RGB struct {
	red, green, blue, total int
}

type Game struct {
	games []RGB
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

func deserialize(lines []string) []Game {
	// Game 2: 7 red, 1 green, 4 blue; 13 red, 11 blue; 6 red, 2 blue; 9 blue, 9 red; 4 blue, 11 red; 15 red, 1 green, 3 blue
	r, _ := regexp.Compile("Game \\d+: (.*)$")

	games := make([]Game, 0)
	for _, line := range lines {
		// chop off "Game N: "
		parts := r.FindStringSubmatch(line)
		rest := parts[1]

		// split up successive draws
		draws := make([]RGB, 0)
		for _, draw := range strings.Split(rest, "; ") {
			// 7 red, 1 green, 4 blue
			arrgb := strings.Split(draw, ", ")
			rgb := splitDraw(arrgb)
			draws = append(draws, rgb)
		}
		games = append(games, Game{draws})
	}
	return games
}


func splitDraw(arrgb []string) RGB {
	vals := map[string]int{"red": 0, "green": 0, "blue": 0}
	for _, color := range arrgb {
		parts := strings.Split(color, " ")
		count := atoi(parts[0])
		vals[parts[1]] = count
	}
	total := vals["red"]+vals["green"]+vals["blue"]
	return RGB{vals["red"], vals["green"], vals["blue"], total}

}


func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func isIllegal(g Game) bool {
	cubeCap := RGB{12, 13, 14, 39}
	for _, rgb := range g.games {
		if rgb.red > cubeCap.red || rgb.green > cubeCap.green || rgb.blue > cubeCap.blue {
			return true
		}
	}
	return false
}


func main() {
	lines := read_input("input/2")
	games := deserialize(lines)
	sum := 0
	for i, game := range games {
		if ! isIllegal(game) {
			sum += i+1
		} else {
			fmt.Println("Illegal game # ", i+1)
		}
	}
	fmt.Println("Total: ", sum)
}
