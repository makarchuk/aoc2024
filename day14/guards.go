package day14

import (
	"bufio"
	"fmt"
	"io"

	"github.com/makarchuk/aoc2024/pkg/field"
	"github.com/makarchuk/aoc2024/pkg/set"
)

type Input struct {
	Guards []Guard
	Size   field.Point
}

type Guard struct {
	Position field.Point
	Velocity field.Point
}

func ParseInput(r io.Reader) (Input, error) {
	input := Input{
		Guards: []Guard{},
		Size:   field.Point{X: 100, Y: 102},
	}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var guard Guard
		_, err := fmt.Sscanf(scanner.Text(), "p=%d,%d v=%d,%d", &guard.Position.X, &guard.Position.Y, &guard.Velocity.X, &guard.Velocity.Y)
		if err != nil {
			return Input{}, err
		}
		input.Guards = append(input.Guards, guard)
	}
	if err := scanner.Err(); err != nil {
		return Input{}, err
	}
	return input, nil
}

func (g Guard) Move(size field.Point) Guard {
	return Guard{
		Position: field.Point{
			X: wrappingAdd(g.Position.X, g.Velocity.X, size.X),
			Y: wrappingAdd(g.Position.Y, g.Velocity.Y, size.Y),
		},
		Velocity: g.Velocity,
	}
}

func wrappingAdd(a, b, size int) int {
	res := (a + b) % size
	if res < 0 {
		return size + res
	}
	return res
}

func (i *Input) SafetyFactor() int {
	quardrants := [4]int{}

	for _, guard := range i.Guards {
		if guard.Position.X < i.Size.X/2 && guard.Position.Y < i.Size.Y/2 {
			quardrants[0]++
		} else if guard.Position.X > i.Size.X/2 && guard.Position.Y < i.Size.Y/2 {
			quardrants[1]++
		} else if guard.Position.X < i.Size.X/2 && guard.Position.Y > i.Size.Y/2 {
			quardrants[2]++
		} else if guard.Position.X > i.Size.X/2 && guard.Position.Y > i.Size.Y/2 {
			quardrants[3]++
		}
	}

	factor := 1
	fmt.Println(quardrants)
	for _, q := range quardrants {
		factor *= q
	}
	return factor
}

func (i *Input) Step() {
	for idx := range i.Guards {
		i.Guards[idx] = i.Guards[idx].Move(i.Size)
	}
}

func (i *Input) Print() {
	for y := range i.Size.Y {
		for x := range i.Size.X {
			found := false
			for _, guard := range i.Guards {
				if guard.Position.X == x && guard.Position.Y == y {
					found = true
					break
				}
			}
			if found {
				fmt.Printf("*")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (i *Input) ContainsVerticalLine() bool {
	coordinates := set.New[field.Point]()
	for _, guard := range i.Guards {
		coordinates.Add(guard.Position)
	}

externalLoop:
	for _, point := range coordinates.List() {
		for offset := range 10 {
			if !coordinates.Contains(field.Point{X: point.X, Y: point.Y + offset}) {
				break externalLoop
			}
		}
		return true
	}

	return false
}
