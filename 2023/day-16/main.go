package main

import (
	"fmt"
)

type Contraption struct {
	Grid      []string
	Energized map[Vec]bool
	Beams     []Beam
	cache     map[Beam]bool
}

type Beam struct {
	Pos       Vec
	Direction Direction
}

func ParseContraption() *Contraption {
	initialBeam := Beam{Pos: Vec{-1, 0}, Direction: East}
	return &Contraption{
		Grid:      parseInput(),
		Energized: make(map[Vec]bool),
		Beams:     []Beam{initialBeam},
		cache:     make(map[Beam]bool),
	}
}

func (c *Contraption) Reset() {
	c.Energized = make(map[Vec]bool)
	c.Beams = []Beam{Beam{Pos: Vec{-1, 0}, Direction: East}}
	c.cache = make(map[Beam]bool)
}

func (d Direction) Reflect(mirror byte) Direction {
	if mirror == '/' {
		if d == North {
			return East
		} else if d == West {
			return South
		} else if d == South {
			return West
		} else if d == East {
			return North
		}
	} else if mirror == '\\' {
		if d == North {
			return West
		} else if d == West {
			return North
		} else if d == South {
			return East
		} else if d == East {
			return South
		}
	}
	panic(fmt.Sprintf("unknown mirror %c\n", mirror))
}
func (d Direction) Split(splitter byte) []Direction {
	if splitter == '-' {
		if d == West || d == East {
			return []Direction{d}
		} else {
			return []Direction{East, West}
		}
	} else if splitter == '|' {
		if d == North || d == South {
			return []Direction{d}
		} else {
			return []Direction{North, South}
		}
	}
	panic(fmt.Sprintf("unknown splitter %c\n", splitter))
}

func (c *Contraption) ProcessBeamOnce(id int) {
	beam := &c.Beams[id]
	nextPos := beam.Pos.Add(beam.Direction.ToVec())
	if c.cache[*beam] || nextPos.X < 0 || nextPos.Y < 0 || nextPos.Y >= len(c.Grid) || nextPos.X >= len(c.Grid[nextPos.Y]) {
		c.Beams[id] = c.Beams[len(c.Beams)-1]
		c.Beams = c.Beams[:len(c.Beams)-1]
		return
	}

	c.cache[*beam] = true
	c.Energized[nextPos] = true
	beam.Pos = nextPos

	tile := c.Grid[nextPos.Y][nextPos.X]

	if tile == '/' || tile == '\\' {
		beam.Direction = beam.Direction.Reflect(tile)
	}
	if tile == '-' || tile == '|' {
		dirs := beam.Direction.Split(tile)
		beam.Direction = dirs[0]
		if len(dirs) == 2 {
			c.Beams = append(c.Beams, Beam{Pos: nextPos, Direction: dirs[1]})
		}
	}
}

func (c *Contraption) ProcessAllBeams() int {
	for len(c.Beams) > 0 {
		c.ProcessBeamOnce(0)
	}
	return len(c.Energized)
}

func main() {
	contraption := ParseContraption()

	maxEnergy := 0
	for i := range contraption.Grid[0] {
		contraption.Reset()
		contraption.Beams = []Beam{{Pos: Vec{i, -1}, Direction: South}}
		maxEnergy = max(maxEnergy, contraption.ProcessAllBeams())
	}
	for i := range contraption.Grid[len(contraption.Grid)-1] {
		contraption.Reset()
		contraption.Beams = []Beam{{Pos: Vec{i, len(contraption.Grid)}, Direction: North}}
		maxEnergy = max(maxEnergy, contraption.ProcessAllBeams())
	}
	for i := range contraption.Grid {
		contraption.Reset()
		contraption.Beams = []Beam{{Pos: Vec{-1, i}, Direction: East}}
		maxEnergy = max(maxEnergy, contraption.ProcessAllBeams())
	}
	for i := range contraption.Grid {
		contraption.Reset()
		contraption.Beams = []Beam{{Pos: Vec{len(contraption.Grid[i]), i}, Direction: West}}
		maxEnergy = max(maxEnergy, contraption.ProcessAllBeams())
	}

	fmt.Println(maxEnergy)
}
