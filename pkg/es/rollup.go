package es

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"wrollup/wtools"
)

type Client struct {
	baseURL string
	http    *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		http:    &http.Client{},
	}
}

type RollupJobsResponse struct {
	Jobs []RollupJobWithID `json:"jobs"`
}

type RollupJobWithID struct {
	Config RollupConfig `json:"config"`
	Stats  JobStats     `json:"stats"`
	Status JobStatus    `json:"status"`
}

type RollupConfig struct {
	ID           string                 `json:"id"`
	IndexPattern string                 `json:"index_pattern"`
	RollupIndex  string                 `json:"rollup_index"`
	Cron         string                 `json:"cron"`
	PageSize     int                    `json:"page_size"`
	Groups       map[string]interface{} `json:"groups"`
	Metrics      []interface{}          `json:"metrics"`
	Timeout      string                 `json:"timeout"`
}

type JobStats struct {
	PagesProcessed     int64 `json:"pages_processed"`
	DocumentsProcessed int64 `json:"documents_processed"`
	RollupsIndexed     int64 `json:"rollups_indexed"`
	TriggerCount       int64 `json:"trigger_count"`
	IndexTimeInMs      int64 `json:"index_time_in_ms"`
	IndexTotal         int64 `json:"index_total"`
	IndexFailures      int64 `json:"index_failures"`
	SearchTimeInMs     int64 `json:"search_time_in_ms"`
	SearchTotal        int64 `json:"search_total"`
	SearchFailures     int64 `json:"search_failures"`
	ProcessingTimeInMs int64 `json:"processing_time_in_ms"`
	ProcessingTotal    int64 `json:"processing_total"`
}

type JobStatus struct {
	JobState      string `json:"job_state"`
	UpgradedDocID bool   `json:"upgraded_doc_id"`
}

func (c *Client) GetAllRollupJobs() error {
	resp, err := c.doRequest(http.MethodGet, "/_rollup/job/_all", nil)
	if err != nil {
		return err
	}

	var response RollupJobsResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		return fmt.Errorf("failed to parse response: %v", err)
	}
	fmt.Println("")
	table := tablewriter.NewWriter(os.Stdout)
	fmt.Println("-------------+---------------+--------------+--------+----------------+-----------------+------------------")
	table.SetHeader([]string{"Job ID", "Index Pattern", "Rollup Index", "Status", "Docs Processed", "Pages Processed", "Rollups Indexed"})
	table.SetBorder(false)
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.FgHiGreenColor},
		tablewriter.Colors{tablewriter.FgHiGreenColor},
		tablewriter.Colors{tablewriter.FgHiGreenColor},
		tablewriter.Colors{tablewriter.FgHiGreenColor},
		tablewriter.Colors{tablewriter.FgHiGreenColor},
		tablewriter.Colors{tablewriter.FgHiGreenColor},
		tablewriter.Colors{tablewriter.FgHiGreenColor},
	)

	for _, job := range response.Jobs {
		table.Append([]string{
			job.Config.ID,
			job.Config.IndexPattern,
			job.Config.RollupIndex,
			job.Status.JobState,
			strconv.FormatInt(job.Stats.DocumentsProcessed, 10),
			strconv.FormatInt(job.Stats.PagesProcessed, 10),
			strconv.FormatInt(job.Stats.RollupsIndexed, 10),
		})
	}
	table.Render()

	return nil
}

func (c *Client) GetRollupJob(name string) error {
	resp, err := c.doRequest(http.MethodGet, fmt.Sprintf("/_rollup/job/%s", name), nil)
	if err != nil {
		return err
	}

	var response RollupJobsResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		return fmt.Errorf("failed to parse response: %v", err)
	}

	var job RollupJobWithID
	found := false
	for _, j := range response.Jobs {
		if j.Config.ID == name {
			job = j
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("job %s not found", name)
	}
	fmt.Println("")
	fmt.Printf("\n🔍 Job Details for: %s\n\n", name)
	basicTable := tablewriter.NewWriter(os.Stdout)
	basicTable.SetHeader([]string{"Property", "Value"})
	basicTable.SetBorder(false)
	basicTable.SetHeaderColor(
		tablewriter.Colors{tablewriter.FgHiGreenColor},
		tablewriter.Colors{tablewriter.FgHiGreenColor},
	)

	basicTable.Append([]string{"Index Pattern", job.Config.IndexPattern})
	basicTable.Append([]string{"Rollup Index", job.Config.RollupIndex})
	basicTable.Append([]string{"Cron Schedule", job.Config.Cron})
	basicTable.Append([]string{"Page Size", strconv.Itoa(job.Config.PageSize)})
	basicTable.Append([]string{"Status", job.Status.JobState})
	basicTable.Append([]string{"Timeout", job.Config.Timeout})
	basicTable.Render()

	fmt.Printf("\n📊 Statistics:\n\n")
	statsTable := tablewriter.NewWriter(os.Stdout)
	statsTable.SetHeader([]string{"Metric", "Value"})
	statsTable.SetBorder(false)
	statsTable.SetHeaderColor(
		tablewriter.Colors{tablewriter.FgHiGreenColor},
		tablewriter.Colors{tablewriter.FgHiGreenColor},
	)

	statsTable.Append([]string{"Documents Processed", strconv.FormatInt(job.Stats.DocumentsProcessed, 10)})
	statsTable.Append([]string{"Pages Processed", strconv.FormatInt(job.Stats.PagesProcessed, 10)})
	statsTable.Append([]string{"Rollups Indexed", strconv.FormatInt(job.Stats.RollupsIndexed, 10)})
	statsTable.Append([]string{"Trigger Count", strconv.FormatInt(job.Stats.TriggerCount, 10)})
	statsTable.Append([]string{"Index Time (ms)", strconv.FormatInt(job.Stats.IndexTimeInMs, 10)})
	statsTable.Append([]string{"Index Total", strconv.FormatInt(job.Stats.IndexTotal, 10)})
	statsTable.Append([]string{"Index Failures", strconv.FormatInt(job.Stats.IndexFailures, 10)})
	statsTable.Append([]string{"Search Time (ms)", strconv.FormatInt(job.Stats.SearchTimeInMs, 10)})
	statsTable.Append([]string{"Search Total", strconv.FormatInt(job.Stats.SearchTotal, 10)})
	statsTable.Append([]string{"Search Failures", strconv.FormatInt(job.Stats.SearchFailures, 10)})
	statsTable.Append([]string{"Processing Time (ms)", strconv.FormatInt(job.Stats.ProcessingTimeInMs, 10)})
	statsTable.Append([]string{"Processing Total", strconv.FormatInt(job.Stats.ProcessingTotal, 10)})
	statsTable.Render()

	configBytes, err := json.MarshalIndent(job.Config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %v", err)
	}
	fmt.Printf("\n⚙️  Configuration:\n\n%s\n", string(configBytes))

	return nil
}

func (c *Client) DeleteRollupJob(name string) error {
	_, err := c.doRequest(http.MethodDelete, fmt.Sprintf("/_rollup/job/%s", name), nil)
	return err
}

func (c *Client) StartRollupJob(name string) error {
	_, err := c.doRequest(http.MethodPost, fmt.Sprintf("/_rollup/job/%s/_start", name), nil)
	return err
}

func (c *Client) StopRollupJob(name string) error {
	_, err := c.doRequest(http.MethodPost, fmt.Sprintf("/_rollup/job/%s/_stop", name), nil)
	return err
}

func (c *Client) PutRollupJob(name string, config map[string]interface{}) error {
	body, err := json.Marshal(config)
	if err != nil {
		return err
	}

	_, err = c.doRequest(http.MethodPut, fmt.Sprintf("/_rollup/job/%s", name), bytes.NewBuffer(body))
	if err != nil {
		if err1 := c.DeleteIndice(config["rollup_index"].(string)); err1 != nil {
			wtools.Error(fmt.Sprintf("Failed to delete rollup index[%s]: ", config["rollup_index"].(string)) + err.Error())
			return err1
		}
		_, err = c.doRequest(http.MethodPut, fmt.Sprintf("/_rollup/job/%s", name), bytes.NewBuffer(body))
	}
	return err
}

func (c *Client) doRequest(method, path string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, c.baseURL+path, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

func (c *Client) CheckRollupJob(name string) (bool, error) {
	resp, err := c.doRequest(http.MethodGet, fmt.Sprintf("/_rollup/job/%s", name), nil)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return false, nil
		}
		return false, err
	}

	var response RollupJobsResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		return false, fmt.Errorf("failed to parse response: %v", err)
	}

	for _, job := range response.Jobs {
		if job.Config.ID == name {
			c.GetRollupJob(name)
			return true, nil
		}
	}

	return false, nil
}
