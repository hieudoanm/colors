// Package cmd ...
package cmd

import (
	"colors-cli/utils/colors"
	"colors-cli/utils/figlet"
	"fmt"

	"github.com/spf13/cobra"
)

// hexToRgbCmd represents the colorsHexToRgb command
var hexToRgbCmd = &cobra.Command{
	Use:   "hex2rgb",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		figlet.LogProgramName()
		// Get URL
		fmt.Print("HEX: ")
		var hex string
		fmt.Scanln(&hex)
		r, g, b, err := colors.HexToRgb(hex)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("RGB: rgb(%d, %d, %d)\n", r, g, b)
	},
}

func init() {
	rootCmd.AddCommand(hexToRgbCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hexToRgbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hexToRgbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
