package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	Left  []int
	Right map[int]int
}

func main() {
	input := Input{
		Left:  []int{},
		Right: make(map[int]int),
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		pair := strings.Split(scanner.Text(), "   ")
		left, err := strconv.Atoi(pair[0])
		if err != nil {
			panic(err)
		}
		right, err := strconv.Atoi(pair[1])
		if err != nil {
			panic(err)
		}
		input.Left = append(input.Left, left)
		input.Right[right] += 1
	}

	similarity := 0
	for _, left := range input.Left {
		similarity += input.Right[left] * left
	}

	fmt.Println(similarity)
}
