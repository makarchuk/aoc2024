package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/makarchuk/aoc2024/pkg/orderedlist"
)

type Input struct {
	Left  orderedlist.List[int]
	Right orderedlist.List[int]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input Input
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
		input.Left.Add(left)
		input.Right.Add(right)
	}

	totalDistance := 0
	rightItems := input.Right.Items()
	for i, left := range input.Left.Items() {
		distance := rightItems[i] - left
		if distance < 0 {
			distance = -distance
		}
		totalDistance += distance
	}

	fmt.Println(totalDistance)
}
