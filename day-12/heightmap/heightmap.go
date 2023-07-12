package heightmap

import (
	"bufio"
	"os"
)

type HeightMap [][]rune
type Path []Vec2
type Vec2 struct {
	X int
	Y int
}

func (l Vec2) Add(other Vec2) Vec2 {
	return Vec2{X: l.X + other.X, Y: l.Y + other.Y}
}

var orthogonalNeighbors = []Vec2{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func (m HeightMap) Dijkstra(destination Vec2) [][]Path {
	res := make([][]Path, len(m))
	for i := range res {
		res[i] = make([]Path, len(m[i]))
	}

	res[destination.Y][destination.X] = Path{}
	toProcess := []Vec2{destination}

	for len(toProcess) > 0 {
		currentLocation := toProcess[0]
		currentLocationPath := res[currentLocation.Y][currentLocation.X]
		toProcess = toProcess[1:]

		for _, v := range orthogonalNeighbors {
			neighbor := currentLocation.Add(v)
			if neighbor.X < 0 || neighbor.Y < 0 || neighbor.Y >= len(m) || neighbor.X >= len(m[neighbor.Y]) {
				// neighbor is outside the map
				continue
			}
			if m[currentLocation.Y][currentLocation.X] > m[neighbor.Y][neighbor.X]+1 {
				// currentLocation is too high, can't climb to it from neighbor
				// remember that we are moving from the destination so it's reversed
				continue
			}

			neighborPath := res[neighbor.Y][neighbor.X]

			if neighborPath != nil && len(neighborPath) <= len(currentLocationPath)+1 {
				// there is already a shorter path to neighbor
				continue
			}

			res[neighbor.Y][neighbor.X] = make(Path, len(currentLocationPath)+1)
			res[neighbor.Y][neighbor.X][0] = currentLocation
			copy(res[neighbor.Y][neighbor.X][1:], currentLocationPath)
			toProcess = append(toProcess, neighbor)
		}
	}

	return res
}

func (m HeightMap) ShortestPath(from, to Vec2) Path {
	dijkstra := m.Dijkstra(to)
	return dijkstra[from.Y][from.X]
}

type Input struct {
	HeightMap   HeightMap
	Start       Vec2
	Destination Vec2
}

func FromFile(filename string) (Input, error) {
	var res Input

	inputFile, err := os.Open(filename)
	if err != nil {
		return res, err
	}
	scanner := bufio.NewScanner(inputFile)

	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		res.HeightMap = append(res.HeightMap, []rune(line))
		for x, r := range line {
			if r == 'S' {
				res.Start = Vec2{X: x, Y: y}
				res.HeightMap[y][x] = 'a'
			}
			if r == 'E' {
				res.Destination = Vec2{X: x, Y: y}
				res.HeightMap[y][x] = 'z'
			}
		}
	}

	return res, nil
}
