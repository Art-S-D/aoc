package main

import "fmt"

type Direction int

const (
	Up = Direction(iota)
	Down
	Right
	Left
)

var Directions [4]Direction = [4]Direction{Right, Down, Left, Up}

func (d Direction) ToVec() Vec {
	switch d {
	case Up:
		return Vec{0, -1}
	case Down:
		return Vec{0, 1}
	case Right:
		return Vec{1, 0}
	case Left:
		return Vec{-1, 0}
	}
	panic(fmt.Sprintf("wrong direction %d", d))
}

func (d Direction) String() string {
	return []string{"Up", "Down", "Right", "Left"}[d]
}

func DirectionFromByte(b byte) Direction {
	switch b {
	case 'U':
		return Up
	case 'D':
		return Down
	case 'L':
		return Left
	case 'R':
		return Right
	}
	panic(fmt.Sprintf("unknown direction %c\n", b))
}

func (d Direction) Opposite() Direction {
	switch d {
	case Up:
		return Down
	case Down:
		return Up
	case Left:
		return Right
	case Right:
		return Left
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
func (v Vec) Scale(by int) Vec {
	return Vec{v.X * by, v.Y * by}
}
