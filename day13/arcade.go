package day13

import (
	"bufio"
	"fmt"
	"io"

	"github.com/makarchuk/aoc2024/pkg/field"
)

type Arcade struct {
	OffsetA field.Point
	OffsetB field.Point
	Target  field.Point
}

func (m *Arcade) OptimalSolution() int {
	x := m.Target.X
	x1 := m.OffsetA.X
	x2 := m.OffsetB.X

	y := m.Target.Y
	y1 := m.OffsetA.Y
	y2 := m.OffsetB.Y

	aNumerator := (x*y2 - x2*y)
	aDenominator := (x1*y2 - x2*y1)

	if aNumerator%aDenominator != 0 {
		return 0
	}

	a := aNumerator / aDenominator

	bNumerator := x - a*x1
	bDenominator := x2

	if bNumerator%bDenominator != 0 {
		return 0
	}

	b := bNumerator / bDenominator

	if a < 0 || b < 0 {
		return 0
	}

	return 3*a + b
}

func (m *Arcade) calculateStep() (int, int, bool) {
	remTarget := m.Target.X % m.OffsetB.X
	remA := m.OffsetA.X % m.OffsetB.X
	fmt.Printf("remTarget: %v, remA: %v. Offset: %v\n", remTarget, remA, m.OffsetB.X)
	if remA == 0 && remTarget != 0 {
		return 0, 0, false
	}
	i := 0
	firstMatch := -1
	for {
		if (i*remA)%m.OffsetB.X == remTarget {
			if firstMatch == -1 {
				firstMatch = i
			} else {
				return firstMatch, i - firstMatch, true
			}
		}
		i++
	}
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
