package main

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Card rune

var cardValues = map[Card]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	// 'J': 11,
	'J': 1,
	'Q': 12,
	'K': 13,
	'A': 14,
}

func (c Card) Value() int {
	v, ok := cardValues[c]
	if !ok {
		panic(fmt.Sprintf("unknown card %v", v))
	}
	return v
}

func (c Card) String() string {
	return string(c)
}

type Hand [5]Card

func (h Hand) CountCards() map[Card]uint {
	res := make(map[Card]uint)
	for _, card := range h {
		res[card] += 1
	}
	return res
}

func (h Hand) CompareCards(other Hand) int {
	return slices.CompareFunc[[]Card, []Card, Card, Card]([]Card(h[:]), []Card(other[:]), func(c1, c2 Card) int {
		return cmp.Compare(c1.Value(), c2.Value())
	})
}

func HandFromString(in string) Hand {
	if len(in) != 5 {
		panic(fmt.Sprintf("tried to create a hand with invalid string %v", in))
	}
	var res Hand
	for i := range res {
		if _, ok := cardValues[Card(in[i])]; !ok {
			panic(fmt.Sprintf("invalid card %v", in[i]))
		}
		res[i] = Card(in[i])
	}
	return res
}

func (h Hand) String() string {
	res := ""
	for _, v := range h {
		res += v.String()
	}
	return res
}

type HandType int

const (
	HighCard = HandType(iota)
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func (t HandType) String() string {
	switch t {
	case HighCard:
		return "High-Card"
	case OnePair:
		return "One-Pair"
	case TwoPair:
		return "Two-Pairs"
	case ThreeOfAKind:
		return "Three-of-a-Kind"
	case FullHouse:
		return "Full-House"
	case FourOfAKind:
		return "Four-of-a-Kind"
	case FiveOfAKind:
		return "Five-of-a-Kind"
	}
	panic("wrong type")
}

func (h Hand) Type() HandType {
	count := h.CountCards()
	var sortedCounts []uint
	for _, v := range count {
		sortedCounts = append(sortedCounts, v)
	}
	slices.Sort(sortedCounts)
	slices.Reverse(sortedCounts)

	if sortedCounts[0] == 5 {
		return FiveOfAKind
	} else if sortedCounts[0] == 4 && sortedCounts[1] == 1 {
		return FourOfAKind
	} else if sortedCounts[0] == 3 && sortedCounts[1] == 2 {
		return FullHouse
	} else if sortedCounts[0] == 3 {
		return ThreeOfAKind
	} else if sortedCounts[0] == 2 && sortedCounts[1] == 2 {
		return TwoPair
	} else if sortedCounts[0] == 2 {
		return OnePair
	} else {
		return HighCard
	}
}

func (h Hand) Compare(other Hand) int {
	hType := h.Jokered().Type()
	otherType := other.Jokered().Type()

	// fmt.Printf("%v - %v / %v - %v\n", h, hType, other, otherType)

	if hType == otherType {
		return h.CompareCards(other)
	} else {
		return cmp.Compare(hType, otherType)
	}
}

func (h Hand) Jokered() Hand {
	counts := h.CountCards()

	// find the most frequent card in h
	// jokers will take the value of this card
	// if none appear more than one time, all jokers will become an ace
	mostFrequentCard := h[0]
	mostFrequestCardCount := uint(1)
	for k, v := range counts {
		if k == 'J' {
			continue
		}
		if v > mostFrequestCardCount {
			mostFrequentCard = k
			mostFrequestCardCount = v
		} else if v == mostFrequestCardCount && k.Value() > mostFrequentCard.Value() {
			mostFrequentCard = k
			mostFrequestCardCount = v
		}
	}
	if mostFrequentCard == 'J' {
		mostFrequentCard = 'A'
	}
	var res Hand
	for i := range h {
		if h[i] != 'J' {
			res[i] = h[i]
		} else {
			res[i] = mostFrequentCard
		}
	}
	return res
}

type HandBid struct {
	Hand Hand
	Bid  uint
}

func parseHandBids() []HandBid {
	input := parseInput()
	var res []HandBid
	for _, line := range input {
		fields := strings.Fields(line)
		hand := HandFromString(fields[0])
		bid, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err.Error())
		}
		res = append(res, HandBid{Hand: hand, Bid: uint(bid)})
	}
	return res
}

func main() {
	handBids := parseHandBids()

	for _, h := range handBids {
		fmt.Println(h.Hand, h.Hand.Jokered(), h.Hand.Jokered().Type())
	}

	slices.SortFunc(handBids, func(h1, h2 HandBid) int {
		return h1.Hand.Compare(h2.Hand)
	})
	slices.Reverse(handBids)

	winnings := uint(0)
	for i, hand := range handBids {
		winnings += hand.Bid * (uint(len(handBids)) - uint(i))
	}
	fmt.Println(winnings)
}
