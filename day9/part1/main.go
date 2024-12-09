package main

import (
	"fmt"
	"os"

	"github.com/makarchuk/aoc2024/day9"
)

func main() {
	compressed, err := day9.New(os.Stdin)
	if err != nil {
		panic(err)
	}
	rendered := compressed.Render()
	day9.Defragment(rendered)
	// day9.PrintDisk(rendered)
	fmt.Println(day9.CheckSum(rendered))
}
