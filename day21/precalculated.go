package day21

import (
	"fmt"

	"github.com/makarchuk/aoc2024/pkg/field"
	"github.com/makarchuk/aoc2024/pkg/set"
)

type CostProvider interface {
	PressCost(from, to byte) int
}

type ManualPresser struct{}

func (ManualPresser) PressCost(from, to byte) int {
	return 1
}

type pair = [2]byte
type costmap = map[pair]int

type PrecaclulatedCostProvider struct {
	costs costmap
}

func (p PrecaclulatedCostProvider) PressCost(from, to byte) int {
	c, ok := p.costs[pair{from, to}]
	if !ok {
		panic(fmt.Sprintf("cost from %c to %c not found", from, to))
	}
	return c
}

func EnterCombinationPrecalculated(combination string, provider CostProvider) int {
	sum := 0
	currentPosition := 'A'
	for _, c := range combination {
		sum += provider.PressCost(byte(currentPosition), byte(c))
		currentPosition = c
	}
	return sum
}

func BuildPrecalculatedCostProvider(pad keypad, previousProvider CostProvider) PrecaclulatedCostProvider {
	costs := make(costmap)
	for from, start := range pad.buttons {
		for to, finish := range pad.buttons {
			if from == to {
				//press A again if we'are there already
				costs[pair{from, to}] = previousProvider.PressCost('A', 'A')
				// fmt.Printf("%c->%c: %d\n", from, to, costs[pair{from, to}])
				continue
			}
			costs[pair{from, to}] = findCheapestPath(start, finish, pad, previousProvider)
			// fmt.Printf("%c->%c: %d\n", from, to, costs[pair{from, to}])
		}
	}
	return PrecaclulatedCostProvider{costs: costs}
}

type traverseState struct {
	position         field.Point
	providerPosition byte
	cost             int
}

func findCheapestPath(start, finish field.Point, pad keypad, provider CostProvider) int {
	intiailState := traverseState{
		position:         start,
		providerPosition: 'A',
		cost:             0,
	}
	frontier := []traverseState{intiailState}
	visited := set.New[traverseState]()
	visited.Add(intiailState)

	bestCost := -1

	for {
		newFrontier := []traverseState{}
		for _, state := range frontier {
			if bestCost != -1 && state.cost >= bestCost {
				continue
			}

			for _, button := range moveButtons {
				additionalCost := provider.PressCost(state.providerPosition, button.button)
				newPosition := state.position.Move(button.direction)

				if newPosition == finish {
					cost := state.cost + additionalCost + provider.PressCost(button.button, 'A')
					if bestCost == -1 || cost < bestCost {
						bestCost = cost
					}
					continue
				}

				if !pad.AllowedPoint(newPosition) {
					continue
				}

				newState := traverseState{
					position:         newPosition,
					cost:             state.cost + additionalCost,
					providerPosition: button.button,
				}
				if visited.Contains(newState) {
					continue
				}

				visited.Add(newState)
				newFrontier = append(newFrontier, newState)
			}
		}

		frontier = newFrontier
		if len(frontier) == 0 {
			break
		}
	}

	if bestCost == -1 {
		fmt.Sprintf("could not find path from %v to %v", start, finish)
	}

	return bestCost
}
