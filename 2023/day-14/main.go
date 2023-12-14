package main

import (
	"fmt"
)

type Platform [][]rune

func ParsePlatform() Platform {
	var res Platform
	for _, line := range parseInput() {
		res = append(res, []rune(line))
	}
	return res
}

func (p Platform) TiltNorth() {
	for y := range p {
		for x := range p[y] {
			j := y
			for j > 0 && p[j][x] == 'O' && p[j-1][x] == '.' {
				p[j][x] = '.'
				p[j-1][x] = 'O'
				j--
			}
		}
	}
}

func (p Platform) TiltSouth() {
	for y := len(p) - 1; y >= 0; y-- {
		for x := range p[y] {
			j := y
			for j < len(p)-1 && p[j][x] == 'O' && p[j+1][x] == '.' {
				p[j][x] = '.'
				p[j+1][x] = 'O'
				j++
			}
		}
	}
}

func (p Platform) TiltWest() {
	for y := range p {
		for x := range p[y] {
			j := x
			for j > 0 && p[y][j] == 'O' && p[y][j-1] == '.' {
				p[y][j] = '.'
				p[y][j-1] = 'O'
				j--
			}
		}
	}
}

func (p Platform) TiltEast() {
	for y := range p {
		for x := len(p[y]) - 1; x >= 0; x-- {
			j := x
			for j < len(p[y])-1 && p[y][j] == 'O' && p[y][j+1] == '.' {
				p[y][j] = '.'
				p[y][j+1] = 'O'
				j++
			}
		}
	}
}

func (p Platform) NorthLoad() int {
	res := 0
	for y := range p {
		for x := range p[y] {
			if p[y][x] == 'O' {
				res += len(p) - y
			}
		}
	}
	return res
}

func (p Platform) String() string {
	var res string
	for _, l := range p {
		res += string(l) + "\n"
	}
	return res
}

func main() {
	platform := ParsePlatform()

	cache := make(map[string]int)

	iterations := 1_000_000_000
	for i := 0; i < iterations; i++ {
		platform.TiltNorth()
		platform.TiltWest()
		platform.TiltSouth()
		platform.TiltEast()

		if count, ok := cache[platform.String()]; ok {
			fmt.Printf("found loop from %d to %d iterations\n", count, i)
			loopSize := i - count
			for i+loopSize < iterations {
				i += loopSize
			}
		}
		cache[platform.String()] = i
	}

	fmt.Println(platform.NorthLoad())
}
