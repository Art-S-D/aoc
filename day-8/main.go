package main

import (
	"bufio"
	"fmt"
	"os"
)

func isVisible(input [][]rune, x, y int) bool {
	height := input[x][y]

	// left
	for i := x - 1; i >= -1; i-- {
		if i == -1 {
			return true
		}
		if input[i][y] >= height {
			break
		}
	}
	// right
	for i := x + 1; i <= len(input); i++ {
		if i == len(input) {
			return true
		}
		if input[i][y] >= height {
			break
		}
	}
	// top
	for i := y - 1; i >= -1; i-- {
		if i == -1 {
			return true
		}
		if input[x][i] >= height {
			break
		}
	}
	// bottom
	for i := y + 1; i <= len(input[x]); i++ {
		if i == len(input[x]) {
			return true
		}
		if input[x][i] >= height {
			break
		}
	}

	return false
}

func scenicScore(input [][]rune, x, y int) int {
	height := input[x][y]

	left := 0
	for i := x - 1; i >= 0; i-- {
		if input[i][y] >= height {
			left += 1
			break
		} else {
			left += 1
		}
	}

	right := 0
	for i := x + 1; i < len(input); i++ {
		if input[i][y] >= height {
			right += 1
			break
		} else {
			right += 1
		}
	}

	top := 0
	for i := y - 1; i >= 0; i-- {
		if input[x][i] >= height {
			top += 1
			break
		} else {
			top += 1
		}
	}

	down := 0
	for i := y + 1; i < len(input[x]); i++ {
		if input[x][i] >= height {
			down += 1
			break
		} else {
			down += 1
		}
	}

	return left * right * top * down
}

func main() {
	var input [][]rune

	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	buf := bufio.NewScanner(inputFile)
	for buf.Scan() {
		line := buf.Text()
		input = append(input, []rune(line))
	}

	maxScenic := 0
	maxScenicX := 0
	maxScenicY := 0
	for x, row := range input {
		for y := range row {
			currentScenic := scenicScore(input, x, y)
			if currentScenic > maxScenic {
				maxScenic = currentScenic
				maxScenicX = x
				maxScenicY = y
			}
		}
	}
	fmt.Println(maxScenic)
	fmt.Printf("%d-%d\n", maxScenicX, maxScenicY)
}
