package day13

import (
	"bufio"
	"fmt"
	"io"
	"slices"

	"github.com/makarchuk/aoc2024/pkg/field"
)

type Arcade struct {
	OffsetA field.Point
	OffsetB field.Point
	Target  field.Point
}

func (a *Arcade) OptimalSolution() int {
	possibleSolutions := []int{}

	minIncluded := min(a.Target.X/a.OffsetA.X, a.Target.Y/a.OffsetA.Y)
	for i := 0; i <= minIncluded; i++ {
		remainder := a.Target.Sub(a.OffsetA.Mul(i))
		if remainder.X%a.OffsetB.X == 0 && remainder.Y%a.OffsetB.Y == 0 {
			if remainder.X/a.OffsetB.X == remainder.Y/a.OffsetB.Y {
				price := 3*i + 1*(remainder.X/a.OffsetB.X)
				possibleSolutions = append(possibleSolutions, price)
			}
		}
	}
	if len(possibleSolutions) == 0 {
		return 0
	}
	return slices.Min(possibleSolutions)
}

func ParseArcades(input io.Reader) ([]Arcade, error) {
	scanner := bufio.NewScanner(input)
	var arcades []Arcade

	for {
		if !scanner.Scan() {
			return nil, fmt.Errorf("unexpected EOF, expected button A")
		}

		offsetA := field.Point{}
		n, err := fmt.Sscanf(scanner.Text(), "Button A: X+%v, Y+%v", &offsetA.X, &offsetA.Y)
		if n != 2 {
			fmt.Println(scanner.Text(), offsetA, n)
			return nil, fmt.Errorf("invalid button A format: %s", scanner.Text())
		}
		if err != nil {
			return nil, fmt.Errorf("failed to parse button A: %w", err)
		}

		if !scanner.Scan() {
			return nil, fmt.Errorf("unexpected EOF, expected button B")
		}

		offsetB := field.Point{}
		n, err = fmt.Sscanf(scanner.Text(), "Button B: X+%d, Y+%d", &offsetB.X, &offsetB.Y)
		if n != 2 {
			return nil, fmt.Errorf("invalid button B format: %s", scanner.Text())
		}
		if err != nil {
			return nil, fmt.Errorf("failed to parse button B: %w", err)
		}

		if !scanner.Scan() {
			return nil, fmt.Errorf("unexpected EOF, expected prize")
		}

		target := field.Point{}
		n, err = fmt.Sscanf(scanner.Text(), "Prize: X=%d, Y=%d", &target.X, &target.Y)
		if n != 2 {
			return nil, fmt.Errorf("invalid prize format: %s", scanner.Text())
		}
		if err != nil {
			return nil, fmt.Errorf("failed to parse prize: %w", err)
		}

		arcades = append(arcades, Arcade{
			OffsetA: offsetA,
			OffsetB: offsetB,
			Target:  target,
		})

		if !scanner.Scan() {
			break
		}
	}

	return arcades, nil
}
