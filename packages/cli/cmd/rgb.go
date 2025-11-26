// Package cmd ...
package cmd

import (
	"bufio"
	"colors-cli/utils/colors"
	"colors-cli/utils/figlet"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// rgbCmd represents the RGB-to-other-color-spaces command
var rgbCmd = &cobra.Command{
	Use:   "rgb",
	Short: "Convert RGB color to HEX, HSL, HCL, OKLCH, and CMYK",
	Long: `Convert an RGB color to multiple color spaces:
- HEX
- HSL (Hue, Saturation, Lightness)
- HCL (Hue, Chroma, Lightness)
- OKLCH (Lightness, Chroma, Hue)
- CMYK (Cyan, Magenta, Yellow, Black)

Example:
  colors-cli rgb`,
	Run: func(cmd *cobra.Command, args []string) {
		figlet.LogProgramName()

		reader := bufio.NewReader(os.Stdin)

		// Prompt for R
		fmt.Print("R (0-255): ")
		rStr, _ := reader.ReadString('\n')
		rStr = strings.TrimSpace(rStr)
		r, err := strconv.Atoi(rStr)
		if err != nil {
			fmt.Println("Error (R)    :", err)
			return
		}

		// Prompt for G
		fmt.Print("G (0-255): ")
		gStr, _ := reader.ReadString('\n')
		gStr = strings.TrimSpace(gStr)
		g, err := strconv.Atoi(gStr)
		if err != nil {
			fmt.Println("Error (G)    :", err)
			return
		}

		// Prompt for B
		fmt.Print("B (0-255): ")
		bStr, _ := reader.ReadString('\n')
		bStr = strings.TrimSpace(bStr)
		b, err := strconv.Atoi(bStr)
		if err != nil {
			fmt.Println("Error (B)    :", err)
			return
		}

		rgb := colors.RGB{R: r, G: g, B: b}

		// RGB → HEX
		hexStr, err := rgb.ToHex()
		if err != nil {
			fmt.Println("Error (HEX)  :", err)
		} else {
			fmt.Printf("HEX    : %s\n", hexStr)
		}

		// RGB → HSL
		hHSL, sHSL, lHSL, err := rgb.ToHSL()
		if err != nil {
			fmt.Println("Error (HSL)  :", err)
		} else {
			fmt.Printf("HSL    : h=%.2f°, s=%.2f%%, l=%.2f%%\n", hHSL, sHSL, lHSL)
		}

		// RGB → HCL
		hHCL, cHCL, lHCL, err := rgb.ToHCL()
		if err != nil {
			fmt.Println("Error (HCL)  :", err)
		} else {
			fmt.Printf("HCL    : h=%.2f°, c=%.2f, l=%.2f\n", hHCL, cHCL, lHCL)
		}

		// RGB → OKLCH
		oklL, oklC, oklH, err := rgb.ToOKLCH()
		if err != nil {
			fmt.Println("Error (OKLCH):", err)
		} else {
			fmt.Printf("OKLCH  : L=%.3f, C=%.3f, H=%.2f°\n", oklL, oklC, oklH)
		}

		// RGB → CMYK
		C, M, Y, K, err := rgb.ToCMYK()
		if err != nil {
			fmt.Println("Error (CMYK) :", err)
		} else {
			fmt.Printf("CMYK   : C=%.3f, M=%.3f, Y=%.3f, K=%.3f\n", C, M, Y, K)
		}
	},
}

func init() {
	rootCmd.AddCommand(rgbCmd)
}
