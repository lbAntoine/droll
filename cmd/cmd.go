package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/lbAntoine/droll/internal/commands"
	"github.com/spf13/cobra"
)

var (
	diceNumber int
	diceType   int
	showSum    bool
	showUnit   bool
	verbose    bool
)

var rootCmd = &cobra.Command{
	Use: `droll [--number, -n dice_number] --dice, -d dice_type [options]
  droll dice_number dice_type [options]
  droll dice_type [options]
  droll --help, -h`,
	Short: "dRoll is a CLI dice roller: v0.1.1",
	Long: `A Fast and Easy to use CLI dice roller built
with love and passion by lbAntoine in Go.
Complete documentation is available at https://github.com/lbAntoine/droll`,
	Version: "0.1.1",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}

		processCmd(args)
	},
}

func processCmd(args []string) {
	parseArgs(args)

	if diceNumber == 0 {
		diceNumber = 1
	}

	if diceType <= 0 {
		fmt.Println("Error: Dice must be a positive number > 0")
		os.Exit(1)
	}

	results := rollDice(diceNumber, diceType)
	displayResults(results)
}

func parseArgs(args []string) {
	if diceNumber == 0 && diceType == 0 {
		if len(args) == 1 {
			if d, err := strconv.Atoi(args[0]); err == nil {
				diceNumber = 1
				diceType = d
			}
		}

		if len(args) >= 2 {
			if n, err := strconv.Atoi(args[0]); err == nil {
				diceNumber = n
			}
			if d, err := strconv.Atoi(args[1]); err == nil {
				diceType = d
			}
		}
	}
}

func rollDice(number, diceType int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	results := make([]int, number)
	for i := range number {
		results[i] = r.Intn(diceType) + 1
	}

	return results
}

func calculateSum(results []int) int {
	sum := 0
	for _, r := range results {
		sum += r
	}

	return sum
}

func displayResults(results []int) {
	sum := calculateSum(results)

	fmt.Printf("%d dice rolled (d%d):\n", diceNumber, diceType)

	if diceNumber == 1 {
		fmt.Printf("Rolled: %d\n", results[0])
		if verbose {
			printVerboseMessage(results)
		}
		return
	}

	if showSum && !showUnit {
		fmt.Printf("Sum: %d\n", sum)
		if verbose {
			printVerboseMessage(results)
		}
		return
	}

	if showUnit && !showSum {
		fmt.Println("Throws:")
		for _, result := range results {
			fmt.Printf("  * %d\n", result)
		}
		if verbose {
			printVerboseMessage(results)
		}
		return
	}

	fmt.Printf("Sum: %d\n", sum)
	fmt.Println("Throws:")
	for _, result := range results {
		fmt.Printf("  * %d\n", result)
	}

	if verbose {
		printVerboseMessage(results)
	}
}

func printVerboseMessage(results []int) {
	if diceType == 20 {
		for _, result := range results {
			switch result {
			case 20:
				fmt.Printf("Natural 20!! Critical success!!!\n")
				return
			case 1:
				fmt.Printf("Natural 1... Critical failure!!!\n")
				return
			}
		}
	}

	sum := calculateSum(results)
	maxPossible := diceNumber * diceType
	percentage := float64(sum) / float64(maxPossible)

	switch {
	case percentage >= 0.9:
		fmt.Printf("Exceptional roll! The gods smile upon you!\n")
	case percentage >= 0.7:
		fmt.Printf("Good roll! Fortune favors you today.\n")
	case percentage >= 0.5:
		fmt.Printf("Decent roll. Could be better, could be worse.\n")
	case percentage >= 0.3:
		fmt.Printf("Not great. The fates are testing you.\n")
	default:
		fmt.Printf("Abysmal roll!! Better luck next time... I guess?\n")
	}
}

func init() {
	rootCmd.AddCommand(commands.NewFlipCommand())

	rootCmd.Flags().IntVarP(&diceNumber, "number", "n", 0, "Number of dice to roll")
	rootCmd.Flags().IntVarP(&diceType, "dice", "d", 0, "Type of dice to roll (e.g., 6 for d6)")
	rootCmd.Flags().BoolVar(&showSum, "sum", false, "Only show the sum of dice")
	rootCmd.Flags().BoolVar(&showUnit, "unit", false, "Only show individual dice throws")
	rootCmd.Flags().BoolVarP(&verbose, "comment", "c", false, "Show verbose DnD-style messages")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
