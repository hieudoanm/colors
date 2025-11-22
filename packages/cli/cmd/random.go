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
	Short: "Generate random HEX colors with conversions",
	Long: `Generate a specified number of random HEX colors and show their RGB values.
Example:
  colors-cli random --max 5`,
	Run: func(cmd *cobra.Command, args []string) {
		figlet.LogProgramName()

		for i := 0; i < maxHEX; i++ {
			newHEX := colors.GenerateRandomHexColor()
			hex := colors.Hex(newHEX) // wrap string in Hex type

			r, g, b, err := hex.ToRGB()
			if err != nil {
				fmt.Println("Error converting HEX to RGB:", err)
				continue
			}

			fmt.Printf("%s - rgb(%d, %d, %d)\n", newHEX, r, g, b)
		}
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

	randomCmd.PersistentFlags().IntVarP(&maxHEX, "max", "m", 1, "Number of Colors")
}
