package cmd

// cron: https://cron.ciding.cc/

var (
	mapping = map[string]string{
		"vsd":         vsd,
		"vsp":         vsp,
		"cc-p-h1-acc": vsp,
		"cc-d-h1-acc": vsd,
		"lsep1":       lsep1,
		"lsep2":       lsep2,
	}
	indiceMapping = map[string]string{
		"cc-d-h1-acc": "cc-d-h1-acc",
		"cc-p-h1-acc": "cc-p-h1-acc",
		"vsp":         "cc-p-h1-acc",
		"vsd":         "cc-d-h1-acc",
		"lsep1":       "lsep1",
		"lsep2":       "lsep2",
	}
)

var (
	lsep1 = `{
  "index_pattern": "lsep1-*",
  "rollup_index": "rollup-lsep1",
  "cron": "0 0/5 * * * ?",
  "page_size": 5000,
  "groups": {
    "date_histogram": {
      "field": "time_local",
      "fixed_interval": "5m",
      "delay": "5m",
      "time_zone": "Asia/Shanghai"
    },
    "terms": {
      "fields": [
        "host.keyword",
        "client_ip.keyword",
        "status.keyword",
        "upstream_cache_status.keyword",
		"client_facing.keyword",
		"zone.keyword"
      ]
    }
  },
  "metrics": [
    {
      "field": "size",
      "metrics": ["sum"]
    }
  ]
}`
	lsep2 = `{
  "index_pattern": "lsep2-*",
  "rollup_index": "rollup-lsep2",
  "cron": "0 0/5 * * * ?",
  "page_size": 5000,
  "groups": {
    "date_histogram": {
      "field": "time_local",
      "fixed_interval": "5m",
      "delay": "5m",
      "time_zone": "Asia/Shanghai"
    },
    "terms": {
      "fields": [
        "host.keyword",
        "client_ip.keyword",
        "status.keyword",
        "upstream_cache_status.keyword",
		"client_facing.keyword",
		"zone.keyword"
      ]
    }
  },
  "metrics": [
    {
      "field": "size",
      "metrics": ["sum"]
    }
  ]
}`
	vsd = `{
  "index_pattern": "cc-d-h1-acc-*",
  "rollup_index": "rollup-cc-d-h1-acc",
  "cron": "0 0/5 * * * ?",
  "page_size": 5000,
  "groups": {
    "date_histogram": {
      "field": "time_local",
      "fixed_interval": "5m",
      "delay": "5m",
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
      "metrics": ["sum"]
    },
    {
      "field": "ivsd_vts_upstream_resp_body_len",
      "metrics": ["sum"]
    },
    {
      "field": "ivsd_vts_time_spent_upstream",
      "metrics": ["avg"]
    },
    {
      "field": "ivsd_vts_ttfb",
      "metrics": ["avg"]
    }
  ]
}`

	vsp = `{
  "index_pattern": "cc-p-h1-acc-*",
  "rollup_index": "rollup-cc-p-h1-acc",
  "cron": "0 0/5 * * * ?",
  "page_size": 5000,
  "groups": {
    "date_histogram": {
      "field": "time_local",
      "fixed_interval": "5m",
      "delay": "5m",
      "time_zone": "Asia/Shanghai"
    },
    "terms": {
      "fields": [
        "host.keyword",
        "client_ip.keyword",
        "status.keyword",
        "upstream_status.keyword",
        "upstream_cache_status.keyword",
		"client_facing.keyword",
		"zone.keyword"
      ]
    }
  },
  "metrics": [
    {
      "field": "size",
      "metrics": ["sum", "avg"]
    },
    {
      "field": "request_time",
      "metrics": ["avg"]
    }
  ]
}`
)
