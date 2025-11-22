// Package cmd ...
package cmd

import (
	"colors-cli/utils/colors"
	"colors-cli/utils/figlet"
	"fmt"

	"github.com/spf13/cobra"
)

// hexCmd represents the colorsHexToRgb command
var hexCmd = &cobra.Command{
	Use:   "hex",
	Short: "Convert HEX color to RGB, HSL, HCL, and OKLCH",
	Long: `Convert a HEX color to multiple color spaces:
- RGB
- HSL (Hue, Saturation, Lightness)
- HCL (Hue, Chroma, Lightness)
- OKLCH (Lightness, Chroma, Hue)

Example:
  colors-cli hex #FF5733`,
	Run: func(cmd *cobra.Command, args []string) {
		figlet.LogProgramName()

		// Prompt for HEX input
		fmt.Print("HEX: ")
		var hex string
		fmt.Scanln(&hex)

		// HEX → RGB
		r, g, b, err := colors.HexToRgb(hex)
		if err != nil {
			fmt.Println("Error (RGB):", err)
			return
		}
		fmt.Printf("RGB: rgb(%d, %d, %d)\n", r, g, b)

		// HEX → HSL
		hHSL, sHSL, lHSL, err := colors.HexToHSL(hex)
		if err != nil {
			fmt.Println("Error (HSL):", err)
		} else {
			fmt.Printf("HSL: h=%.2f°, s=%.2f%%, l=%.2f%%\n", hHSL, sHSL, lHSL)
		}

		// HEX → HCL
		hHCL, cHCL, lHCL, err := colors.HexToHCL(hex)
		if err != nil {
			fmt.Println("Error (HCL):", err)
		} else {
			fmt.Printf("HCL: h=%.2f°, c=%.2f, l=%.2f\n", hHCL, cHCL, lHCL)
		}

		// HEX → OKLCH
		oklL, oklC, oklH, err := colors.HexToOKLCH(hex)
		if err != nil {
			fmt.Println("Error (OKLCH):", err)
		} else {
			fmt.Printf("OKLCH: L=%.3f, C=%.3f, H=%.2f°\n", oklL, oklC, oklH)
		}
	},
}

func init() {
	rootCmd.AddCommand(hexCmd)
}
