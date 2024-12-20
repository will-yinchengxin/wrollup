package cmd

import (
	"wrollup/pkg/es"
	"wrollup/wtools"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get rollup job information",
	Long:  `Get information about one or all rollup jobs`,
	Run: func(cmd *cobra.Command, args []string) {
		client := es.NewClient(esURL)
		if jobName != "" {
			if err := client.GetRollupJob(jobName); err != nil {
				wtools.Error("Failed to get rollup job: " + err.Error())
				return
			}
		} else {
			if err := client.GetAllRollupJobs(); err != nil {
				wtools.Error("Failed to get all rollup jobs: " + err.Error())
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
