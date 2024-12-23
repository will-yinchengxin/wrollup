package cmd

import (
	"fmt"
	"wrollup/pkg/es"
	"wrollup/wtools"

	"github.com/spf13/cobra"
)

var (
	duration string
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean old data from indices",
	Long:  `Delete data older than three months from specified indices`,
	Run: func(cmd *cobra.Command, args []string) {
		if indice == "" {
			wtools.Error("Indice pattern is required")
			return
		}
		if _, ok := mapping[indice]; !ok {
			wtools.Error(fmt.Sprintf("Indice「%s」is not exist on local map", indice))
			return
		}

		client := es.NewClient(esURL)

		// 提示功能, 可选
		fmt.Printf("Are you sure you want to delete data older than 3 months from %s? (y/N): ", indice)
		var response string
		fmt.Scanln(&response)
		if response != "y" && response != "Y" {
			fmt.Println("Operation cancelled")
			return
		}
		if duration == "" {
			duration = "3M"
		}
		if err := client.DeleteOldData(indice, duration); err != nil {
			wtools.Error("Failed to clean old data: " + err.Error())
			return
		}

		wtools.Info(fmt.Sprintf("Successfully cleaned old data from %s", indice))
	},
}

func init() {
	cleanCmd.Flags().StringVarP(&indice, "indice", "i", "", "The index pattern to clean (required)")
	cleanCmd.Flags().StringVarP(&duration, "duration", "d", "", "The duration you want to clean, the format is like 1h/H/d/D/w/W/m/M/y/Y")
	rootCmd.AddCommand(cleanCmd)
}
