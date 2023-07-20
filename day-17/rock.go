package main

type Rock struct {
	Shape  []Vec2
	Width  uint8
	Height uint8
	Type   RockType
}

var rocks []Rock = []Rock{
	{Shape: []Vec2{{0, 0}, {1, 0}, {2, 0}, {3, 0}}, Width: 4, Height: 1, Type: HLine},        // -
	{Shape: []Vec2{{1, 0}, {0, 1}, {1, 1}, {2, 1}, {1, 2}}, Width: 3, Height: 3, Type: Plus}, // +
	{Shape: []Vec2{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}}, Width: 3, Height: 3, Type: L},    // L
	{Shape: []Vec2{{0, 0}, {0, 1}, {0, 2}, {0, 3}}, Width: 1, Height: 4, Type: VLine},        // |
	{Shape: []Vec2{{0, 0}, {0, 1}, {1, 0}, {1, 1}}, Width: 2, Height: 2, Type: Square},       // O
}

type RockType int

const (
	HLine = RockType(iota)
	Plus
	L
	VLine
	Square
)

func (r Rock) Clone() Rock {
	shape := make([]Vec2, len(r.Shape))
	copy(shape, r.Shape)
	return Rock{Shape: shape, Width: r.Width, Height: r.Height, Type: r.Type}
}

func NewRock(shape RockType) Rock {
	return rocks[shape].Clone()
}

func (r RockType) Next() RockType {
	return RockType((int(r) + 1) % len(rocks))
}
