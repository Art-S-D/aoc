package main

import (
	"fmt"
	"os"
)

type FallingRock struct {
	Rock
	Pos Vec2
}

type VerticalChamber struct {
	FallingRock  *FallingRock
	Chamber      map[Vec2]bool
	Width        uint8
	highestChunk uint
	jetPattern   []Jet
	currentJet   int
}

type Jet rune

const (
	Left  = Jet('<')
	Right = Jet('>')
)

func readJetFn() func() []Jet {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err.Error())
	}
	jetPattern := []Jet(string(input))
	return func() []Jet {
		return jetPattern
	}
}

var ReadJet = readJetFn()

func NewChamber() *VerticalChamber {
	return &VerticalChamber{
		FallingRock:  nil,
		Chamber:      make(map[Vec2]bool),
		Width:        7,
		highestChunk: 0,
		jetPattern:   ReadJet(),
		currentJet:   0,
	}
}

func (chamber *VerticalChamber) Clear() {
	chamber.Chamber = make(map[Vec2]bool)
}

func (chamber *VerticalChamber) IsValid(rock FallingRock) bool {
	if rock.Pos.X < 0 || rock.Pos.Y < 0 {
		return false
	}
	if rock.Pos.X+int(rock.Width) > int(chamber.Width) {
		return false
	}

	// check for intersection with other rocks
	for _, v := range rock.Shape {
		chunk := rock.Pos.Add(v)
		if chamber.Chamber[chunk] {
			return false
		}
	}

	return true
}

func (chamber *VerticalChamber) SpawnNextRock() {
	var newType RockType
	if chamber.FallingRock == nil {
		newType = HLine
	} else {
		newType = chamber.FallingRock.Type.Next()
	}
	rock := NewRock(newType)
	chamber.FallingRock = &FallingRock{rock, Vec2{2, int(chamber.highestChunk) + 3}}
}

func (chamber *VerticalChamber) SaveRock() {
	for _, v := range chamber.FallingRock.Shape {
		chunk := chamber.FallingRock.Pos.Add(v)
		chamber.Chamber[chunk] = true
		chunkTop := chunk.Y + 1
		if chunkTop > int(chamber.highestChunk) {
			chamber.highestChunk = uint(chunkTop)
		}
	}
}

func (chamber *VerticalChamber) JetPush() {
	jet := chamber.jetPattern[chamber.currentJet]
	chamber.currentJet = (chamber.currentJet + 1) % len(chamber.jetPattern)

	var dir Vec2
	if jet == Left {
		dir = Vec2{-1, 0}
	} else if jet == Right {
		dir = Vec2{1, 0}
	} else {
		panic(fmt.Sprintf("unknown pattern %c", jet))
	}

	previousPos := chamber.FallingRock.Pos
	chamber.FallingRock.Pos = chamber.FallingRock.Pos.Add(dir)
	if !chamber.IsValid(*chamber.FallingRock) {
		chamber.FallingRock.Pos = previousPos
	}
}

func (chamber *VerticalChamber) RockFall() {
	chamber.SpawnNextRock()
	for {
		chamber.JetPush()

		previousPos := chamber.FallingRock.Pos
		chamber.FallingRock.Pos.Y -= 1
		if !chamber.IsValid(*chamber.FallingRock) {
			chamber.FallingRock.Pos = previousPos
			chamber.SaveRock()
			return
		}
	}
}

func (chamber *VerticalChamber) Debug() {
	for y := int(chamber.highestChunk); y >= 0; y-- {
		fmt.Print("|")
		for x := 0; x < int(chamber.Width); x++ {
			if chamber.Chamber[Vec2{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("|")
		fmt.Println()
	}
	fmt.Println("+-------+")
}
