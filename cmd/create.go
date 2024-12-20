package cmd

import (
	"encoding/json"
	"fmt"
	"wrollup/pkg/es"
	"wrollup/wtools"

	"github.com/spf13/cobra"
)

var indice string

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new rollup job",
	Long:  `Create a new rollup job with the specified configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		if jobName == "" {
			wtools.Error("Job name is required")
			return
		}
		client := es.NewClient(esURL)
		jobExist, err := client.CheckRollupJob(jobName)
		if err != nil {
			wtools.Error("Failed to check rollup job: " + err.Error())
			return
		}
		if jobExist {
			wtools.Error("Rollup job already exists: " + jobName)
			return
		}
		if indice == "" {
			wtools.Error("Config indice is required")
			return
		}
		configData, ok := mapping[indice]
		if !ok {
			wtools.Error("Invalid config indice: " + indice + ", the indice we have are 「vsd」")
			return
		}

		var c map[string]interface{}
		if err := json.Unmarshal([]byte(configData), &c); err != nil {
			wtools.Error("Failed to parse config file: " + err.Error())
			return
		}
		if err := client.PutRollupJob(jobName, c); err != nil {
			wtools.Error(fmt.Sprintf("Failed to create rollup job[%s]: ", jobName) + err.Error())
			return
		}
		err = client.StartRollupJob(jobName)
		if err != nil {
			wtools.Error(fmt.Sprintf("Failed to start rollup job[%s]: ", jobName) + err.Error())
			return
		}
		wtools.Info("Successfully created rollup job: " + jobName)
	},
}

func init() {
	createCmd.Flags().StringVarP(&indice, "indice", "i", "", "The rollup job configuration indice")
	rootCmd.AddCommand(createCmd)
}
