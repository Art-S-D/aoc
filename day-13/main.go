package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const (
	Integer = 0
	List    = 1
)

type Packet struct {
	Type    int
	Integer int
	List    []*Packet
}

func (packet Packet) String() string {
	if packet.Type == Integer {
		return fmt.Sprintf("%d", packet.Integer)
	} else {
		var sb strings.Builder
		sb.WriteRune('[')
		for i, p := range packet.List {
			sb.WriteString(p.String())
			if i < len(packet.List)-1 {
				sb.WriteRune(',')
			}
		}
		sb.WriteRune(']')
		return sb.String()
	}
}

func cmp(i, j int) int {
	if i < j {
		return -1
	} else if i > j {
		return 1
	} else {
		return 0
	}
}
func (packet *Packet) Cmp(other *Packet) int {
	if packet.Type == Integer && other.Type == Integer {
		return cmp(packet.Integer, other.Integer)
	} else if packet.Type == List && other.Type == List {
		for i := range packet.List {
			if i >= len(other.List) {
				return 1
			}
			left := packet.List[i]
			right := other.List[i]

			cmp := left.Cmp(right)
			if cmp != 0 {
				return cmp
			}
		}
		if len(other.List) == len(packet.List) {
			return 0
		} else {
			return -1
		}
	} else if packet.Type == Integer {
		list := Packet{Type: List, List: []*Packet{packet}}
		return list.Cmp(other)
	} else if other.Type == Integer {
		list := Packet{Type: List, List: []*Packet{other}}
		return packet.Cmp(&list)
	} else {
		panic(fmt.Sprintf("unknown types: %d, %d\n", packet.Type, other.Type))
	}
}

// type PacketPair struct {
// 	Left, Right Packet
// }

// func (pair *PacketPair) IsCorrectlyOrdered() bool {
// 	return pair.Left.Cmp(&pair.Right) < 0
// }

// assumes line has the ScanRunes split and has been scanned to the beginning of the current packet
func parsePacket(line *bufio.Scanner) Packet {
	nextChar := line.Text()[0]

	if nextChar == '[' {
		res := Packet{Type: List, List: []*Packet{}}

		// skip initial [
		line.Scan()
		nextChar = line.Text()[0]

		for nextChar != ']' {
			parsed := parsePacket(line)
			res.List = append(res.List, &parsed)

			nextChar = line.Text()[0]
			// skip ,
			if nextChar == ',' {
				line.Scan()
				nextChar = line.Text()[0]
			}
		}
		// skip final ]
		line.Scan()
		return res
	} else if nextChar <= '9' && nextChar >= '0' {
		res := Packet{Type: Integer}

		for nextChar <= '9' && nextChar >= '0' {
			res.Integer *= 10
			res.Integer += int(nextChar - '0')

			line.Scan()
			nextChar = line.Text()[0]
		}
		return res
	} else {
		panic(fmt.Sprintf("cannot packet starting with caracter %c", nextChar))
	}
}
func ParsePacket(line string) Packet {
	scanner := bufio.NewScanner(strings.NewReader(line))
	scanner.Split(bufio.ScanRunes)
	scanner.Scan()
	return parsePacket(scanner)
}

type Signal []*Packet

func (a Signal) Len() int           { return len(a) }
func (a Signal) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Signal) Less(i, j int) bool { return a[i].Cmp(a[j]) < 0 }

func ParseInput() Signal {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	res := []*Packet{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		packet := ParsePacket(line)
		res = append(res, &packet)
	}
	return res
}

func main() {
	signal := ParseInput()

	dividerPacket1 := &Packet{Type: List, List: []*Packet{{Type: List, List: []*Packet{{Type: Integer, Integer: 2}}}}}
	dividerPacket2 := &Packet{Type: List, List: []*Packet{{Type: List, List: []*Packet{{Type: Integer, Integer: 6}}}}}
	signal = append(signal, dividerPacket1, dividerPacket2)

	sort.Sort(signal)

	res := 1
	for i, packet := range signal {
		if packet == dividerPacket1 {
			res *= i + 1
		} else if packet == dividerPacket2 {
			res *= i + 1
		}
	}
	fmt.Println(res)
}
