package main

import (
	// general
	"net/http"
	"os"

	// Prometheus for logging/metrics
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	// go-kit
	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	// Prometheus metrics settings for this service
	fieldKeys := []string{"method", "error"} // label names
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "EDGECAST_SERVICE",
		Subsystem: "service_metrics",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "EDGECAST_SERVICE",
		Subsystem: "service_metrics",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
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

	// initiate server with service + endpoint
	getDataHandler := httptransport.NewServer(
		makeGetDataEndpoint(svc),
		nil,
		encodeResponse,
	)

	// connect handlers
	http.Handle("/", getDataHandler)
	http.Handle("/metrics2", promhttp.Handler())

	// set up logger and start service on port 8090
	logger.Log("msg", "HTTP", "addr", ":8090")
	logger.Log("err", http.ListenAndServe(":8090", nil))
}
