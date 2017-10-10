package main

import (
	// general
	"net/http"
	"os"

	// Prometheus for logging/metrics
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	// go-kit
	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	// Prometheus metrics settings for this service
	fieldKeys := []string{"method", "error"} // label names
	requestCount := kitprometheus.NewCounterFrom(prometheus.CounterOpts{
		Namespace: "EDGECAST_SERVICE",
		Subsystem: "service_metrics",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(prometheus.SummaryOpts{
		Namespace: "EDGECAST_SERVICE",
		Subsystem: "service_metrics",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(prometheus.SummaryOpts{
		Namespace: "EDGECAST_SERVICE",
		Subsystem: "service_metrics",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{}) // no fields here

	// create the actual service
	var svc EdgecastService
	svc = NewEdgecastService()
	svc = loggingMiddleware{logger, svc} // attach logger to service
	svc = instrumentingMiddleware{requestCount, requestLatency, countResult, svc}

	// connect handlers
	http.Handle("/metrics", promhttp.Handler())

	// set up logger and start service on port 8090
	logger.Log("msg", "HTTP", "addr", ":8090")
	logger.Log("err", http.ListenAndServe(":8090", nil))
}
