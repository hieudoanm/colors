// Package cmd ...
package cmd

import (
	"colors-cli/utils/colors"
	"colors-cli/utils/figlet"
	"fmt"

	"github.com/spf13/cobra"
)

var maxHEX int

// randomCmd represents the colorsRandom command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		figlet.LogProgramName()
		for i := 0; i < maxHEX; i++ {
			newHEX := colors.GenerateRandomHexColor()
			r, g, b, err := colors.HexToRgb(newHEX)
			if err != nil {
				fmt.Println("Error:", err)
			}
			fmt.Printf("%s - rgb(%d, %d, %d)\n", newHEX, r, g, b)
		}
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

	randomCmd.PersistentFlags().IntVarP(&maxHEX, "max", "m", 1, "Number of Colors")
}
