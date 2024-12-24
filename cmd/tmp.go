package cmd

var (
	mapping = map[string]string{
		"vsd": vsd,
		"vsp": vsp,
	}
)

var (
	vsd = `{
  "index_pattern": "vsd-access-*",
  "rollup_index": "rollup-vsd",
  "cron": "0 0/1 * * * ?",
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
        "ivsd_vts_cache_result.keyword"
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
  "index_pattern": "vsp-access-*",
  "rollup_index": "rollup-vsp",
  "cron": "0 0/1 * * * ?",
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
        "upstream_cache_status.keyword"
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
