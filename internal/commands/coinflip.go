package commands

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

func NewFlipCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "flip [number]",
		Short: "Flip a coin",
		Run: func(cmd *cobra.Command, args []string) {
			numFlips := 1
			if len(args) > 0 {
				if n, err := strconv.Atoi(args[0]); err == nil && n > 0 {
					numFlips = n
				} else {
					fmt.Println("Error: Invalid number of flips")
					os.Exit(1)
				}
			}

			results := flipCoin(numFlips)
			displayFlipResults(results)
		},
	}
}

func flipCoin(numFlips int) []string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	results := make([]string, numFlips)

	for i := range numFlips {
		if r.Intn(2) == 0 {
			results[i] = "Heads"
		} else {
			results[i] = "Tails"
		}
	}

	return results
}

func displayFlipResults(results []string) {
	fmt.Printf("%d coin flips:\n", len(results))
	for i, res := range results {
		fmt.Printf("  - Flip %d: %s\n", i+1, res)
	}
}
