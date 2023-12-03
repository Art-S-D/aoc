package main

import (
	"fmt"
	"unicode"
)

func isSymbol(r rune) bool {
	isNumber := unicode.IsDigit(r)
	return !isNumber && r != '.'
}

type Schematic [][]rune

func SchematicFromStringArray(in []string) Schematic {
	res := Schematic{}
	for _, line := range in {
		res = append(res, []rune(line))
	}
	return res
}

// PART ONE

func (s Schematic) IsAdjacentToSymbol(x, y int) bool {
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i < 0 || j < 0 || j >= len(s) || i >= len(s[j]) {
				continue
			}
			current := s[j][i]
			if isSymbol(current) {
				return true
			}
		}
	}
	return false
}

func (s Schematic) ProcessNumber(x, y int) (value, length int, adjacentToSymbol bool) {
	for i := x; i < len(s[y]) && unicode.IsDigit(s[y][i]); i++ {
		length++
		value = value*10 + int(s[y][i]-'0')
		if s.IsAdjacentToSymbol(i, y) {
			adjacentToSymbol = true
		}
	}
	return value, length, adjacentToSymbol
}

func (s Schematic) PartsSum() (sum int) {
	for y := range s {
		for x := 0; x < len(s[y]); x++ {
			if unicode.IsDigit(s[y][x]) {
				value, length, adjacentToSymbol := s.ProcessNumber(x, y)
				if adjacentToSymbol {
					sum += value
				}
				x += length // ?
			}
		}
	}
	return sum
}

// PART TWO

func (s Schematic) NumberAt(x, y int) int {
	i := x
	for i >= 0 && unicode.IsDigit(s[y][i]) {
		i--
	}
	// we end up with i to the left of the number so we need to shift one to the right
	value, _, _ := s.ProcessNumber(i+1, y)
	return value
}

func Unique(intSlice []int) []int {
	keys := make(map[int]bool)
	res := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			res = append(res, entry)
		}
	}
	return res
}

func (s Schematic) NumbersAround(x, y int) []int {
	res := []int{}
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i < 0 || j < 0 || j >= len(s) || i >= len(s[j]) {
				continue
			}
			if unicode.IsDigit(s[j][i]) {
				res = append(res, s.NumberAt(i, j))
			}
		}
	}
	return Unique(res)
}

func (s Schematic) GearRatioSum() int {
	res := 0
	for y := range s {
		for x := range s[y] {
			if s[y][x] == '*' {
				ratios := s.NumbersAround(x, y)
				if len(ratios) != 2 {
					fmt.Printf("gear with %d numbers adjacent to it ar (%d,%d)\n", len(ratios), x, y)
				}
				if len(ratios) == 2 {
					res += ratios[0] * ratios[1]
				}
			}
		}
	}
	return res
}

func main() {
	input := parseInput()
	schematic := SchematicFromStringArray(input)
	// sum := schematic.PartsSum()
	sum := schematic.GearRatioSum()
	fmt.Println(sum)
}
