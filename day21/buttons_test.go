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
			directPresser := day21.NewManualPresser()

			indirectPresser := day21.NewRobotPresser(
				day21.NumberPadButtons,
				directPresser,
				"only",
			)

			total := []byte{}

			for _, c := range tc.combination {
				res := indirectPresser.Press(byte(c))
				total = append(total, res...)
			}

			require.Len(t, total, tc.pressesCount, day21.PrintPresses(total))
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
			directPresser := day21.NewManualPresser()

			radiationRobot := day21.NewRobotPresser(
				day21.ArrowPadButtons,
				directPresser,
				"banner",
			)

			freezingRobot := day21.NewRobotPresser(
				day21.NumberPadButtons,
				radiationRobot,
				"frosty",
			)

			total := []byte{}

			for _, c := range tc.combination {
				res := freezingRobot.Press(byte(c))
				total = append(total, res...)
			}

			require.Len(t, total, tc.pressesCount, day21.PrintPresses(total))
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
			directPresser := day21.NewManualPresser()

			historianRobot := day21.NewRobotPresser(
				day21.ArrowPadButtons,
				directPresser,
				"historian",
			)

			frostyRobot := day21.NewRobotPresser(
				day21.ArrowPadButtons,
				historianRobot,
				"frosty",
			)

			radiationRobot := day21.NewRobotPresser(
				day21.NumberPadButtons,
				frostyRobot,
				"banner",
			)

			total := []byte{}

			for _, c := range tc.combination {
				res := radiationRobot.Press(byte(c))
				total = append(total, res...)
			}

			require.Len(t, total, tc.pressesCount, day21.PrintPresses(total))
		})
	}
}

func TestSingleArrowsPadBot(t *testing.T) {
	type testCase struct {
		from  byte
		to    byte
		moves int
	}

	for _, tc := range []testCase{{
		from:  '<',
		to:    'A',
		moves: len(">>^A"),
	}} {
		robot := day21.NewRobotPresser(
			day21.ArrowPadButtons,
			day21.ManualPresser{},
			"indirect",
		)

		//discard the value, just moving to the starting position
		_ = robot.Press(tc.from)
		res := robot.Press(tc.to)
		require.Len(t, res, tc.moves, day21.PrintPresses(res))
	}
}

func TestRobotPresser(t *testing.T) {
	type testCase struct {
		from   byte
		to     byte
		length int
		result string
	}

	for _, tc := range []testCase{
		{
			from:   '0',
			to:     '1',
			length: 3,
			result: "^<A",
		},
		{
			from:   '0',
			to:     '9',
			length: 5,
		},
		{
			from:   'A',
			to:     'A',
			length: 1,
			result: "A",
		},
	} {
		testName := fmt.Sprintf("direct %c-%c", tc.from, tc.to)
		t.Run(testName, func(t *testing.T) {
			directPresser := day21.NewManualPresser()

			indirectPresser := day21.NewRobotPresser(
				day21.NumberPadButtons,
				directPresser,
				"only",
			)

			//discard the value, just moving to the starting position
			_ = indirectPresser.Press(tc.from)
			res := indirectPresser.Press(tc.to)

			require.Len(t, res, tc.length, day21.PrintPresses(res))
			if tc.result != "" {
				require.Equal(t, tc.result, string(res), day21.PrintPresses(res))
			}
		})
	}

}

func TestDirect(t *testing.T) {
	type testCase struct {
		from   byte
		to     byte
		length int
	}

	testCases := []testCase{}
	for _, i := range []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0', 'A'} {
		for _, j := range []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0', 'A'} {
			testCases = append(testCases, testCase{
				from:   i,
				to:     j,
				length: 1,
			})
		}
	}
	for _, tc := range testCases {
		testName := fmt.Sprintf("direct %c-%c", tc.from, tc.to)
		t.Run(testName, func(t *testing.T) {
			presser := day21.NewManualPresser()

			//discard the value, just moving to the starting position
			_ = presser.Press(tc.from)
			res := presser.Press(tc.to)

			require.Len(t, res, tc.length, day21.PrintPresses(res))
		})
	}
}
