package cmd

import (
	"fmt"
	"os"

	"github.com/lbAntoine/droll/internal/commands"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "droll",
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
	},
}

func init() {
	rootCmd.AddCommand(commands.NewRollCommand())
	rootCmd.AddCommand(commands.NewFlipCommand())
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
