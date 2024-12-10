package day10

import (
	"fmt"
	"io"

	"github.com/makarchuk/aoc2024/pkg/field"
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

func TerrainScore(t *Terrain) (int, int) {
	totalPeaks := 0
	totalTrails := 0

	for point := range t.space.Iter() {
		elevation, _ := t.space.Get(point)
		if elevation == 0 {
			peaks, trails := score(point, t.space)
			totalPeaks += peaks
			totalTrails += trails
		}

	}
	return totalPeaks, totalTrails
}

func score(p field.Point, space field.Field) (int, int) {
	var expectedElevation byte = 1
	currentSurface := map[field.Point]int{
		p: 1,
	}

	for {
		nextSurface := map[field.Point]int{}
		for point, trails := range currentSurface {
			for _, surroundingPoint := range surroundingPoints {
				newPoint := point.Add(surroundingPoint)
				elevation, ok := space.Get(newPoint)
				if !ok {
					continue
				}
				if elevation == expectedElevation {
					nextSurface[newPoint] += trails
				}
			}
		}
		if expectedElevation == 9 {
			totalScore := 0
			for _, value := range nextSurface {
				totalScore += value
			}
			return len(nextSurface), totalScore
		}
		if len(nextSurface) == 0 {
			return 0, 0
		}
		currentSurface = nextSurface
		expectedElevation++
	}
}
