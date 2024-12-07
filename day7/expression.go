package day7

import (
	"errors"
	"strconv"
	"strings"

	"github.com/makarchuk/aoc2024/pkg/helpers"
)

type Operator string

const (
	OperatorAdd    Operator = "+"
	OperatorMul    Operator = "*"
	OperatorConcat Operator = "||"
)

func (o Operator) Apply(a, b int) int {
	switch o {
	case OperatorAdd:
		return a + b
	case OperatorMul:
		return a * b
	case OperatorConcat:
		return Concat(a, b)
	}
	panic("unknown operator")
}

func Concat(a, b int) int {
	order := 10

	for order <= b {
		order *= 10
	}

	return a*order + b
}

type Expression struct {
	Result int

	Operands []int
}

func (e Expression) BruteforceOperators(operators []Operator) bool {
	return e.bruteforcePartial(e.Operands[0], e.Operands[1:], operators)
}

func (e Expression) bruteforcePartial(subres int, tail []int, operators []Operator) bool {
	head := tail[0]
	for _, operator := range operators {
		subres := operator.Apply(subres, head)
		if len(tail) == 1 {
			if subres == e.Result {
				return true
			}
		} else {
			if e.bruteforcePartial(subres, tail[1:], operators) {
				return true
			}
		}
	}
	return false
}

func (e *Expression) Parse(s string) error {
	split := strings.SplitN(s, ": ", 2)
	if len(split) < 2 {
		return errors.New("invalid expression")
	}

	result, err := strconv.Atoi(split[0])
	if err != nil {
		return err
	}
	e.Result = result
	e.Operands, err = helpers.ParseIntsArray(split[1], " ")
	if err != nil {
		return err
	}

	return nil
}
