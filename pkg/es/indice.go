package es

import (
	"bytes"
	"encoding/json"
	"fmt"
	"wrollup/wtools"
)

var (
	jobPrefix = "rollup-"
)

// DeleteOldData 删除指定索引中指定时间之前的数据
func (c *Client) DeleteOldData(indexPattern string, duration string) error {
	timestamp, err := wtools.ParseDuration(duration)
	if err != nil {
		return fmt.Errorf("failed to parse duration: %v", err)
	}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"range": map[string]interface{}{
				"@timestamp.date_histogram.timestamp": map[string]interface{}{
					"lt": timestamp,
				},
			},
		},
	}

	body, err := json.Marshal(query)
	if err != nil {
		return fmt.Errorf("failed to marshal query: %v", err)
	}

	path := fmt.Sprintf("/%s/_delete_by_query", jobPrefix+indexPattern)
	resp, err := c.doRequest("POST", path, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to delete old data: %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp, &result); err != nil {
		return fmt.Errorf("failed to parse response: %v", err)
	}

	if deleted, ok := result["deleted"].(float64); ok {
		fmt.Printf("Successfully deleted %d documents\n", int(deleted))
	}

	return nil
}

// DeleteIndice 删除指定索引
func (c *Client) DeleteIndice(index string) error {
	_, err := c.doRequest("DELETE", "/"+index, nil)
	if err != nil {
		return fmt.Errorf("failed to delete index %s: %v", index, err)
	}
	return nil
}
