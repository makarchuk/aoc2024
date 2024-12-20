package day20

import (
	"fmt"
	"io"

	"github.com/makarchuk/aoc2024/pkg/field"
	"github.com/makarchuk/aoc2024/pkg/set"
)

const (
	EMPTY  = '.'
	WALL   = '#'
	START  = 'S'
	FINISH = 'E'
)

type Input struct {
	Track  field.Field
	Start  field.Point
	Finish field.Point
}

// fake coordinates to indicate that cheat is not used
var cheatNotUsed = field.Point{-1, -1}

var neightbours = []field.Point{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func Part1(in io.Reader) (string, error) {
	i, err := ParseInput(in)
	if err != nil {
		return "", err
	}

	cheatingRoutes := i.FindAllPossibleCheats(2, 100)

	differentCheatPoints := set.New[[2]field.Point]()
	for _, route := range cheatingRoutes {
		differentCheatPoints.Add(route.Cheat)
	}

	return fmt.Sprintf("%d", differentCheatPoints.Len()), nil
}

func Part2(in io.Reader) (string, error) {
	i, err := ParseInput(in)
	if err != nil {
		return "", err
	}

	cheatingRoutes := i.FindAllPossibleCheats(20, 100)

	differentCheatPoints := set.New[[2]field.Point]()
	for _, route := range cheatingRoutes {
		differentCheatPoints.Add(route.Cheat)
	}

	return fmt.Sprintf("%d", differentCheatPoints.Len()), nil
}

func ParseInput(in io.Reader) (Input, error) {
	f, err := field.New(in)
	if err != nil {
		return Input{}, err
	}

	extracted, track := f.Extract([]byte{START, FINISH}, '.')
	if len(extracted[START]) != 1 {
		return Input{}, fmt.Errorf("expected 1 start point, got %d", len(extracted[START]))
	}

	if len(extracted[FINISH]) != 1 {
		return Input{}, fmt.Errorf("expected 1 finish point, got %d", len(extracted[FINISH]))
	}

	return Input{
		Track:  track,
		Start:  extracted[START][0],
		Finish: extracted[FINISH][0],
	}, nil
}

func (i Input) FindAllPossibleCheats(maxCheat int, saveSteps int) []Route {
	result := []Route{}

	distanceToFinish := i.FindAllReachableFrom(i.Finish)
	maxSteps := distanceToFinish[i.Start] - saveSteps

	distanceToStart := i.FindAllReachableFrom(i.Start)

	for cheatStart, preCheatScore := range distanceToStart {
		for cheatEnd, postCheatScore := range distanceToFinish {
			cheatLength := cheatStart.ManhattanDistance(cheatEnd)
			if cheatLength > maxCheat {
				continue
			}

			totalScore := preCheatScore + postCheatScore + cheatLength
			if totalScore > maxSteps {
				continue
			}

			result = append(result, Route{
				Cheat: [2]field.Point{cheatStart, cheatEnd},
				Steps: totalScore,
			})
		}
	}
	return result
}

func (i Input) FindAllReachableFrom(p field.Point) map[field.Point]int {
	frontier := set.From([]field.Point{p})
	visited := map[field.Point]int{p: 0}

	steps := 0
	for {
		steps++
		newFrontier := set.New[field.Point]()

		for p := range frontier.Iter() {
			for _, neighbour := range neightbours {
				next := p.Add(neighbour)

				if v, ok := i.Track.Get(next); !ok || v == WALL {
					continue
				}

				if _, ok := visited[next]; ok {
					continue
				}

				visited[next] = steps
				newFrontier.Add(next)
			}
		}
		if newFrontier.Len() == 0 {
			return visited
		}

		frontier = newFrontier
	}
}

type Route struct {
	Cheat [2]field.Point
	Steps int
}
