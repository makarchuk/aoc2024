package day17

import (
	"fmt"
	"io"
	"math"
	"strings"

	"github.com/makarchuk/aoc2024/pkg/helpers"
	"github.com/makarchuk/aoc2024/pkg/set"
)

var ErrEndOfInput = fmt.Errorf("end of input")

type Computer struct {
	RegisterA int
	RegisterB int
	RegisterC int

	Program Program

	Output []byte
}

func ParseInput(in io.Reader) (Computer, error) {
	var c Computer

	var program string

	_, err := fmt.Fscanf(in, `Register A: %d
Register B: %d
Register C: %d

Program: %s`, &c.RegisterA, &c.RegisterB, &c.RegisterC, &program)

	if err != nil {
		return c, err
	}

	program = strings.Trim(program, "\n")
	nums, err := helpers.ParseIntsArray(program, ",")
	if err != nil {
		return c, err
	}
	for _, n := range nums {
		c.Program.Memory = append(c.Program.Memory, byte(n))
	}
	return c, nil
}

func Part1(in io.Reader) (string, error) {
	comp, err := ParseInput(in)
	if err != nil {
		return "", err
	}
	output := CompiledProgram(comp.RegisterA)

	out := []string{}
	for _, b := range output {
		out = append(out, fmt.Sprintf("%d", b))
	}

	return strings.Join(out, ","), nil
}

func Part2(in io.Reader) (string, error) {
	comp, err := ParseInput(in)
	if err != nil {
		return "", err
	}

	candidates := set.New[int]()
	candidates.Add(0)
	for i := 0; i < len(comp.Program.Memory); i++ {
		newCandidates := set.New[int]()
		memSlice := comp.Program.Memory[len(comp.Program.Memory)-i-1:]
		for _, knownNum := range candidates.List() {
			for d := 0; d < 8; d++ {
				input := knownNum*8 + d
				out := CompiledProgram(input)
				// fmt.Println("Input:", input, "Output:", out)
				match := true
				if len(out) != i+1 {
					// fmt.Printf("skipping. Unexpected output len: %v. Expected: %v\n", len(out), i+1)
					continue
				}

				for j := 0; j <= i; j++ {
					if out[j] != int(memSlice[j]) {
						match = false
						break
					}
				}
				if match {
					newCandidates.Add(input)
					// fmt.Printf("stored input: %d\n", input)
				}
			}
		}
		if newCandidates.Len() == 0 {
			return "", fmt.Errorf("no candidates found")
		}
		candidates = newCandidates
		// fmt.Println("candidates:", candidates.List())
	}
	if candidates.Len() != 1 {
		return "", fmt.Errorf("expected a single solution, got %d", candidates.Len())
	}
	return fmt.Sprintf("%d", candidates.List()[0]), nil
}

func CompiledProgram(A int) []int {
	var result []int
	for {
		L0 := A % 8
		L2 := L0 ^ 6
		L4 := A / int(math.Pow(2, float64(L2)))
		L6 := L2 ^ L4
		L8 := L6 ^ 7
		A = A / 8
		result = append(result, L8%8)
		if A == 0 {
			return result
		}
	}
}

func (c *Computer) Run() error {
	for {
		err := c.Execute()
		// fmt.Printf("A: %d, B: %d, C: %d, Pointer: %d\n", c.RegisterA, c.RegisterB, c.RegisterC, c.Program.Pointer)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

func (c *Computer) Execute() error {
	instruction, err := c.Read()
	if err != nil {
		return io.EOF
	}

	switch instruction {
	case 0:
		err := c.executeAdv()
		if err != nil {
			return fmt.Errorf("error executing ADV: %w", err)
		}
	case 1:
		err := c.executeBxl()
		if err != nil {
			return fmt.Errorf("error executing BXL: %w", err)
		}
	case 2:
		err := c.executeBst()
		if err != nil {
			return fmt.Errorf("error executing BST: %w", err)
		}
	case 3:
		err := c.executeJnz()
		if err != nil {
			return fmt.Errorf("error executing JNZ: %w", err)
		}
	case 4:
		err := c.executeBxc()
		if err != nil {
			return fmt.Errorf("error executing BXC: %w", err)
		}
	case 5:
		err := c.executeOut()
		if err != nil {
			return fmt.Errorf("error executing OUT: %w", err)
		}
	case 6:
		err := c.executeBdv()
		if err != nil {
			return fmt.Errorf("error executing BDV: %w", err)
		}
	case 7:
		err := c.executeCdv()
		if err != nil {
			return fmt.Errorf("error executing CDV: %w", err)
		}
	default:
		return fmt.Errorf("unknown instruction %d", instruction)
	}
	return nil
}

func (c *Computer) executeAdv() error {
	// oldA := c.RegisterA
	res := c.division(&c.RegisterA)
	// fmt.Printf("ADV: oldA=%d, newA=%d\n", oldA, c.RegisterA)
	return res
}

func (c *Computer) executeBxl() error {
	left := c.RegisterB
	right, err := c.loadLiteralOperand()
	if err != nil {
		return err
	}
	xor := left ^ right
	// fmt.Printf("BXL: left=%d, right=%d, B=%d\n", left, right, xor)
	c.RegisterB = xor
	return nil
}

func (c *Computer) executeBst() error {
	op, err := c.loadComboOperand()
	if err != nil {
		return err
	}
	c.RegisterB = op % 8
	// fmt.Printf("BST: op=%v, B=%v\n", op, c.RegisterB)
	return nil
}

func (c *Computer) executeJnz() error {
	op, err := c.loadLiteralOperand()
	if err != nil {
		return err
	}

	if c.RegisterA == 0 {
		return nil
	}

	c.Program.Pointer = op
	// panic("jnz")
	return nil
}

func (c *Computer) executeBxc() error {
	_, err := c.loadLiteralOperand()
	if err != nil {
		return err
	}
	c.RegisterB = c.RegisterC ^ c.RegisterB
	return nil
}

func (c *Computer) executeOut() error {
	op, err := c.loadComboOperand()
	if err != nil {
		return err
	}
	res := op % 8
	// fmt.Printf("OUT: op=%v, res=%v\n", op, res)
	c.Output = append(c.Output, byte(res))
	return nil
}

func (c *Computer) executeBdv() error {
	return c.division(&c.RegisterB)
}

func (c *Computer) executeCdv() error {
	// fmt.Printf("CDV: C=%d\n", c.RegisterC)
	return c.division(&c.RegisterC)
}

func (c *Computer) division(target *int) error {
	num := c.RegisterA
	op, err := c.loadComboOperand()
	if err != nil {
		return err
	}
	denom := int(math.Pow(2, float64(op)))
	result := num / denom
	// fmt.Printf("DIV: num=%d, denom=%d, res=%d\n", num, int(denom), result)
	*target = result
	return nil
}

type Program struct {
	Memory  []byte
	Pointer int
}

func (c *Computer) Read() (byte, error) {
	if c.Program.Pointer >= len(c.Program.Memory) {
		return 0, ErrEndOfInput
	}
	value := c.Program.Memory[c.Program.Pointer]
	c.Program.Pointer++
	return value, nil
}

func (c *Computer) loadLiteralOperand() (int, error) {
	operand, err := c.Read()
	if err != nil {
		return 0, err
	}
	return int(operand), nil
}

func (c *Computer) loadComboOperand() (int, error) {
	v, err := c.Read()
	if err != nil {
		return 0, err
	}

	switch v {
	case 0, 1, 2, 3:
		return int(v), nil
	case 4:
		return c.RegisterA, nil
	case 5:
		return c.RegisterB, nil
	case 6:
		return c.RegisterC, nil
	default:
		return 0, fmt.Errorf("invalid combo operand %d", v)
	}
}
