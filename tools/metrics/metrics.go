package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var DiffCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "Diff_request_count",
		Help: "No of request handled by Diff Handlers",
	},
)
