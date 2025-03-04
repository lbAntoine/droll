package cmd

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestParseArgs(t *testing.T) {
	tcs := []struct {
		name            string
		args            []string
		expectedNumber  int
		expectedDice    int
		setupGlobalVars func()
	}{
		{
			name:           "Single argument sets dice type",
			args:           []string{"6"},
			expectedNumber: 1,
			expectedDice:   6,
			setupGlobalVars: func() {
				diceNumber = 0
				diceType = 0
			},
		},
		{
			name:           "Two arguments sets number and type",
			args:           []string{"3", "6"},
			expectedNumber: 3,
			expectedDice:   6,
			setupGlobalVars: func() {
				diceNumber = 0
				diceType = 0
			},
		},
		{
			name:           "Flags already set take precedence",
			args:           []string{"2", "20"},
			expectedNumber: 5,
			expectedDice:   10,
			setupGlobalVars: func() {
				diceNumber = 5
				diceType = 10
			},
		},
		{
			name:           "Big numbs for big bums",
			args:           []string{"500", "487"},
			expectedNumber: 500,
			expectedDice:   487,
			setupGlobalVars: func() {
				diceNumber = 0
				diceType = 0
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupGlobalVars()

			parseArgs(tc.args)

			if diceNumber != tc.expectedNumber {
				t.Errorf("Expected dice number %d, got %d", tc.expectedNumber, diceNumber)
			}
			if diceType != tc.expectedDice {
				t.Errorf("Expected dice type %d, got %d", tc.expectedDice, diceType)
			}
		})
	}
}

func TestRollDice(t *testing.T) {
	tcs := []struct {
		number   int
		diceType int
	}{
		{number: 1, diceType: 6},
		{number: 3, diceType: 20},
		{number: 5, diceType: 100},
		{number: 500, diceType: 487},
	}

	for _, tc := range tcs {
		t.Run(fmt.Sprintf("%dd%d", tc.number, tc.diceType), func(t *testing.T) {
			res := rollDice(tc.number, tc.diceType)

			if len(res) != tc.number {
				t.Errorf("Expected %d results, got %d", tc.number, len(res))
			}

			for _, roll := range res {
				if roll < 1 || roll > tc.diceType {
					t.Errorf("Roll %d out of range for d%d", roll, tc.diceType)
				}
			}
		})
	}
}

func TestCalculateSum(t *testing.T) {
	tcs := []struct {
		name     string
		results  []int
		expected int
	}{
		{
			name:     "Simple sum",
			results:  []int{1, 2, 3},
			expected: 6,
		},
		{
			name:     "Big nums",
			results:  []int{12039, 90012938, 882382366616},
			expected: 882472391593,
		},
		{
			name:     "Single die",
			results:  []int{4},
			expected: 4,
		},
		{
			name:     "Empty slice",
			results:  []int{},
			expected: 0,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			sum := calculateSum(tc.results)
			if sum != tc.expected {
				t.Errorf("Expected sum %d, got %d", tc.expected, sum)
			}
		})
	}
}

func TestDisplayResults(t *testing.T) {
	tcs := []struct {
		name            string
		results         []int
		setupGlobalVars func()
		expectedOutput  string
	}{
		{
			name:    "Single die roll",
			results: []int{5},
			setupGlobalVars: func() {
				diceNumber = 1
				diceType = 6
				showSum = false
				showUnit = false
				verbose = false
			},
			expectedOutput: "1 dice rolled (d6):\nRolled: 5\n",
		},
		{
			name:    "Double 20 roll",
			results: []int{19, 13},
			setupGlobalVars: func() {
				diceNumber = 2
				diceType = 20
				showSum = false
				showUnit = false
				verbose = false
			},
			expectedOutput: "2 dice rolled (d20):\nSum: 32\nThrows:\n  * 19\n  * 13\n",
		},
		{
			name:    "Big big big roll, only sum",
			results: []int{7, 18, 39, 32, 26, 38, 36, 48, 47, 7, 35, 45, 3, 8, 13, 35, 2, 45, 17, 48, 1, 1, 38, 13, 38, 41, 33, 39, 9, 25, 21, 9},
			setupGlobalVars: func() {
				diceNumber = 32
				diceType = 20
				showSum = true
				showUnit = false
				verbose = false
			},
			expectedOutput: "32 dice rolled (d20):\nSum: 817\n",
		},
		{
			name:    "Double roll, only throws",
			results: []int{19, 12},
			setupGlobalVars: func() {
				diceNumber = 2
				diceType = 20
				showSum = false
				showUnit = true
				verbose = false
			},
			expectedOutput: "2 dice rolled (d20):\nThrows:\n  * 19\n  * 13\n",
		},
		{
			name:    "Double roll, verbose",
			results: []int{19, 13},
			setupGlobalVars: func() {
				diceNumber = 2
				diceType = 20
				showSum = false
				showUnit = false
				verbose = true
			},
			expectedOutput: "2 dice rolled (d20):\nSum: 32\nThrows:\n  * 19\n  * 13\nGood roll! Fortune favors you today.\n",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			tc.setupGlobalVars()

			displayResults(tc.results)

			w.Close()
			os.Stdout = oldStdout
			var buf bytes.Buffer
			buf.ReadFrom(r)

			output := buf.String()
			if output != tc.expectedOutput {
				t.Errorf("Expected output:\n%s\nGot:\n%s", tc.expectedOutput, output)
			}
		})
	}
}

func TestPrintVerboseMessage(t *testing.T) {
	tcs := []struct {
		name     string
		result   []int
		diceType int
		expected string
	}{
		{
			name:     "D20 Critical Success",
			result:   []int{20},
			diceType: 20,
			expected: "Natural 20!! Critical success!!!\n",
		},
		{
			name:     "D20 Critical Miss",
			result:   []int{1},
			diceType: 20,
			expected: "Natural 1... Critical failure!!!\n",
		},
		{
			name:     "Exceptional roll",
			result:   []int{12, 12},
			diceType: 12,
			expected: "Exceptional roll! The gods smile upon you!\n",
		},
		{
			name:     "Good roll",
			result:   []int{10, 10},
			diceType: 12,
			expected: "Good roll! Fortune favors you today.\n",
		},
		{
			name:     "Decent roll",
			result:   []int{12, 1},
			diceType: 12,
			expected: "Decent roll. Could be better, could be worse.\n",
		},
		{
			name:     "Not great roll",
			result:   []int{5, 4},
			diceType: 12,
			expected: "Not great. The fates are testing you.\n",
		},
		{
			name:     "Abysmal roll",
			result:   []int{1, 1},
			diceType: 12,
			expected: "Abysmal roll!! Better luck next time... I guess?\n",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			diceType = tc.diceType
			verbose = true

			printVerboseMessage(tc.result)

			w.Close()
			os.Stdout = oldStdout
			var buf bytes.Buffer
			buf.ReadFrom(r)

			output := buf.String()
			if output != tc.expected {
				t.Errorf("Expected message:\n%s\nGot:\n%s", tc.expected, output)
			}
		})
	}
}
