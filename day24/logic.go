package day24

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strings"

	"github.com/makarchuk/aoc2024/pkg/set"
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

	num, err := input.readNumber("z")
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", num), nil
}

func Part2(in io.Reader) (string, error) {
	input, err := ParseInput(in)
	if err != nil {
		return "", err
	}

	x, err := input.readNumber("x")
	if err != nil {
		return "", err
	}

	y, err := input.readNumber("y")
	if err != nil {
		return "", err
	}

	expectedResult := x + y

	if err != nil {
		return "", err
	}

	fmt.Printf("Original Diff: %d\n", input.DiffInBits(expectedResult))
	originalRealResult, err := input.readNumber("z")
	if err != nil {
		return "", err
	}
	originalSuspicious := input.findSuspiciousGates(expectedResult, originalRealResult).List()
	slices.Sort(originalSuspicious)
	fmt.Printf("Original Suspicious gates: %v\n", originalSuspicious)

	//swaps are filled in manually by exploring logic provided by input
	newInput := input.Clone()
	newInput.Swap("bjm", "z07")
	newInput.Swap("hsw", "z13")
	newInput.Swap("skf", "z18")
	newInput.Swap("nvr", "wkr")
	swaps := []string{"bjm", "z07", "hsw", "z13", "skf", "z18", "nvr", "wkr"}
	slices.Sort(swaps)

	fmt.Println("Diff: ", newInput.DiffInBits(expectedResult))

	realResult, err := newInput.readNumber("z")
	if err != nil {
		return "", err
	}

	l := input.findSuspiciousGates(expectedResult, realResult).List()
	slices.Sort(l)
	fmt.Printf("Suspicious gates: %v\n", l)

	type testCase struct {
		x    int
		y    int
		name string
	}

	testCases := []testCase{
		{0, 1, ""},
		{0, 2, ""},
		{1 << 7, 1 << 7, ""},
		{1 << 7, 1 << 8, ""},
		{1 << 8, 1 << 8, ""},
		{1 << 6, 1 << 6, ""},
		{1190019232 << 7, 1283718273889, ""},
	}

	for i := range 45 {
		testCases = append(testCases, testCase{
			x:    1 << i,
			y:    0,
			name: fmt.Sprintf("1<<%d, 0", i),
		})
		testCases = append(testCases, testCase{
			x:    0,
			y:    1 << i,
			name: fmt.Sprintf("0, 1<<%d", i),
		})
	}

	for _, tc := range testCases {
		fmt.Printf("testCase: %v\n", tc.name)
		newInput := newInput.SetInputs(tc.x, tc.y)
		expectedResult := tc.x + tc.y
		diff := newInput.DiffInBits(expectedResult)
		fmt.Printf("Diff on %v+%v: %d\n", tc.x, tc.y, diff)
		realResult, err := newInput.readNumber("z")
		if err != nil {
			return "", err
		}

		suspiciousGates := newInput.findSuspiciousGates(expectedResult, realResult).List()
		slices.Sort(suspiciousGates)
		fmt.Printf("Suspicious gates: %v\n", suspiciousGates)
		if diff > 0 {
			return "", fmt.Errorf("Found error")
		}
	}

	return strings.Join(swaps, ","), nil
}

func (input *Input) DiffInBits(expected int) int {
	real, err := input.readNumber("z")
	if err != nil {
		return -1
	}
	realBin := toBinary(real)
	expectedBin := toBinary(expected)
	for len(realBin) < 46 {
		realBin = append(realBin, 0)
	}

	for len(expectedBin) < 46 {
		expectedBin = append(expectedBin, 0)
	}

	fmt.Printf("Real: %v\n", realBin)
	fmt.Printf("Exp.: %v\n", expectedBin)

	diff := 0
	for i := 0; i < len(realBin); i++ {
		if realBin[i] != expectedBin[i] {
			diff++
		}
	}

	return diff
}

func swapId(trueSwaps, falseSwaps [4]string) string {
	pairs := [4]string{}
	for i := 0; i < 4; i++ {
		pairs[i] = fmt.Sprintf("%s:%s", falseSwaps[i], trueSwaps[i])
	}

	slices.Sort(pairs[:])
	return strings.Join(pairs[:], ",")
}

func (input Input) Clone() Input {
	newInput := Input{
		Values: make(map[string]Value),
	}
	newValues := make(map[string]Value)
	for k, v := range input.Values {
		//in this case plain is essentially a cache. Dropping it.
		if v.Expression.Op != "" {
			v.Plain = nil
		}
		newValues[k] = v
	}

	newInput.Values = newValues
	return newInput
}

func (in *Input) SetInputs(x, y int) Input {
	newInput := in.Clone()
	bits := toBinary(x)
	for len(bits) < 46 {
		bits = append(bits, 0)
	}

	for i, bit := range bits {
		newInput.Values[fmt.Sprintf("x%02d", i)] = Value{Plain: ptr(bit == 1)}
	}

	bits = toBinary(y)
	for len(bits) < 46 {
		bits = append(bits, 0)
	}

	for i, bit := range bits {
		newInput.Values[fmt.Sprintf("y%02d", i)] = Value{Plain: ptr(bit == 1)}
	}

	return newInput
}

func (in *Input) Swap(one, two string) {
	oneVal := in.Values[one]
	twoVal := in.Values[two]

	in.Values[one] = twoVal
	in.Values[two] = oneVal
}

func (input *Input) findSuspiciousGates(expectedResult, realResult int) set.Set[string] {
	expected := toBinary(expectedResult)
	real := toBinary(realResult)

	matchingIndices := make([]string, 0, len(real))
	unmatchingIndices := make([]string, 0, len(real))

	for i := len(expected) - 1; i >= 0; i-- {
		key := fmt.Sprintf("z%02d", i)
		if expected[i] != real[i] {
			unmatchingIndices = append(unmatchingIndices, key)
		} else {
			matchingIndices = append(matchingIndices, key)
		}
	}

	return set.From(unmatchingIndices)
}

func ptr(b bool) *bool {
	return &b
}

func toDecimal(in []byte) int {
	res := 0

	for i := len(in) - 1; i >= 0; i-- {
		res = res*2 + int(in[i])
	}
	return res
}

func (input Input) readNumber(prefix string) (int, error) {
	keys := []string{}

	for name, _ := range input.Values {
		if strings.HasPrefix(name, prefix) {
			keys = append(keys, name)
		}
	}

	slices.Sort(keys)

	num := make([]byte, len(keys))
	for i, name := range keys {
		val, err := input.GetValue(name, 100)
		if err != nil {
			return 0, err
		}
		var b byte
		if val {
			b = 1
		}
		num[i] = b
	}

	return toDecimal(num), nil
}

func toBinary(n int) []byte {
	res := []byte{}
	for n > 0 {
		res = append(res, byte(n%2))
		n /= 2
	}
	return res
}

func (in *Input) GetValue(name string, maxDepth int) (bool, error) {
	if maxDepth <= 0 {
		return false, fmt.Errorf("Max depth reached")
	}
	val, ok := in.Values[name]
	if !ok {
		return false, fmt.Errorf("Value %s not found", name)
	}

	if val.Plain != nil {
		return *val.Plain, nil
	}

	left, err := in.GetValue(val.Expression.Left, maxDepth-1)
	if err != nil {
		return false, err
	}

	right, err := in.GetValue(val.Expression.Right, maxDepth-1)
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
