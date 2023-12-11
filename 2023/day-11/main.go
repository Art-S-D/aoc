package main

import "fmt"

type Vec struct {
	X int
	Y int
}

func (v Vec) Add(other Vec) Vec {
	return Vec{v.X + other.X, v.Y + other.Y}
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
func (v Vec) Dist(other Vec) int {
	return abs(v.X-other.X) + abs(v.Y-other.Y)
}

type Map [][]rune

func (m Map) EmptyRow(row int) bool {
	for i := 0; i < len(m[row]); i++ {
		if m[row][i] != '.' {
			return false
		}
	}
	return true
}
func (m Map) EmptyCol(col int) bool {
	for i := 0; i < len(m); i++ {
		if m[i][col] != '.' {
			return false
		}
	}
	return true
}

func (m Map) GalaxiesPositions() []Vec {
	var galaxies []Vec
	for i := range m {
		for j := range m[i] {
			if m[i][j] == '#' {
				galaxies = append(galaxies, Vec{j, i})
			}
		}
	}
	offset := make([]Vec, len(galaxies))
	copy(offset, galaxies)
	for i := range m {
		if m.EmptyRow(i) {
			for v := range galaxies {
				if galaxies[v].Y > i {
					offset[v].Y += 1000000 - 1
				}
			}
		}
	}
	for i := range m[0] {
		if m.EmptyCol(i) {
			for v := range galaxies {
				if galaxies[v].X > i {
					offset[v].X += 1000000 - 1
				}
			}
		}
	}
	return offset
}

func sumOfDistances(galaxies []Vec) int {
	res := 0
	for i := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			// fmt.Println(i, j, i.Dist(j))
			res += galaxies[i].Dist(galaxies[j])
		}
	}
	return res
}

func parseMap() Map {
	var res Map
	for _, line := range parseInput() {
		res = append(res, []rune(line))
	}
	return res
}

func main() {
	m := parseMap()
	galaxies := m.GalaxiesPositions()
	fmt.Println(galaxies)
	fmt.Println(sumOfDistances(galaxies))
}
