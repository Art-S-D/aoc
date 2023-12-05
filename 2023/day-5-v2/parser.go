package main

import (
	"strconv"
	"strings"
)

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
		res.Seeds = append(res.Seeds, Range{Start: uint(start), Length: uint(len)})
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
				MappingRange{Range: Range{Start: uint(src), Length: uint(len)}, Dest: uint(dest)},
			)
			i++
		}

		res.Maps = append(res.Maps, currentMap)
		currentMap = CategoryMap{}
	}

	return res
}
