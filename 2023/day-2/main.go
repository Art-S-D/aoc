package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CubeSet struct {
	Red   uint
	Green uint
	Blue  uint
}

func CubeSetFromColors(colors string) CubeSet {
	var res CubeSet
	for _, color := range strings.Split(colors, ", ") {
		splittedColor := strings.Split(color, " ")

		count, err := strconv.Atoi(splittedColor[0])
		if err != nil {
			panic(err.Error())
		}
		color := splittedColor[1]

		if color == "blue" {
			res.Blue += uint(count)
		} else if color == "red" {
			res.Red += uint(count)
		} else if color == "green" {
			res.Green += uint(count)
		} else {
			panic(fmt.Sprintf("unknown color %s", color))
		}
	}
	return res
}

func (cs CubeSet) Possible() bool {
	maxRed, maxGreen, maxBlue := uint(12), uint(13), uint(14)
	return cs.Red <= maxRed && cs.Green <= maxGreen && cs.Blue <= maxBlue
}

type Game []CubeSet

func (g *Game) Fewest() CubeSet {
	var res CubeSet
	for _, set := range *g {
		res.Red = max(res.Red, set.Red)
		res.Green = max(res.Green, set.Green)
		res.Blue = max(res.Blue, set.Blue)
	}
	return res
}
func (g *Game) Power() uint {
	fewest := g.Fewest()
	return fewest.Red * fewest.Green * fewest.Blue
}

func parseInput() []Game {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	res := []Game{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		gameStr := strings.Split(line, ": ")
		sets := strings.Split(gameStr[1], "; ")
		var game Game
		for _, set := range sets {
			game = append(game, CubeSetFromColors(set))
		}
		res = append(res, game)
	}
	return res
}

func main() {
	input := parseInput()

	result := uint(0)
	for _, game := range input {
		result += game.Power()
	}
	fmt.Println(result)
}
