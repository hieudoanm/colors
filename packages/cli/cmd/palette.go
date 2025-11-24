/*
Copyright © 2025
*/
package cmd

import (
	"bufio"
	"colors-cli/utils/colors"
	"fmt"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// paletteCmd represents the palette command
var paletteCmd = &cobra.Command{
	Use:   "palette",
	Short: "Generate a 3-color palette from a HEX input",
	Long: `Generate a 3-color palette using 1 of 3 design styles:
1. Balanced professional (Triadic)
2. High-contrast (Complementary)
3. Soft aesthetic (Analogous)`,
	Run: func(cmd *cobra.Command, args []string) {

		reader := bufio.NewReader(os.Stdin)

		// ----------------------------
		// 1) Ask for base HEX input
		// ----------------------------
		fmt.Print("Enter base HEX (e.g. #ff6600): ")
		hexInput, _ := reader.ReadString('\n')
		hexInput = strings.TrimSpace(hexInput)

		if !colors.IsValidHex(hexInput) {
			fmt.Println("❌ Invalid hex color.")
			return
		}

		// HEX → HSL
		h, s, l, err := colors.HexToHSL(hexInput)
		if err != nil {
			fmt.Println(err)
			return
		}
		baseHue := h
		baseS := s
		baseL := l

		// ----------------------------
		// 2) Ask for palette style
		// ----------------------------
		paletteOptions := []string{
			"Balanced professional (Triadic)",
			"High-contrast (Complementary)",
			"Soft aesthetic (Analogous)",
		}

		var selected string

		prompt := &survey.Select{
			Message: "Choose palette style:",
			Options: paletteOptions,
		}

		survey.AskOne(prompt, &selected)

		// ----------------------------
		// 3) Compute palette hues
		// ----------------------------
		var base, support, accent float64

		switch selected {
		case "Balanced professional (Triadic)":
			base = baseHue
			support = baseHue + 120
			accent = baseHue + 240

		case "High-contrast (Complementary)":
			base = baseHue
			support = baseHue + 30
			accent = baseHue + 180

		case "Soft aesthetic (Analogous)":
			base = baseHue
			support = baseHue - 30
			accent = baseHue + 30
		}

		// Normalize hues
		hues := []*float64{&base, &support, &accent}
		for _, h := range hues {
			for *h < 0 {
				*h += 360
			}
			for *h >= 360 {
				*h -= 360
			}
		}

		// ----------------------------
		// 4) Convert back to HEX
		// ----------------------------
		baseHex := colors.HSLToHex(colors.HSL{H: base, S: baseS, L: baseL})
		supportHex := colors.HSLToHex(colors.HSL{H: support, S: baseS, L: baseL})
		accentHex := colors.HSLToHex(colors.HSL{H: accent, S: baseS, L: baseL})

		// ----------------------------
		// 5) Output
		// ----------------------------
		fmt.Println("\n🎨 Generated Palette:")
		fmt.Printf("Base:    %s  (H=%.0f)\n", baseHex, base)
		fmt.Printf("Support: %s  (H=%.0f)\n", supportHex, support)
		fmt.Printf("Accent:  %s  (H=%.0f)\n", accentHex, accent)
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(paletteCmd)
}
