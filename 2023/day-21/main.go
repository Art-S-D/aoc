package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Tile int

const (
	Rock = Tile(iota)
	GardenPlot
)

type Garden struct {
	Tiles            [][]Tile
	StartingPosition Vec
}

func ParseGardenMap() Garden {
	var res Garden
	for y, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		var l []Tile
		for x, c := range line {
			if c == 'S' {
				res.StartingPosition = Vec{x, y}
				l = append(l, GardenPlot)
			} else if c == '.' {
				l = append(l, GardenPlot)
			} else {
				l = append(l, Rock)
			}
		}
		res.Tiles = append(res.Tiles, l)
	}
	return res
}

func (g Garden) PlotsReachedAfter(steps int) map[Vec]bool {
	positions := make(map[Vec]bool)
	positions[g.StartingPosition] = true

	for i := 0; i < steps; i++ {
		nextPositions := make(map[Vec]bool)
		for position := range positions {
			for _, dir := range Directions {
				pos := position.Add(dir.ToVec())
				if pos.X < 0 ||
					pos.Y < 0 ||
					pos.Y >= len(g.Tiles) ||
					pos.X >= len(g.Tiles[pos.Y]) ||
					g.Tiles[pos.Y][pos.X] == Rock {
					continue
				}
				if _, ok := nextPositions[pos]; ok {
					continue
				}
				nextPositions[pos] = true
			}
		}
		positions = nextPositions
	}

	return positions
}

func (g Garden) PlotsReachedAfter2(steps int) int {
	positions := make(map[Vec]bool)
	positions[g.StartingPosition] = true

	cache := make(map[Vec]int)

	for i := 0; i < steps; i++ {
		// fmt.Println(i, len(positions), len(cache))
		nextPositions := make(map[Vec]bool)
		for position := range positions {
			for _, dir := range Directions {
				pos := position.Add(dir.ToVec())
				// if pos.X < 0 ||
				// 	pos.Y < 0 ||
				// 	pos.Y > len(g.Tiles) ||
				// 	pos.X > len(g.Tiles[pos.Y]) ||
				// 	g.Tiles[pos.Y][pos.X] == Rock ||
				// 	slices.Contains(nextPositions, pos) {
				// 	continue
				// }

				if _, ok := cache[pos]; ok {
					continue
				}

				mapPos := pos
				for mapPos.Y < 0 {
					mapPos.Y += len(g.Tiles)
				}
				mapPos.Y = mapPos.Y % len(g.Tiles)
				for mapPos.X < 0 {
					mapPos.X += len(g.Tiles[mapPos.Y])
				}
				mapPos.X = mapPos.X % len(g.Tiles[mapPos.Y])
				if g.Tiles[mapPos.Y][mapPos.X] == Rock {
					continue
				}
				nextPositions[pos] = true
				cache[pos] = i + 1
			}
		}
		positions = nextPositions
	}

	res := 0
	for _, i := range cache {
		if i%2 == steps%2 {
			res += 1
		}
	}
	return res
}

func main() {
	garden := ParseGardenMap()

	for i := 0; i < 5; i++ {
		count := garden.PlotsReachedAfter2(i*len(garden.Tiles) + 65)
		fmt.Printf("(%d,%d)\n", i, count)
	}
	// find quadratic equation that fits the output and solve for x=26501365 // 131 (or: steps // input length)
	// https://www.wolframalpha.com/input?i=quadratic+fit+calculator&assumption=%7B%22F%22%2C+%22QuadraticFitCalculator%22%2C+%22data3x%22%7D+-%3E%22%7B0%2C1%2C2%7D%22&assumption=%7B%22F%22%2C+%22QuadraticFitCalculator%22%2C+%22data3y%22%7D+-%3E%22%7B3859%2C+34324%2C+95135%7D%22
}
