package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type Range struct {
	Src    uint
	Dest   uint
	Length uint
}

func (r *Range) Contains(src uint) bool {
	return r.Src <= src && r.Src+r.Length > src
}

// assumes r contains src
func (r *Range) Map(src uint) (dest uint) {
	return (src - r.Src) + r.Dest
}

type CategoryMap struct {
	Src    string
	Dest   string
	Ranges []Range
}

func (m *CategoryMap) Map(src uint) (dest uint) {
	for _, r := range m.Ranges {
		if r.Contains(src) {
			return r.Map(src)
		}
	}
	return src
}

type SeedRange struct {
	Start  uint
	Length uint
}
type Almanach struct {
	Seeds []SeedRange
	Maps  []CategoryMap
}

func (a *Almanach) SeedLocation(seed uint) uint {
	currentCategory := "seed"
	currentValue := seed

	for currentCategory != "location" {
		for _, categoryMap := range a.Maps {
			if categoryMap.Src == currentCategory {
				currentCategory = categoryMap.Dest
				currentValue = categoryMap.Map(currentValue)
			}
		}
	}
	return currentValue
}

func (almanach *Almanach) MinimumLocationForSeedRange(seed SeedRange) uint {
	min := almanach.SeedLocation(seed.Start)
	for i := seed.Start + 1; i < seed.Start+seed.Length; i++ {
		location := almanach.SeedLocation(i)
		if location < min {
			min = location
		}
	}
	return min
}

func (almanach *Almanach) MinimumLocation() uint {
	var wg sync.WaitGroup
	minimums := make(chan uint, len(almanach.Seeds))

	for _, seed := range almanach.Seeds {
		wg.Add(1)
		go func(almanach *Almanach, seed SeedRange) {
			minimums <- almanach.MinimumLocationForSeedRange(seed)
			wg.Done()
		}(almanach, seed)
	}
	wg.Wait()
	min := <-minimums
	for len(minimums) > 0 {
		m := <-minimums
		if m < min {
			min = m
		}
	}
	return min
}

func AlmanacFromString(lines []string) Almanach {
	res := Almanach{}
	seeds := strings.Split(lines[0][7:], " ")
	for i := 0; i < len(seeds); i += 2 {
		start, err := strconv.Atoi(seeds[i])
		if err != nil {
			panic(err.Error())
		}
		len, err := strconv.Atoi(seeds[i+1])
		if err != nil {
			panic(err.Error())
		}
		res.Seeds = append(res.Seeds, SeedRange{Start: uint(start), Length: uint(len)})
	}

	currentMap := CategoryMap{}
	for i := 2; i < len(lines); i++ {
		mapName := strings.Split(lines[i], " ")[0]
		categoryNames := strings.Split(mapName, "-to-")
		currentMap.Src = categoryNames[0]
		currentMap.Dest = categoryNames[1]

		// skip map title line
		i++

		for i < len(lines) && len(lines[i]) > 0 {
			rangeValues := strings.Split(lines[i], " ")
			dest, err := strconv.Atoi(rangeValues[0])
			if err != nil {
				panic(err.Error())
			}
			src, err := strconv.Atoi(rangeValues[1])
			if err != nil {
				panic(err.Error())
			}
			len, err := strconv.Atoi(rangeValues[2])
			if err != nil {
				panic(err.Error())
			}
			currentMap.Ranges = append(currentMap.Ranges,
				Range{
					Src:    uint(src),
					Dest:   uint(dest),
					Length: uint(len),
				},
			)
			i++
		}

		res.Maps = append(res.Maps, currentMap)
		currentMap = CategoryMap{}
	}

	return res
}

func main() {
	input := parseInput()
	almanach := AlmanacFromString(input)
	fmt.Println(almanach.MinimumLocation())
}
