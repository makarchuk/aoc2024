package day8

import (
	"io"

	"github.com/makarchuk/aoc2024/pkg/field"
	"github.com/makarchuk/aoc2024/pkg/set"
)

type Frequency byte

type Input struct {
	Space field.Field

	Antennas map[Frequency][]field.Point
}

func ReadInput(in io.Reader) (*Input, error) {
	space, err := field.New(in)
	if err != nil {
		return nil, err
	}

	i := &Input{
		Space:    *space,
		Antennas: make(map[Frequency][]field.Point),
	}

	for x := range space.Size().X + 1 {
		for y := range space.Size().Y + 1 {
			char := space.Get(field.Point{X: x, Y: y})
			if char == '.' {
				continue
			}
			if char == 0 {
				panic("out of range")
			}
			i.Antennas[Frequency(char)] = append(i.Antennas[Frequency(char)], field.Point{X: x, Y: y})
		}
	}

	return i, nil
}

func FindHarmonicAntinodes(space field.Field, antenas map[Frequency][]field.Point) set.Set[field.Point] {
	antinodes := set.New[field.Point]()
	antinodeArr := []field.Point{}

	for _, antenas := range antenas {
		for i, p1 := range antenas {
			for _, p2 := range antenas[i+1:] {
				localAntinodes := HarmonicAntiNodes(p1, p2, space)
				for _, antinode := range localAntinodes {
					antinodes.Add(antinode)
					antinodeArr = append(antinodeArr, antinode)
				}
			}
		}
	}
	return antinodes
}

func FindAntinodes(space field.Field, antenas map[Frequency][]field.Point) set.Set[field.Point] {
	// fmt.Printf("Size: %v\n", space.Size())
	antinodes := set.New[field.Point]()
	antinodeArr := []field.Point{}

	for _, antenas := range antenas {
		for i, p1 := range antenas {
			for j, p2 := range antenas[i+1:] {
				_ = j
				localAntinodes := AntiNodes(p1, p2)
				for _, antinode := range localAntinodes {
					if !space.InField(antinode) {
						continue
					}
					antinodes.Add(antinode)
					antinodeArr = append(antinodeArr, antinode)
				}
			}
		}
	}
	return antinodes
}

func AntiNodes(p1, p2 field.Point) []field.Point {
	a1 := p1.Add(p1).Sub(p2)
	a2 := p2.Add(p2).Sub(p1)

	d := p2.Sub(p1)

	res := []field.Point{a1, a2}

	if d.X%3 == 0 && d.Y%3 == 0 {
		deltaVector := field.Point{X: d.X / 3, Y: d.Y / 3}
		res = append(res, p1.Add(deltaVector), p2.Sub(deltaVector))
	}
	return res
}

func HarmonicAntiNodes(p1, p2 field.Point, space field.Field) []field.Point {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y

	//simplify dy/dx to lowest terms
	for i := 2; i <= max(dx, dy); i++ {
		for dx%i == 0 && dy%i == 0 {
			dx /= i
			dy /= i
		}
	}

	stepVector := field.Point{X: dx, Y: dy}

	antinodes := []field.Point{p1}
	point := p1
	for {
		point = point.Add(stepVector)
		// fmt.Printf("Up: %v\n", point)
		if !space.InField(point) {
			break
		}
		antinodes = append(antinodes, point)
	}

	point = p1
	for {
		point = point.Sub(stepVector)
		// fmt.Printf("Down: %v\n", point)
		if !space.InField(point) {
			break
		}
		antinodes = append(antinodes, point)
	}

	return antinodes
}
