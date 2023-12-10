package main

import (
	"fmt"
	"slices"
)

type Direction int

const (
	North = Direction(iota)
	South
	East
	West
)

var Directions [4]Direction = [4]Direction{North, South, East, West}

func (d Direction) ToVec() Vec {
	switch d {
	case North:
		return Vec{0, -1}
	case South:
		return Vec{0, 1}
	case East:
		return Vec{1, 0}
	case West:
		return Vec{-1, 0}
	}
	panic(fmt.Sprintf("wrong direction %d", d))
}

func (d Direction) Opposite() Direction {
	switch d {
	case North:
		return South
	case South:
		return North
	case East:
		return West
	case West:
		return East
	}
	panic(fmt.Sprintf("wrong direction %d", d))
}

type Vec struct {
	X int
	Y int
}

func (v Vec) Add(other Vec) Vec {
	return Vec{v.X + other.X, v.Y + other.Y}
}

type Pipe struct {
	From Direction
	To   Direction
	repr rune
}

func PipeFromRune(r rune) *Pipe {
	switch r {
	case '|':
		return &Pipe{North, South, r}
	case '-':
		return &Pipe{East, West, r}
	case 'L':
		return &Pipe{North, East, r}
	case 'J':
		return &Pipe{North, West, r}
	case '7':
		return &Pipe{West, South, r}
	case 'F':
		return &Pipe{East, South, r}
	}
	return nil
}

// dir relative to p
// so we might need to reverse dir before calling this function
func (p Pipe) CanComeFrom(dir Direction) bool {
	return p.From == dir || p.To == dir
}

func (p Pipe) ExitDirection(comingFrom Direction) Direction {
	if comingFrom == p.From {
		return p.To
	} else {
		return p.From
	}
}

type Map struct {
	Tiles            [][]*Pipe
	StartingPosition Vec
}

func (m *Map) At(pos Vec) *Pipe {
	return m.Tiles[pos.Y][pos.X]
}

func (m *Map) Valid(pos Vec) bool {
	return pos.X >= 0 && pos.Y >= 0 && pos.Y < len(m.Tiles) && pos.X < len(m.Tiles[pos.Y])
}

func (m *Map) Traverse(position Vec, direction Direction) (newPosition Vec, newDirection Direction, ok bool) {
	newPosition = position.Add(direction.ToVec())
	if !m.Valid(newPosition) {
		return newPosition, direction, false
	}
	pipe := m.At(newPosition)
	if pipe == nil || !pipe.CanComeFrom(direction.Opposite()) {
		return newPosition, direction, false
	}
	return newPosition, pipe.ExitDirection(direction.Opposite()), true
}

func (m *Map) findLoop(pos Vec, dir Direction) *[]Vec {
	nextPos, nextDir, ok := m.Traverse(pos, dir)
	if nextPos == m.StartingPosition {
		return &[]Vec{pos}
	}
	if !ok {
		return nil
	}
	restOfTheLoop := m.findLoop(nextPos, nextDir)
	if restOfTheLoop == nil {
		return nil
	}
	res := append(*restOfTheLoop, pos)
	return &res
}
func (m *Map) FindLoop() []Vec {
	for _, dir := range Directions {
		loop := m.findLoop(m.StartingPosition, dir)
		if loop != nil {
			// return append([]Vec{m.StartingPosition}, (*loop)...)
			return *loop
		}
	}
	panic("no loop found in the map")
}

func (m *Map) IsInLoop(pos Vec, loop []Vec) bool {
	if slices.Contains(loop, pos) {
		return false
	}

	count := 0
	for i := pos.X; i < len(m.Tiles[pos.Y]); i++ {
		p := Vec{i, pos.Y}
		if slices.Contains(loop, p) {
			if p != m.StartingPosition {
				if m.At(p).repr == '|' || m.At(p).repr == 'L' || m.At(p).repr == 'J' {
					count += 1
				}
			}
		}
	}
	return count%2 == 1
}

func (m *Map) PrintTiles() {
	for y := range m.Tiles {
		for x := range m.Tiles[y] {
			if m.Tiles[y][x] != nil {
				fmt.Printf("%c", m.Tiles[y][x].repr)
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func parseMap() *Map {
	input := parseInput()
	res := &Map{}
	for y, line := range input {
		pipeLine := []*Pipe{}
		for x, r := range line {
			if r == 'S' {
				res.StartingPosition = Vec{x, y}
			}
			pipeLine = append(pipeLine, PipeFromRune(r))
		}
		res.Tiles = append(res.Tiles, pipeLine)
	}
	return res
}

func main() {
	m := parseMap()
	m.PrintTiles()
	loop := m.FindLoop()
	fmt.Println(loop)
	fmt.Println(len(loop) / 2)

	count := 0
	for y := range m.Tiles {
		for x := range m.Tiles[y] {
			if m.IsInLoop(Vec{x, y}, loop) {
				fmt.Printf("I")
				count += 1
			} else {
				fmt.Printf("O")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Println(count)
}
