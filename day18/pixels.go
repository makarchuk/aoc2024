package day18

import (
	"bufio"
	"fmt"
	"io"

	"github.com/makarchuk/aoc2024/pkg/field"
	"github.com/makarchuk/aoc2024/pkg/set"
)

type Input struct {
	Screen field.Field
	Pixels []field.Point
	Size   field.Point
}

const (
	EMPTY  = '.'
	BROKEN = '#'
)

const DEBUG = false

func Part1(in io.Reader) (string, error) {
	input, err := ReadInput(in, DEBUG)
	if err != nil {
		return "", err
	}

	pixels := 1024
	if DEBUG {
		pixels = 12
	}

	for i := range pixels {
		coord := input.Pixels[i]
		input.Screen = input.Screen.Replace(coord, BROKEN)
	}

	s, err := input.FindRoute(field.Point{0, 0}, input.Size)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", s), nil
}

func Part2(in io.Reader) (string, error) {
	input, err := ReadInput(in, DEBUG)
	if err != nil {
		return "", err
	}
	for _, pixel := range input.Pixels {
		input.Screen = input.Screen.Replace(pixel, BROKEN)
		_, err := input.FindRoute(field.Point{0, 0}, input.Size)
		if err != nil {
			return fmt.Sprintf("%d,%d", pixel.X, pixel.Y), nil
		}
	}

	return "", fmt.Errorf("found a solution for every pixel")
}

func ReadInput(in io.Reader, debug bool) (Input, error) {
	input := Input{}
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		p := field.Point{}
		_, err := fmt.Sscanf(scanner.Text(), "%d,%d", &p.X, &p.Y)
		if err != nil {
			return input, err
		}
		input.Pixels = append(input.Pixels, p)
	}

	input.Size = field.Point{X: 70, Y: 70}
	if debug {
		input.Size = field.Point{X: 6, Y: 6}
	}

	input.Screen = field.Create(input.Size, EMPTY)

	return input, nil
}

var neighbours = []field.Point{
	{X: 0, Y: -1},
	{X: 0, Y: 1},
	{X: -1, Y: 0},
	{X: 1, Y: 0},
}

func (i Input) FindRoute(from, to field.Point) (int, error) {
	steps := 0

	surface := set.New[field.Point]()
	surface.Add(from)
	visited := set.New[field.Point]()
	visited.Add(from)

	for {
		nextSurface := set.New[field.Point]()
		steps++

		for _, p := range surface.List() {
			for _, neighbour := range neighbours {
				newPoint := p.Add(neighbour)

				v, ok := i.Screen.Get(newPoint)
				if !ok {
					continue
				}

				if v == BROKEN {
					continue
				}

				if visited.Contains(newPoint) {
					continue
				}

				if newPoint == to {
					return steps, nil
				}

				visited.Add(newPoint)
				nextSurface.Add(newPoint)
			}
		}

		if nextSurface.Len() == 0 {

			for y := range i.Screen.Size().Y {
				for x := range i.Screen.Size().X {
					v, _ := i.Screen.Get(field.Point{X: x, Y: y})
					if v == BROKEN {
						fmt.Printf("%c", BROKEN)
					} else if visited.Contains(field.Point{X: x, Y: y}) {
						fmt.Printf("%c", 'O')
					} else {
						fmt.Printf("%c", EMPTY)
					}
				}
				fmt.Println()
			}

			return -1, fmt.Errorf("no route found")
		}

		surface = nextSurface
	}
}
