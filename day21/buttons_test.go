package day21_test

import (
	"fmt"
	"testing"

	"github.com/makarchuk/aoc2024/day21"
	"github.com/stretchr/testify/require"
)

func TestSingleRobotE2E(t *testing.T) {
	type testCase struct {
		combination  string
		pressesCount int
	}

	for _, tc := range []testCase{
		{
			combination:  "029A",
			pressesCount: len("<A^A>^^AvvvA"),
		},
	} {
		testName := fmt.Sprintf("combination %s", tc.combination)
		t.Run(testName, func(t *testing.T) {
			// directPresser := day21.NewManualPresser()

			// indirectPresser := day21.NewRobotPresser(
			// 	day21.NumberPadButtons,
			// 	directPresser,
			// 	"only",
			// )

			// total := []byte{}

			// for _, c := range tc.combination {
			// 	res := indirectPresser.Press(byte(c))
			// 	total = append(total, res...)
			// }

			// require.Len(t, total, tc.pressesCount, day21.PrintPresses(total))

			robot := day21.BuildPrecalculatedCostProvider(
				*day21.NewPad(day21.NumberPadButtons),
				day21.ManualPresser{},
			)
			precalculatedRes := day21.EnterCombinationPrecalculated(tc.combination, robot)
			require.Equal(t, tc.pressesCount, precalculatedRes)
		})
	}
}

func TestTwoRobotsE2E(t *testing.T) {
	type testCase struct {
		combination  string
		pressesCount int
	}

	for _, tc := range []testCase{
		{
			combination:  "029A",
			pressesCount: len("v<<A>>^A<A>AvA<^AA>A<vAAA>^A"),
		},
	} {
		testName := fmt.Sprintf("combination %s", tc.combination)
		t.Run(testName, func(t *testing.T) {
			// directPresser := day21.NewManualPresser()

			// radiationRobot := day21.NewRobotPresser(
			// 	day21.ArrowPadButtons,
			// 	directPresser,
			// 	"banner",
			// )

			// freezingRobot := day21.NewRobotPresser(
			// 	day21.NumberPadButtons,
			// 	radiationRobot,
			// 	"frosty",
			// )

			// total := []byte{}

			// for _, c := range tc.combination {
			// 	res := freezingRobot.Press(byte(c))
			// 	total = append(total, res...)
			// }
			// require.Len(t, total, tc.pressesCount, day21.PrintPresses(total))

			robot := day21.BuildPrecalculatedCostProvider(
				*day21.NewPad(day21.ArrowPadButtons),
				day21.ManualPresser{},
			)
			fmt.Println("===============")
			robot2 := day21.BuildPrecalculatedCostProvider(
				*day21.NewPad(day21.NumberPadButtons),
				robot,
			)

			precalculatedRes := day21.EnterCombinationPrecalculated(tc.combination, robot2)
			require.Equal(t, tc.pressesCount, precalculatedRes)
		})
	}
}

func TestThreeRobotsE2E(t *testing.T) {
	type testCase struct {
		combination  string
		pressesCount int
	}

	for _, tc := range []testCase{
		{
			combination:  "029A",
			pressesCount: len("<vA<AA>>^AvAA<^A>A<v<A>>^AvA^A<vA>^A<v<A>^A>AAvA^A<v<A>A>^AAAvA<^A>A"),
		},
		{
			combination:  "980A",
			pressesCount: len("<v<A>>^AAAvA^A<vA<AA>>^AvAA<^A>A<v<A>A>^AAAvA<^A>A<vA>^A<A>A"),
		},
		{
			combination:  "179A",
			pressesCount: len("<v<A>>^A<vA<A>>^AAvAA<^A>A<v<A>>^AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A"),
		},
		{
			combination:  "456A",
			pressesCount: len("<v<A>>^AA<vA<A>>^AAvAA<^A>A<vA>^A<A>A<vA>^A<A>A<v<A>A>^AAvA<^A>A"),
		},
		{
			combination:  "379A",
			pressesCount: 64, //len("<v<A>>^AvA^A<vA<AA>>^AAvA<^A>AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A"),
		},
	} {
		testName := fmt.Sprintf("combination %s", tc.combination)
		t.Run(testName, func(t *testing.T) {
			robot := day21.BuildPrecalculatedCostProvider(
				*day21.NewPad(day21.ArrowPadButtons),
				day21.ManualPresser{},
			)
			fmt.Println("===============")
			robot2 := day21.BuildPrecalculatedCostProvider(
				*day21.NewPad(day21.ArrowPadButtons),
				robot,
			)
			robot3 := day21.BuildPrecalculatedCostProvider(
				*day21.NewPad(day21.NumberPadButtons),
				robot2,
			)

			precalculatedRes := day21.EnterCombinationPrecalculated(tc.combination, robot3)
			require.Equal(t, tc.pressesCount, precalculatedRes)
		})
	}
}
