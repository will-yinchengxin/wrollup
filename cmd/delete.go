package cmd

import (
	"wrollup/pkg/es"
	"wrollup/wtools"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a rollup job",
	Long:  `Delete an existing rollup job by name`,
	Run: func(cmd *cobra.Command, args []string) {
		if jobName == "" {
			wtools.Error("Job name is required")
			return
		}
		client := es.NewClient(esURL)
		client.StopRollupJob(jobName)
		if err := client.DeleteRollupJob(jobName); err != nil {
			wtools.Error("Failed to delete rollup job: " + err.Error())
			return
		}

		wtools.Info("Successfully deleted rollup job: " + jobName)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
