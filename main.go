package main

import (
	// general
	"net/http"
	"os"

	// Edgecast Client
	"github.com/iwilltry42/edgecast"

	// Prometheus for logging/metrics
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	// go-kit
	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
)

var (
	accountID = "testID"
	token     = "testToken"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	// Prometheus metrics settings for this service
	fieldKeys := []string{"method", "error"} // label names
	requestCount := kitprometheus.NewCounterFrom(prometheus.CounterOpts{
		Namespace: "EDGECAST",
		Subsystem: "service_metrics",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(prometheus.SummaryOpts{
		Namespace: "EDGECAST",
		Subsystem: "service_metrics",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)

	// create the actual service
	collector := NewEdgecastCollector(edgecast.NewEdgecastClient(accountID, token))
	prometheus.MustRegister(collector)
	var svc edgecast.Edgecast
	svc = collector.ec
	svc = loggingMiddleware{logger, svc} // attach logger to service
	svc = instrumentingMiddleware{requestCount, requestLatency, svc}

	// connect handlers
	http.Handle("/metrics", promhttp.Handler())

	// set up logger and start service on port 8090
	logger.Log("msg", "HTTP", "addr", ":8090")
	logger.Log("err", http.ListenAndServe(":8090", nil))
}
