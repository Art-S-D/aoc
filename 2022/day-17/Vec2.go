package main

type Vec2 struct {
	X, Y int
}

func absInt(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

func (v Vec2) Dist(other Vec2) int {
	return absInt(v.X-other.X) + absInt(v.Y-other.Y)
}

func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{v.X + other.X, v.Y + other.Y}
}
