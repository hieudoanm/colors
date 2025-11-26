// Package cmd ...
package cmd

import (
	"colors-cli/utils/colors"
	"colors-cli/utils/figlet"
	"fmt"

	"github.com/spf13/cobra"
)

// hclCmd represents the colorsHCL command
var hclCmd = &cobra.Command{
	Use:   "hcl",
	Short: "Convert HCL color to HEX, RGB, HSL, OKLCH, and CMYK",
	Long: `Convert a HCL color (Hue, Chroma, Lightness) to multiple color spaces:
- HEX (Hexadecimal)
- RGB (Red, Green, Blue)
- HSL (Hue, Saturation, Lightness)
- OKLCH (Lightness, Chroma, Hue)
- CMYK (Cyan, Magenta, Yellow, Black)

Example:
  colors-cli hcl 30 80 50`,
	Run: func(cmd *cobra.Command, args []string) {
		figlet.LogProgramName()

		// Prompt for HCL input
		var h, c, l float64
		fmt.Print("Hue (0–360)      : ")
		fmt.Scanln(&h)
		fmt.Print("Chroma (0–100)   : ")
		fmt.Scanln(&c)
		fmt.Print("Lightness (0–100): ")
		fmt.Scanln(&l)

		hcl := colors.HCL{H: h, C: c, L: l}

		// HCL → RGB
		r, g, b, err := hcl.ToRGB()
		if err != nil {
			fmt.Println("Error (RGB)  :", err)
			return
		}
		rgb := colors.RGB{R: r, G: g, B: b}

		// RGB → HEX
		hex, err := rgb.ToHex()
		if err != nil {
			fmt.Println("Error (HEX)  :", err)
		} else {
			fmt.Printf("HEX    : %s\n", hex)
		}

		// RGB → RGB
		if !rgb.IsValid() {
			fmt.Println("Error (RGB)  : invalid RGB values")
		} else {
			fmt.Printf("RGB    : rgb(%d, %d, %d)\n", r, g, b)
		}

		// RGB → HSL
		hHSL, sHSL, lHSL, err := rgb.ToHSL()
		if hHSL < 0 || sHSL < 0 || lHSL < 0 || err != nil {
			fmt.Println("Error (HSL)  : conversion failed")
		} else {
			fmt.Printf("HSL    : h=%.2f°, s=%.2f%%, l=%.2f%%\n", hHSL, sHSL, lHSL)
		}

		// RGB → OKLCH
		oklL, oklC, oklH, err := rgb.ToOKLCH()
		if oklL < 0 || oklC < 0 || oklH < 0 || err != nil {
			fmt.Println("Error (OKLCH): conversion failed")
		} else {
			fmt.Printf("OKLCH  : L=%.3f, C=%.3f, H=%.2f°\n", oklL, oklC, oklH)
		}

		// RGB → CMYK
		C, M, Y, K, err := rgb.ToCMYK()
		if C < 0 || M < 0 || Y < 0 || K < 0 || err != nil {
			fmt.Println("Error (CMYK) : conversion failed")
		} else {
			fmt.Printf("CMYK   : C=%.3f, M=%.3f, Y=%.3f, K=%.3f\n", C, M, Y, K)
		}
	},
}

func init() {
	rootCmd.AddCommand(hclCmd)
}
