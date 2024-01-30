package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	// Replace with your module path
)

var exportFilePath string

var memCmd = &cobra.Command{
	Use:   "memory",
	Short: "Get Memory usage",
	Run: func(cmd *cobra.Command, args []string) {
		result := GetMemoryUsage()
		fmt.Println(result)
		ExportToFile(result, exportFilePath)
	},
}

func init() {
	rootCmd.AddCommand(memCmd)
	memCmd.Flags().StringVarP(&exportFilePath, "export", "e", "", "Export to file (provide file path)")
}
