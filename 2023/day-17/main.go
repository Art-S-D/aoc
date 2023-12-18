package main

import (
	"fmt"
	"strconv"
)

type Map struct {
	Tiles [][]int
	cache map[Crucible]int
}

func (m *Map) Inside(v Vec) bool {
	return v.X >= 0 && v.Y >= 0 && v.Y < len(m.Tiles) && v.X < len(m.Tiles[v.Y])
}

type Crucible struct {
	Pos           Vec
	Dir           Direction
	StraightMoves int
}

func ParseMap() Map {
	var res Map
	res.cache = make(map[Crucible]int)
	for _, line := range parseInput() {
		var l []int
		for _, c := range line {
			n, _ := strconv.Atoi(string(c))
			l = append(l, n)
		}
		res.Tiles = append(res.Tiles, l)
	}
	return res
}

// ordered to go to the bottom right first
var Dirs []Direction = []Direction{South, East, West, North}

func (m *Map) LowestHeatLost(crucible Crucible) int {
	if v, ok := m.cache[crucible]; ok {
		fmt.Println("cache hit")
		return v
	}
	if crucible.Pos.Y == len(m.Tiles)-1 && crucible.Pos.X == len(m.Tiles[crucible.Pos.Y])-1 {
		// reached the end
		return 0
	}
	localMin := -1
	for _, dir := range Dirs {
		if dir.Opposite() == crucible.Dir {
			// no turning back
			continue
		}
		if dir == crucible.Dir && crucible.StraightMoves == 3 {
			// no more than three times in a straight line
			continue
		}
		nextCrucible := Crucible{
			Pos:           crucible.Pos.Add(dir.ToVec()),
			Dir:           dir,
			StraightMoves: 1,
		}
		if dir == crucible.Dir {
			nextCrucible.StraightMoves = crucible.StraightMoves + 1
		}
		if !m.Inside(nextCrucible.Pos) {
			continue
		}
		lowest := m.Tiles[nextCrucible.Pos.Y][nextCrucible.Pos.X] + m.LowestHeatLost(nextCrucible)
		if localMin == -1 {
			localMin = lowest
		} else {
			localMin = min(localMin, lowest)
		}
	}
	m.cache[crucible] = localMin
	return localMin
}

func (m *Map) LowestHeatLost2(crucible Crucible, accumulatedHead int, localMin *int) {
	if v, ok := m.cache[crucible]; ok {
		if v <= accumulatedHead {
			return
		} else {
			m.cache[crucible] = accumulatedHead
		}
	} else {
		m.cache[crucible] = accumulatedHead
	}

	if *localMin != -1 && accumulatedHead >= *localMin {
		return
	}
	if crucible.Pos.Y == len(m.Tiles)-1 && crucible.Pos.X == len(m.Tiles[crucible.Pos.Y])-1 {
		// reached the end
		if *localMin == -1 {
			fmt.Println("initial minimum", accumulatedHead)
			*localMin = accumulatedHead
		} else {
			fmt.Println("new min", accumulatedHead)
			*localMin = min(*localMin, accumulatedHead)
		}
		return
	}
	for _, dir := range Dirs {
		if dir.Opposite() == crucible.Dir {
			// no turning back
			continue
		}
		if dir == crucible.Dir && crucible.StraightMoves == 3 {
			// no more than three times in a straight line
			continue
		}
		nextCrucible := Crucible{
			Pos:           crucible.Pos.Add(dir.ToVec()),
			Dir:           dir,
			StraightMoves: 1,
		}
		if dir == crucible.Dir {
			nextCrucible.StraightMoves = crucible.StraightMoves + 1
		}
		if !m.Inside(nextCrucible.Pos) {
			continue
		}
		m.LowestHeatLost2(nextCrucible, accumulatedHead+m.Tiles[nextCrucible.Pos.Y][nextCrucible.Pos.X], localMin)
	}
}

func (m *Map) LowestHeatLost3(crucible Crucible, accumulatedHead int, localMin *int) {
	if v, ok := m.cache[crucible]; ok {
		if v <= accumulatedHead {
			return
		} else {
			m.cache[crucible] = accumulatedHead
		}
	} else {
		m.cache[crucible] = accumulatedHead
	}

	if *localMin != -1 && accumulatedHead >= *localMin {
		return
	}
	if crucible.Pos.Y == len(m.Tiles)-1 && crucible.Pos.X == len(m.Tiles[crucible.Pos.Y])-1 {
		// reached the end
		if *localMin == -1 {
			fmt.Println("initial minimum", accumulatedHead)
			*localMin = accumulatedHead
		} else {
			fmt.Println("new min", accumulatedHead)
			*localMin = min(*localMin, accumulatedHead)
		}
		return
	}
	if crucible.StraightMoves < 4 {
		crucible.Pos = crucible.Pos.Add(crucible.Dir.ToVec())
		crucible.StraightMoves += 1
		if !m.Inside(crucible.Pos) {
			return
		}
		m.LowestHeatLost3(
			crucible,
			accumulatedHead+m.Tiles[crucible.Pos.Y][crucible.Pos.X],
			localMin,
		)
		return
	}
	for _, dir := range Dirs {
		if dir.Opposite() == crucible.Dir {
			// no turning back
			continue
		}
		if dir == crucible.Dir && crucible.StraightMoves >= 10 {
			// no more than 10 times in a straight line
			continue
		}
		nextCrucible := Crucible{
			Pos:           crucible.Pos.Add(dir.ToVec()),
			Dir:           dir,
			StraightMoves: 1,
		}
		if dir == crucible.Dir {
			nextCrucible.StraightMoves = crucible.StraightMoves + 1
		}
		if !m.Inside(nextCrucible.Pos) {
			continue
		}
		m.LowestHeatLost3(nextCrucible, accumulatedHead+m.Tiles[nextCrucible.Pos.Y][nextCrucible.Pos.X], localMin)
	}
}

func main() {
	cityMap := ParseMap()

	lowestHeat := -1
	cityMap.LowestHeatLost3(Crucible{Pos: Vec{0, 0}, Dir: East}, 0, &lowestHeat)

	// lowestHeat := cityMap.LowestHeatLost(Crucible{Pos: Vec{0, 0}, Dir: East})
	fmt.Println(lowestHeat)
}
