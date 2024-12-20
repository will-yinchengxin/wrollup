package cmd

import (
	"bufio"
	"fmt"
	"wrollup/pkg/es"
	"wrollup/wtools"

	"github.com/spf13/cobra"
	"os"
	"strings"
)

type Config struct {
	ElasticsearchServiceUrl string
}

var (
	esURL   string
	jobName string
	config  Config
)

var rootCmd = &cobra.Command{
	Use:   "WRollup",
	Short: "Elasticsearch Rollup job management tool",
	Long:  `A command line tool for managing Elasticsearch Rollup jobs`,
}

func Execute() error {
	return rootCmd.Execute()
}

func readConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	cfg := &Config{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		value = strings.Trim(value, "\"")
		if key == "elasticsearchServiceUrl" {
			cfg.ElasticsearchServiceUrl = value
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func init() {
	var configFile string

	if _, err := os.Stat("/usr/local/vs_conf/vda.conf"); err == nil {
		configFile = "/usr/local/vs_conf/vda.conf"
	} else {
		if _, err := os.Stat("./conf/vda.conf"); err == nil {
			configFile = "./conf/vda.conf"
		}
	}
	if configFile != "" {
		if cfg, err := readConfig(configFile); err == nil {
			config = *cfg
			esURL = config.ElasticsearchServiceUrl
		} else {
			fmt.Printf("Warning: Failed to read config file: %v\n", err)
			esURL = "http://localhost:9200"
		}
	} else {
		esURL = "http://localhost:9200"
	}

	if err := es.CheckESConnection(esURL); err != nil {
		wtools.Error(fmt.Sprintf("Warning: Elasticsearch connection check failed at %s: ", esURL) + err.Error())
		os.Exit(1)
	} else {
		fmt.Printf("Successfully connected to Elasticsearch at %s\n", esURL)
	}

	rootCmd.PersistentFlags().StringVar(&esURL, "url", esURL, "Elasticsearch URL")
	rootCmd.PersistentFlags().StringVar(&jobName, "job", "", "Rollup job name")
}
