package cmd

// cron: https://cron.ciding.cc/

var (
	mapping = map[string]string{
		"vsd": vsd,
		"vsp": vsp,
	}
)

var (
	vsd = `{
  "index_pattern": "cc-d-h1-acc-*",
  "rollup_index": "rollup-cc-d-h1-acc",
  "cron": "0 0/20 * * * ?",
  "page_size": 5000,
  "groups": {
    "date_histogram": {
      "field": "@timestamp",
      "fixed_interval": "5m",
      "delay": "20m",
      "time_zone": "Asia/Shanghai"
    },
    "terms": {
      "fields": [
        "ivsd_vts_http_method.keyword",
        "ivsd_vts_client_ip.keyword",
        "ivsd_vts_request_domain.keyword",
        "ivsd_vts_proxy_resp_status.keyword",
        "ivsd_vts_cache_result.keyword",
		"zone.keyword"
      ]
    }
  },
  "metrics": [
    {
      "field": "ivsd_vts_proxy_resp_body_len",
      "metrics": ["min", "max", "sum", "avg"]
    },
    {
      "field": "ivsd_vts_upstream_resp_body_len",
      "metrics": ["min", "max", "sum", "avg"]
    },
    {
      "field": "ivsd_vts_time_spent_upstream",
      "metrics": ["min", "max", "avg"]
    },
    {
      "field": "ivsd_vts_ttfb",
      "metrics": ["min", "max", "avg"]
    }
  ]
}`

	vsp = `{
  "index_pattern": "cc-p-h1-acc-*",
  "rollup_index": "rollup-cc-p-h1-acc",
  "cron": "0 0/20 * * * ?",
  "page_size": 5000,
  "groups": {
    "date_histogram": {
      "field": "@timestamp",
      "fixed_interval": "5m",
      "delay": "20m",
      "time_zone": "Asia/Shanghai"
    },
    "terms": {
      "fields": [
		"remote.keyword",
        "host.keyword",
        "client_ip.keyword",
        "status.keyword",
        "upstream_status.keyword",
        "upstream_cache_status.keyword",
		"client_facing.keyword",
		"zone.keyword"
      ]
    },
    "histogram": {
       "fields": ["request_time"],
       "interval": 2
    }
  },
  "metrics": [
    {
      "field": "size",
      "metrics": ["min", "max", "sum", "avg"]
    },
    {
      "field": "request_time",
      "metrics": ["min", "max", "avg"]
    },
    {
      "field": "upstream_response_time",
      "metrics": ["min", "max", "avg"]
    }
  ]
}`
)
