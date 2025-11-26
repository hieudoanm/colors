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
	Short: "Convert HEX color to RGB, HSL, HCL, OKLCH, and CMYK",
	Long: `Convert a HEX color to multiple color spaces:
- RGB (Red, Green, Blue)
- HSL (Hue, Saturation, Lightness)
- HCL (Hue, Chroma, Lightness)
- OKLCH (Lightness, Chroma, Hue)
- CMYK (Cyan, Magenta, Yellow, Black)

Example:
  colors-cli hex #FF5733`,
	Run: func(cmd *cobra.Command, args []string) {
		figlet.LogProgramName()

		// Prompt for HEX input
		fmt.Print("HEX: ")
		var hexInput string
		fmt.Scanln(&hexInput)

		hex := colors.Hex(hexInput) // wrap input in Hex type

		// HEX → RGB
		r, g, b, err := hex.ToRGB()
		if err != nil {
			fmt.Println("Error (RGB)  :", err)
			return
		}
		fmt.Printf("RGB    : rgb(%d, %d, %d)\n", r, g, b)

		// HEX → HSL
		hHSL, sHSL, lHSL, err := hex.ToHSL()
		if err != nil {
			fmt.Println("Error (HSL)  :", err)
		} else {
			fmt.Printf("HSL    : h=%.2f°, s=%.2f%%, l=%.2f%%\n", hHSL, sHSL, lHSL)
		}

		// HEX → HCL
		hHCL, cHCL, lHCL, err := hex.ToHCL()
		if err != nil {
			fmt.Println("Error (HCL)  :", err)
		} else {
			fmt.Printf("HCL    : h=%.2f°, c=%.2f, l=%.2f\n", hHCL, cHCL, lHCL)
		}

		// HEX → OKLCH
		oklL, oklC, oklH, err := hex.ToOKLCH()
		if err != nil {
			fmt.Println("Error (OKLCH):", err)
		} else {
			fmt.Printf("OKLCH  : L=%.3f, C=%.3f, H=%.2f°\n", oklL, oklC, oklH)
		}

		// HEX → CMYK
		C, M, Y, K, err := hex.ToCMYK()
		if err != nil {
			fmt.Println("Error (CMYK) :", err)
		} else {
			fmt.Printf("CMYK   : C=%.3f, M=%.3f, Y=%.3f, K=%.3f\n", C, M, Y, K)
		}
	},
}

func init() {
	rootCmd.AddCommand(hexCmd)
}
