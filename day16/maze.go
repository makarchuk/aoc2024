package day16

import (
	"errors"
	"fmt"
	"io"

	"github.com/makarchuk/aoc2024/pkg/field"
	"github.com/makarchuk/aoc2024/pkg/priorityqueue"
	"github.com/makarchuk/aoc2024/pkg/set"
)

const (
	WALL  = '#'
	SPACE = '.'
	START = 'S'
	END   = 'E'
)

type Input struct {
	Maze  field.Field
	Start Location
	End   field.Point
}

func ParseInput(r io.Reader) (Input, error) {
	maze, err := field.New(r)
	if err != nil {
		return Input{}, err
	}
	var extractions map[byte][]field.Point
	extractions, *maze = maze.Extract([]byte{START, END}, '.')

	if len(extractions[START]) != 1 {
		return Input{}, errors.New("expected exactly one start point")
	}

	if len(extractions[END]) != 1 {
		return Input{}, errors.New("expected exactly one end point")
	}

	return Input{
		Maze: *maze,
		Start: Location{
			Position:  extractions[START][0],
			Direction: field.RIGHT,
		},
		End: extractions[END][0],
	}, nil

}

func Part1(r io.Reader) (string, error) {
	inp, err := ParseInput(r)
	if err != nil {
		return "", err
	}

	score, err := inp.FindRoute()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", score), nil
}

func Part2(r io.Reader) (string, error) {
	inp, err := ParseInput(r)
	if err != nil {
		return "", err
	}

	minScore, err := inp.FindRoute()
	if err != nil {
		return "", err
	}

	onBestPaths, err := inp.FindBestPaths(minScore)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", onBestPaths.Len()), nil
}

type Location struct {
	Position  field.Point
	Direction field.Direction
}

type VisitedTracker struct {
	minScore int
	routes   set.Set[Location]
}

func (in Input) FindBestPaths(minScore int) (set.Set[field.Point], error) {
	result := set.New[field.Point]()
	result.Add(in.Start.Position)

	routes := set.New[Location]()
	routes.Add(in.Start)

	visited := map[Location]VisitedTracker{}
	visited[in.Start] = VisitedTracker{
		minScore: 0,
		routes:   set.New[Location](),
	}

	pq := priorityqueue.PriorityQueue[TraverserState]{}

	initialState := TraverserState{
		l:     in.Start,
		score: 0,
	}

	pq.Push(-initialState.score, initialState)

	counter := 0

	for {
		state, ok := pq.Pop()
		if !ok {
			seen := routes.Clone()
			surface := routes
			for {
				newSurface := set.New[Location]()
				for _, location := range surface.List() {
					result.Add(location.Position)
					for _, p := range visited[location].routes.List() {
						if !seen.Contains(p) {
							newSurface.Add(p)
						}
						seen.Add(p)
						result.Add(p.Position)
					}
				}
				surface = newSurface
				if surface.Len() == 0 {
					return result, nil
				}
			}
		}

		if state.score > minScore {
			// fmt.Printf("score=%d, dropping. left in queue: %v\n", state.score, pq.Len())
			continue
		}

		for _, option := range state.Moves() {
			if option.l.Position == in.End {
				if option.score == minScore {
					counter++
					// fmt.Printf("Found best path: %v, len: %v, score: %v\n", counter, option.route.Len(), option.score)
					// fmt.Printf("left in queue: %v\n", pq.Len())
					for _, locationOnRoute := range option.route.List() {
						routes.Add(locationOnRoute)
					}
					continue
				}

				if option.score >= minScore {
					continue
				}

			}

			val, ok := in.Maze.Get(option.l.Position)
			if !ok {
				panic("out of bounds")
			}

			if val == SPACE {
				previousVisits, ok := visited[option.l]

				if !ok || option.score < previousVisits.minScore {
					pq.Push(-option.score, option)
					visited[option.l] = VisitedTracker{
						minScore: option.score,
						routes:   option.route.Clone(),
					}
				} else if previousVisits.minScore == option.score {
					for _, point := range option.route.List() {
						previousVisits.routes.Add(point)
					}
					visited[option.l] = previousVisits
				}
			}
		}
	}
}

func (in Input) FindRoute() (int, error) {
	visited := set.New[Location]()
	visited.Add(in.Start)

	pq := priorityqueue.PriorityQueue[TraverserState]{}

	initialState := TraverserState{
		l:     in.Start,
		route: set.New[Location](),
		score: 0,
	}

	pq.Push(-initialState.score, initialState)

	for {
		state, ok := pq.Pop()
		if !ok {
			return 0, errors.New("no path found")
		}

		for _, option := range state.Moves() {
			if option.l.Position == in.End {
				return option.score, nil
			}

			val, ok := in.Maze.Get(option.l.Position)
			if !ok {
				panic("out of bounds")
			}
			if val == SPACE {
				if !visited.Contains(option.l) {
					visited.Add(option.l)
					pq.Push(-option.score, option)
				}
			}
		}
	}
}

type TraverserState struct {
	l     Location
	route set.Set[Location]
	score int
}

func (t TraverserState) Moves() [3]TraverserState {
	return [3]TraverserState{
		t.Move(),
		t.TurnLeft(),
		t.TurnRight(),
	}
}

func (s TraverserState) Move() TraverserState {
	newRoute := s.route.Clone()
	newPosition := s.l.Position.Move(s.l.Direction)
	newLocation := Location{
		Position:  newPosition,
		Direction: s.l.Direction,
	}

	newRoute.Add(newLocation)

	return TraverserState{
		l:     newLocation,
		route: newRoute,
		score: s.score + 1,
	}
}

func (s TraverserState) TurnLeft() TraverserState {
	newRoute := s.route.Clone()
	newLocation := Location{
		Position:  s.l.Position,
		Direction: s.l.Direction.TurnLeft(),
	}
	newRoute.Add(newLocation)

	return TraverserState{
		l:     newLocation,
		route: newRoute,
		score: s.score + 1000,
	}
}

func (s TraverserState) TurnRight() TraverserState {
	newRoute := s.route.Clone()
	newLocation := Location{
		Position:  s.l.Position,
		Direction: s.l.Direction.TurnRight(),
	}
	newRoute.Add(newLocation)
	return TraverserState{
		l:     newLocation,
		route: newRoute,
		score: s.score + 1000,
	}
}
