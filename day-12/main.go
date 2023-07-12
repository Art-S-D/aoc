package main

import (
	"fmt"

	"github.com/Art-S-D/aoc-2022-day-12/heightmap"
)

func main() {
	input, err := heightmap.FromFile("input.txt")
	if err != nil {
		panic(err.Error())
	}

	dijkstra := input.HeightMap.Dijkstra(input.Destination)

	shortestPath := dijkstra[input.Start.Y][input.Start.X]

	for y, line := range input.HeightMap {
		for x, square := range line {
			if square == 'a' && dijkstra[y][x] != nil && len(dijkstra[y][x]) < len(shortestPath) {
				shortestPath = dijkstra[y][x]
			}
		}
	}
	fmt.Println(len(shortestPath))
}
