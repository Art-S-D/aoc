package main

import (
	"fmt"

	"github.com/Art-S-D/aoc-2022-day-14/cave"
)

func main() {
	scan := cave.Parse()

	count := 0
	for scan.At(500, 0) == cave.Air {
		scan.SpawnSand()
		count += 1
	}
	// scan.Debug(490, 510, 0, 12)
	scan.Debug(450, 550, 0, 60)
	fmt.Println(count)
}
