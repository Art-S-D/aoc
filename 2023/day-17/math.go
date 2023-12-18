package main

import "fmt"

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

func (d Direction) String() string {
	return []string{"North", "South", "East", "West"}[d]
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
