package cmd

import (
	"wrollup/pkg/es"
	"wrollup/wtools"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a rollup job",
	Long:  `Start an existing rollup job by name`,
	Run: func(cmd *cobra.Command, args []string) {
		if jobName == "" {
			wtools.Error("Job name is required")
			return
		}
		client := es.NewClient(esURL)
		if err := client.StartRollupJob(jobName); err != nil {
			wtools.Error("Failed to start rollup job: " + err.Error())
			return
		}

		wtools.Info("Successfully started rollup job: " + jobName)
	},
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop a rollup job",
	Long:  `Stop an existing rollup job by name`,
	Run: func(cmd *cobra.Command, args []string) {
		if jobName == "" {
			wtools.Error("Job name is required")
			return
		}
		client := es.NewClient(esURL)
		if err := client.StopRollupJob(jobName); err != nil {
			wtools.Error("Failed to stop rollup job: " + err.Error())
			return
		}

		wtools.Info("Successfully stopped rollup job: " + jobName)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(stopCmd)
}
