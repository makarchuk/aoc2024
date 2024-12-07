package day6

import (
	"bufio"
	"fmt"
	"io"
)

var ErrorOutOfField = fmt.Errorf("out of field")

type Guard struct {
	position    Point
	orientation Direction
}

func (g Guard) Position() Point {
	return g.position
}

func (g Guard) Clone() Guard {
	return Guard{
		position:    g.position,
		orientation: g.orientation,
	}
}

type Point struct {
	X, Y int
}

func (p Point) Add(p2 Point) Point {
	return Point{p.X + p2.X, p.Y + p2.Y}
}

func (p Point) Move(d Direction) Point {
	switch d {
	case UP:
		return Point{p.X, p.Y - 1}
	case RIGHT:
		return Point{p.X + 1, p.Y}
	case DOWN:
		return Point{p.X, p.Y + 1}
	case LEFT:
		return Point{p.X - 1, p.Y}
	default:
		panic(fmt.Sprintf("invalid direction: %v", d))
	}
}

type Direction byte

const (
	UP    Direction = 0
	RIGHT Direction = 1
	DOWN  Direction = 2
	LEFT  Direction = 3
)

func (d Direction) TurnRight() Direction {
	return (d + 1) % 4
}

type Field struct {
	field [][]byte

	size Point

	originalGuardPosition    Point
	originalGuardOrientation Direction
}

func (f Field) Size() Point {
	return f.size
}

func New(io io.Reader) (*Field, error) {
	field := &Field{}
	scanner := bufio.NewScanner(io)
	y := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		row := make([]byte, 0, len(line))
		for x, char := range line {
			switch char {
			case 'v':
				field.originalGuardOrientation = DOWN
				field.originalGuardPosition = Point{x, y}
				row = append(row, '.')
			case '^':
				field.originalGuardOrientation = UP
				field.originalGuardPosition = Point{x, y}
				row = append(row, '.')
			case '>':
				field.originalGuardOrientation = RIGHT
				field.originalGuardPosition = Point{x, y}
				row = append(row, '.')
			case '<':
				field.originalGuardOrientation = LEFT
				field.originalGuardPosition = Point{x, y}
				row = append(row, '.')
			case '#':
				row = append(row, '#')
			case '.':
				row = append(row, '.')
			default:
				return nil, fmt.Errorf("invalid character %c", char)
			}
		}
		field.field = append(field.field, row)
		y++
	}
	field.size = Point{len(field.field[0]) - 1, len(field.field) - 1}
	return field, nil
}

func (f *Field) Get(x, y int) byte {
	if x < 0 || x > f.size.X || y < 0 || y > f.size.Y {
		return 0
	}
	return f.field[y][x]
}

func (f *Field) GetGuard() Guard {
	return Guard{f.originalGuardPosition, f.originalGuardOrientation}
}

func (f *Field) WithExtraStone(p Point) Field {
	newField := Field{
		field:                    make([][]byte, len(f.field)),
		size:                     f.size,
		originalGuardPosition:    f.originalGuardPosition,
		originalGuardOrientation: f.originalGuardOrientation,
	}

	for y, row := range f.field {
		newRow := make([]byte, len(row))
		copy(newRow, row)
		if y == p.Y {
			newRow[p.X] = '#'
		}
		newField.field[y] = newRow
	}
	return newField
}

func (g *Guard) Move(f Field) error {
	newPos, err := g.NextPosition(f)
	if err != nil {
		return err
	}
	g.position = newPos
	return nil
}

func (g *Guard) NextPosition(f Field) (Point, error) {
	newCoordinates := g.position.Move(g.orientation)
	if newCoordinates.X < 0 || newCoordinates.X > f.size.X || newCoordinates.Y < 0 || newCoordinates.Y > f.size.Y {
		return Point{}, ErrorOutOfField
	}
	if f.Get(newCoordinates.X, newCoordinates.Y) == '#' {
		g.orientation = g.orientation.TurnRight()
		return g.position, nil
	}
	return newCoordinates, nil
}
