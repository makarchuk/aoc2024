package main

import (
	"fmt"
	"os"

	"github.com/makarchuk/aoc2024/day9"
)

func main() {
	dm, err := day9.New(os.Stdin)
	if err != nil {
		panic(err)
	}
	defrag := dm.Defragmenter()

	fmt.Println(day9.CheckSum(defrag.Defragment()))
}
