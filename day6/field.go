package day6

import (
	"fmt"
	"io"

	"github.com/makarchuk/aoc2024/pkg/field"
)

var ErrorOutOfField = fmt.Errorf("out of field")

type Input struct {
	Field         field.Field
	OriginalGuard Guard
}

func New(io io.Reader) (*Input, error) {
	f, err := field.New(io)
	if err != nil {
		return nil, err
	}
	for y := range f.Size().Y {
		for x := range f.Size().X {
			char := f.Get(field.Point{x, y})
			switch char {
			case '^':
				return &Input{
					Field: f.Replace(field.Point{x, y}, '.'),
					OriginalGuard: Guard{
						position:    field.Point{x, y},
						orientation: field.UP,
					},
				}, nil

			case 'v':
				return &Input{
					Field: f.Replace(field.Point{x, y}, '.'),
					OriginalGuard: Guard{
						position:    field.Point{x, y},
						orientation: field.DOWN,
					},
				}, nil
			case '>':
				return &Input{
					Field: f.Replace(field.Point{x, y}, '.'),
					OriginalGuard: Guard{
						position:    field.Point{x, y},
						orientation: field.RIGHT,
					},
				}, nil
			case '<':
				return &Input{
					Field: f.Replace(field.Point{x, y}, '.'),
					OriginalGuard: Guard{
						position:    field.Point{x, y},
						orientation: field.LEFT,
					},
				}, nil
			case '.', '#':
				continue
			default:
				return nil, fmt.Errorf("invalid character: %c", char)
			}
		}
	}
	return nil, fmt.Errorf("guard not found")
}

type Guard struct {
	position    field.Point
	orientation field.Direction
}

func (g Guard) Position() field.Point {
	return g.position
}

func (g Guard) Clone() Guard {
	return Guard{
		position:    g.position,
		orientation: g.orientation,
	}
}

func (g *Guard) Move(f field.Field) error {
	newPos, err := g.NextPosition(f)
	if err != nil {
		return err
	}
	g.position = newPos
	return nil
}

func (g *Guard) NextPosition(f field.Field) (field.Point, error) {
	newCoordinates := g.position.Move(g.orientation)
	if newCoordinates.X < 0 || newCoordinates.X > f.Size().X || newCoordinates.Y < 0 || newCoordinates.Y > f.Size().Y {
		return field.Point{}, ErrorOutOfField
	}
	if f.Get(newCoordinates) == '#' {
		g.orientation = g.orientation.TurnRight()
		return g.position, nil
	}
	return newCoordinates, nil
}
