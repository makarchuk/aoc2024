package day24

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strings"
)

type Input struct {
	Values map[string]Value
}

type Value struct {
	Plain      *bool
	Expression Expression
}

type Expression struct {
	Left  string
	Op    Opertor
	Right string
}

type Opertor string

const (
	OpAnd Opertor = "AND"
	OpOr  Opertor = "OR"
	OpXor Opertor = "XOR"
)

func (in *Input) GetValue(name string) (bool, error) {
	val, ok := in.Values[name]
	if !ok {
		return false, fmt.Errorf("Value %s not found", name)
	}

	if val.Plain != nil {
		return *val.Plain, nil
	}

	left, err := in.GetValue(val.Expression.Left)
	if err != nil {
		return false, err
	}

	right, err := in.GetValue(val.Expression.Right)
	if err != nil {
		return false, err
	}

	var res bool
	switch val.Expression.Op {
	case OpAnd:
		res = left && right
	case OpOr:
		res = left || right
	case OpXor:
		res = left != right
	default:
		return false, fmt.Errorf("Invalid operator %s", val.Expression.Op)
	}

	val.Plain = &res
	return res, nil
}

func ParseInput(in io.Reader) (*Input, error) {
	scanner := bufio.NewScanner(in)
	i := &Input{
		Values: make(map[string]Value),
	}

	//load values
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		line := scanner.Text()
		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			return nil, fmt.Errorf("Invalid input: %s", line)
		}

		name := parts[0]
		switch parts[1] {
		case "0":
			i.Values[name] = Value{Plain: ptr(false)}
		case "1":
			i.Values[name] = Value{Plain: ptr(true)}
		default:
			return nil, fmt.Errorf("Invalid input: %s", line)
		}
	}

	for scanner.Scan() {
		// gpv OR jrk -> whv
		line := scanner.Text()
		parts := strings.Split(line, " -> ")
		if len(parts) != 2 {
			return nil, fmt.Errorf("Invalid input: %s", line)
		}
		exprParts := strings.Split(parts[0], " ")
		if len(exprParts) != 3 {
			return nil, fmt.Errorf("Invalid input: %s", line)
		}

		var op Opertor
		switch exprParts[1] {
		case "AND":
			op = OpAnd
		case "OR":
			op = OpOr
		case "XOR":
			op = OpXor
		default:
			return nil, fmt.Errorf("Invalid input: %s", line)
		}

		i.Values[parts[1]] = Value{
			Expression: Expression{
				Left:  exprParts[0],
				Op:    op,
				Right: exprParts[2],
			},
		}
	}

	return i, nil
}

func Part1(in io.Reader) (string, error) {
	input, err := ParseInput(in)
	if err != nil {
		return "", err
	}

	values := []string{}

	for name, _ := range input.Values {
		if strings.HasPrefix(name, "z") {
			values = append(values, name)
		}
	}

	slices.Sort(values)
	slices.Reverse(values)

	num := make([]byte, len(values))
	for i, name := range values {
		val, err := input.GetValue(name)
		if err != nil {
			return "", err
		}
		var b byte
		if val {
			b = 1
		}
		num[i] = b
	}

	decodedNum := toDecimal(num)

	return fmt.Sprintf("%d", decodedNum), nil
}

func ptr(b bool) *bool {
	return &b
}

func toDecimal(in []byte) int {
	res := 0

	for _, val := range in {
		res = res*2 + int(val)
	}
	return res
}
