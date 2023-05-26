/*
Copyright © 2023 NAZIF ENES KIZILCIN <enes.kizilcin@gmail.com>
*/

package cmd

import (
	"fmt"
	"github.com/eneskzlcn/jsn/clipboard"
	"github.com/eneskzlcn/jsn/formatter"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jsn",
	Short: "That application takes a json and formats it.",
	Long: `That application takes a json, formats it. 
After copies the output to the your clipboard while printing it to your console.
`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		jsonBytes := []byte(args[0])
		output, err := formatter.GetFormattedJSON(jsonBytes)
		if err != nil {
			return err
		}
		if err := clipboard.Add(output); err != nil {
			return err
		}

		fmt.Println(output)

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
