package day10

import (
	"fmt"
	"io"

	"github.com/makarchuk/aoc2024/pkg/field"
	"github.com/makarchuk/aoc2024/pkg/set"
)

type Terrain struct {
	space field.Field
}

var surroundingPoints = [4]field.Point{
	{X: 0, Y: -1},
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: -1, Y: 0},
}

func New(in io.Reader) (*Terrain, error) {
	space, err := field.New(in)
	if err != nil {
		return nil, err
	}
	for point := range space.Iter() {
		char, ok := space.Get(point)
		if !ok {
			panic("out of range")
		}
		if char > '9' || char < '0' {
			panic(fmt.Sprintf("unexpected character: %c", char))
		}
		space.Set(point, char-'0')
	}

	return &Terrain{
		space: *space,
	}, nil
}

func TerrainScore(t *Terrain) int {
	totalScore := 0
	for point := range t.space.Iter() {
		elevation, _ := t.space.Get(point)
		if elevation == 0 {
			score := score(point, t.space)

			totalScore += score
		}

	}
	return totalScore
}

func score(p field.Point, space field.Field) int {
	var expectedElevation byte = 1
	currentSurface := set.New[field.Point]()
	currentSurface.Add(p)
	visited := set.New[field.Point]()

	for {
		nextSurface := set.New[field.Point]()
		for _, point := range currentSurface.List() {
			for _, surroundingPoint := range surroundingPoints {
				newPoint := point.Add(surroundingPoint)
				if visited.Contains(newPoint) {
					continue
				}
				elevation, ok := space.Get(newPoint)
				if !ok {
					continue
				}
				if elevation == expectedElevation {
					visited.Add(newPoint)
					nextSurface.Add(newPoint)
				}
			}
		}
		if expectedElevation == 9 {
			return nextSurface.Len()
		}
		if len(nextSurface.List()) == 0 {
			return 0
		}
		currentSurface = nextSurface
		expectedElevation++
	}
}
