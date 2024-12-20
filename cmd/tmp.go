package cmd

var mapping = map[string]string{
	"vsd": vsd,
}

var (
	vsd = `{
  "index_pattern": "vsd-*",
  "rollup_index": "rollup-vsd",
  "cron": "0 0/1 * * * ?",
  "page_size": 5000,
  "groups": {
    "date_histogram": {
      "field": "@timestamp",
      "fixed_interval": "1m",
      "delay": "1h"
    },
    "terms": {
      "fields": [ "ivsd_vts_http_method.keyword", "ivsd_vts_client_ip.keyword" ]
    }
  },
  "metrics": [
    {
      "field": "ivsd_vts_proxy_resp_body_len",
      "metrics": [ "min", "max", "sum" , "avg"]
    }
  ]
}`
)
