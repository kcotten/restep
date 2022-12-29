package router

import "github.com/prometheus/client_golang/prometheus"

// create a new counter vector
var getRequestCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_get_count", // metric name
		Help: "Number of get requests.",
	},
	[]string{"status"}, // labels
)

// record request latency
var getRequestLatency = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "http_request_get_request_duration_seconds",
		Help:    "Latency of get request in second.",
		Buckets: prometheus.LinearBuckets(0.01, 0.05, 10),
	},
	[]string{"status"},
)

// initialize prometheus counter
func promInit() {
	// must register counter on init
	prometheus.MustRegister(getRequestCounter)
	prometheus.MustRegister(getRequestLatency)
}

// get latency measure based on status
func getTimer(status string) *prometheus.Timer {
	return prometheus.NewTimer(prometheus.ObserverFunc(func(v float64) {
		getRequestLatency.WithLabelValues(status).Observe(v)
	}))
}
