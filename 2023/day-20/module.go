package main

import "fmt"

type Height int

const (
	Low  = Height(0)
	High = Height(1)
)

func (h Height) String() string {
	switch h {
	case Low:
		return "low"
	case High:
		return "high"
	}
	panic(fmt.Sprintf("unknow hight %d", h))
}

type Pulse struct {
	Height      Height
	Source      string
	Destination string
}

func (p Pulse) String() string {
	return fmt.Sprintf("%s -%s -> %s", p.Source, p.Height, p.Destination)
}

type Module interface {
	Receive(pulse Pulse) []Pulse
	Destinations() []string
}
type BaseModule struct {
	Name         string
	destinations []string
}

func (b *BaseModule) Destinations() []string {
	return b.destinations
}

func (b BaseModule) Broadcast(height Height) []Pulse {
	var res []Pulse
	for _, dest := range b.destinations {
		res = append(res, Pulse{Height: height, Source: b.Name, Destination: dest})
	}
	return res
}

type FlipFlop struct {
	BaseModule
	On bool
}

func (m *FlipFlop) Receive(pulse Pulse) []Pulse {
	if pulse.Height == High {
		return []Pulse{}
	} else {
		m.On = !m.On
		if m.On {
			return m.Broadcast(High)
		} else {
			return m.Broadcast(Low)
		}
	}
}

type Conjunction struct {
	BaseModule
	MostRecentPulse map[string]Height
}

func (c *Conjunction) AllHight() bool {
	for _, v := range c.MostRecentPulse {
		if v != High {
			return false
		}
	}
	return true
}
func (c *Conjunction) Receive(pulse Pulse) []Pulse {
	c.MostRecentPulse[pulse.Source] = pulse.Height
	if c.AllHight() {
		return c.Broadcast(Low)
	} else {
		return c.Broadcast(High)
	}
}

type Broadcast struct {
	BaseModule
}

func (b *Broadcast) Receive(pulse Pulse) []Pulse {
	return b.Broadcast(pulse.Height)
}

type Button struct {
	BaseModule
}

func (b *Button) Receive(pulse Pulse) []Pulse {
	return []Pulse{{Source: "button", Destination: "broadcaster", Height: Low}}
}
