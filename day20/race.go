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

	reachableMap := i.FindAllReachableFrom(i.Finish)
	fairScore := reachableMap[i.Start]

	cheatingRoutes := i.FindRoutesWithCheat(fairScore-100, 2, reachableMap)

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

	reachableMap := i.FindAllReachableFrom(i.Finish)
	fairScore := reachableMap[i.Start]

	cheatingRoutes := i.FindRoutesWithCheat(fairScore-50, 20, reachableMap)

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

type TraverserState struct {
	Position               field.Point
	CheatStart             field.Point
	CheatEnd               field.Point
	RemainingCheatDuration int
}

func (s *TraverserState) Clone() *TraverserState {
	return &TraverserState{
		Position:   s.Position,
		CheatStart: s.CheatStart,
	}
}

func (s *TraverserState) AvailableLocations(track field.Field, cheatDuration int) []TraverserState {
	locations := []TraverserState{}

	for _, neighbour := range neightbours {
		next := s.Position.Add(neighbour)
		v, ok := track.Get(next)
		if !ok {
			continue
		}

		if s.RemainingCheatDuration >= 1 {
			if !s.IsCheating() {
				fmt.Printf("%+v\n", s)
				panic("asaaa")
			}
			//stop cheating. can't do it if we haven't went through any walls yet, to avoid loops
			if v == EMPTY {
				locations = append(locations, TraverserState{
					Position:               next,
					CheatStart:             s.CheatStart,
					CheatEnd:               next,
					RemainingCheatDuration: 0,
				})
			}
			if s.RemainingCheatDuration > 1 {
				//keep on cheating
				locations = append(locations, TraverserState{
					Position:               next,
					CheatStart:             s.CheatStart,
					CheatEnd:               cheatNotUsed,
					RemainingCheatDuration: s.RemainingCheatDuration - 1,
				})
			}
			continue
		}

		//just move
		if v == EMPTY {
			locations = append(locations, TraverserState{
				Position:   next,
				CheatStart: s.CheatStart,
				CheatEnd:   s.CheatEnd,
			})
		}

		if !s.UsedCheat() {
			//start cheating
			locations = append(locations, TraverserState{
				Position:               next,
				CheatStart:             s.Position,
				CheatEnd:               cheatNotUsed,
				RemainingCheatDuration: cheatDuration - 1,
			})
		}

	}

	return locations
}

func (s *TraverserState) cacheKey() cacheKey {
	return cacheKey{
		Position:   s.Position,
		CheatStart: s.CheatStart,
		CheatEnd:   s.CheatEnd,
	}
}

func (s *TraverserState) IsCheating() bool {
	return s.CheatStart != cheatNotUsed && s.CheatEnd == cheatNotUsed
}

func (s *TraverserState) UsedCheat() bool {
	return s.CheatStart != cheatNotUsed && s.CheatEnd != cheatNotUsed
}

type cacheKey struct {
	Position   field.Point
	CheatStart field.Point
	CheatEnd   field.Point
}

func (in *Input) FindRoutesWithCheat(maxSteps int, cheatDuration int, fairlyReachable map[field.Point]int) []Route {
	startingState := TraverserState{
		Position:   in.Start,
		CheatStart: cheatNotUsed,
		CheatEnd:   cheatNotUsed,
	}

	frontier := set.New[TraverserState]()
	frontier.Add(startingState)
	visited := set.New[cacheKey]()
	visited.Add(startingState.cacheKey())

	routes := []Route{}

	steps := 0
	for {
		steps++
		if steps > maxSteps {
			return routes
		}

		newFrontier := set.New[TraverserState]()
		for s := range frontier.Iter() {
			for _, next := range s.AvailableLocations(in.Track, cheatDuration) {
				if next.UsedCheat() {
					fairScoreFromPoint, ok := fairlyReachable[next.Position]
					if !ok {
						continue
					}

					totalScore := fairScoreFromPoint + steps
					if totalScore > maxSteps {
						continue
					}
					// fmt.Printf("found solution for %v %v steps\n", [2]field.Point{next.CheatStart, next.CheatEnd}, totalScore)
					routes = append(routes, Route{
						Cheat: [2]field.Point{next.CheatStart, next.CheatEnd},
						Steps: totalScore,
					})
					continue
				}

				if next.Position == in.Finish {
					if next.CheatEnd == cheatNotUsed {
						next.CheatEnd = in.Finish
						next.RemainingCheatDuration = 0
					} else {
						fmt.Printf("%+v\n", next)
						panic("should not be here, if we've finished the cheat already, we should've looked the map up")
					}

					// fmt.Printf("found solution for %v %v steps\n", [2]field.Point{next.CheatStart, next.CheatEnd}, steps)

					routes = append(routes, Route{
						Cheat: [2]field.Point{next.CheatStart, next.CheatEnd},
						Steps: steps,
					})
					// fmt.Printf(
					// 	"found route in %v steps. Cheat: %+v frontier: %v\n",
					// 	steps,
					// 	[2]field.Point{next.CheatStart, next.CheatEnd},
					// 	frontier.Len(),
					// )
					// continue
				}

				cacheKey := next.cacheKey()
				if visited.Contains(cacheKey) {
					continue
				}

				visited.Add(cacheKey)
				newFrontier.Add(next)
			}
		}

		if newFrontier.Len() == 0 {
			return routes
		}
		frontier = newFrontier
	}

}
