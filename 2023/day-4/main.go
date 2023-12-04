package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Card struct {
	Id      uint
	Winning []uint
	Has     []uint
}

func (c *Card) Worth() uint {
	winning := make(map[uint]bool)
	for _, number := range c.Winning {
		winning[number] = true
	}
	value := uint(0)
	for _, number := range c.Has {
		if _, ok := winning[number]; ok {
			if value == 0 {
				value = 1
			} else {
				value *= 2
			}
		}
	}
	return value
}

func (c *Card) Matching() []uint {
	winning := make(map[uint]bool)
	for _, number := range c.Winning {
		winning[number] = true
	}
	var res []uint
	for _, number := range c.Has {
		if _, ok := winning[number]; ok {
			res = append(res, number)
		}
	}
	return res
}

func ParseCards() []Card {
	input := parseInput()
	res := []Card{}
	for _, line := range input {
		card := Card{}

		cardSplit := strings.Split(line, ": ")
		cardId, err := strconv.Atoi(strings.Trim(cardSplit[0][4:], " "))
		if err != nil {
			panic(err.Error())
		}
		card.Id = uint(cardId)
		cardDescription := cardSplit[1]
		splittedCard := strings.Split(cardDescription, " | ")
		winningNumbers := splittedCard[0]
		pickedNumbers := splittedCard[1]

		for i := 0; i < len(winningNumbers); i += 3 {
			n, err := strconv.Atoi(strings.Trim(winningNumbers[i:i+2], " "))
			if err != nil {
				panic(err.Error())
			}
			card.Winning = append(card.Winning, uint(n))
		}
		for i := 0; i < len(pickedNumbers); i += 3 {
			n, err := strconv.Atoi(strings.Trim(pickedNumbers[i:i+2], " "))
			if err != nil {
				panic(err.Error())
			}
			card.Has = append(card.Has, uint(n))
		}
		res = append(res, card)
	}
	return res
}

func main() {
	initialCards := ParseCards()
	// totalWinnings := uint(0)
	// for _, card := range cards {
	// 	totalWinnings += card.Worth()
	// }
	// fmt.Println(totalWinnings)

	count := 0

	currentCards := initialCards

	var nextCards []Card

	for len(currentCards) > 0 {
		for _, card := range currentCards {
			count += 1
			matches := card.Matching()
			for i := 0; i < len(matches); i++ {
				nextCards = append(nextCards, initialCards[int(card.Id)+i])
			}
		}
		currentCards = nextCards
		nextCards = []Card{}
	}

	fmt.Println(count)
}
