package field

import (
	"bufio"
	"fmt"
	"io"
)

type Point struct {
	X, Y int
}

func (p Point) Add(other Point) Point {
	return Point{p.X + other.X, p.Y + other.Y}
}

func (p Point) Sub(other Point) Point {
	return Point{p.X - other.X, p.Y - other.Y}
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
		row := make([]byte, len(line))
		copy(row, line)
		field.field = append(field.field, row)
		y++
	}
	field.size = Point{len(field.field[0]) - 1, len(field.field) - 1}
	return field, nil
}

func (f *Field) InField(p Point) bool {
	return p.X >= 0 && p.Y >= 0 && p.X <= f.size.X && p.Y <= f.size.Y
}

func (f *Field) Get(p Point) byte {
	if !f.InField(p) {
		return 0
	}
	return f.field[p.Y][p.X]
}

func (f *Field) Replace(p Point, value byte) Field {
	newField := Field{
		field: make([][]byte, len(f.field)),
		size:  f.size,
	}

	for y, row := range f.field {
		newRow := make([]byte, len(row))
		copy(newRow, row)
		if y == p.Y {
			newRow[p.X] = value
		}
		newField.field[y] = newRow
	}
	return newField
}
